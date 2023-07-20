package utils

import (
	"fmt"
	"html/template"
	"math"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/prysmaticlabs/go-bitfield"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func FormatETH(num string) string {
	floatNum, _ := strconv.ParseFloat(num, 64)
	return fmt.Sprintf("%.4f", floatNum/math.Pow10(18)) + " ETH"
}

func FormatETHFromGwei(gwei uint64) string {
	return fmt.Sprintf("%.4f", float64(gwei)/math.Pow10(9)) + " ETH"
}

func FormatFullETHFromGwei(gwei uint64) string {
	return fmt.Sprintf("%v ETH", uint64(float64(gwei)/math.Pow10(9)))
}

func FormatFloat(num float64, precision int) string {
	p := message.NewPrinter(language.English)
	f := fmt.Sprintf("%%.%vf", precision)
	s := strings.TrimRight(strings.TrimRight(p.Sprintf(f, num), "0"), ".")
	r := []rune(p.Sprintf(s, num))
	return string(r)
}

func FormatAddCommasFormated(num float64, precision uint) template.HTML {
	p := message.NewPrinter(language.English)
	s := p.Sprintf(fmt.Sprintf("%%.%vf", precision), num)
	if precision > 0 {
		s = strings.TrimRight(strings.TrimRight(s, "0"), ".")
	}
	return template.HTML(strings.ReplaceAll(string([]rune(p.Sprintf(s, num))), ",", `<span class="thousands-separator"></span>`))
}

func FormatBigNumberAddCommasFormated(val hexutil.Big, precision uint) template.HTML {
	return FormatAddCommasFormated(float64(val.ToInt().Int64()), 0)
}

func FormatAddCommas(n uint64) template.HTML {
	number := FormatFloat(float64(n), 2)

	number = strings.ReplaceAll(number, ",", `<span class="thousands-separator"></span>`)
	return template.HTML(number)
}

func FormatBitlist(b []byte) template.HTML {
	p := bitfield.Bitlist(b)
	return formatBits(p.BytesNoTrim(), int(p.Len()))
}

func formatBits(b []byte, length int) template.HTML {
	var buf strings.Builder
	buf.WriteString("<div class=\"text-bitfield text-monospace\">")
	perLine := 8
	for y := 0; y < len(b); y += perLine {
		start, end := y*8, (y+perLine)*8
		if end >= length {
			end = length
		}
		for x := start; x < end; x++ {
			if x%8 == 0 {
				if x != 0 {
					buf.WriteString("</span> ")
				}
				buf.WriteString("<span>")
			}
			bit := BitAtVector(b, x)
			if bit {
				buf.WriteString("1")
			} else {
				buf.WriteString("0")
			}
		}
		buf.WriteString("</span><br/>")
	}
	buf.WriteString("</div>")
	return template.HTML(buf.String())
}

func formatBitvectorValidators(bits []byte, validators []uint64) template.HTML {
	invalidLen := false
	if len(bits)*8 != len(validators) {
		invalidLen = true
	}
	var buf strings.Builder
	buf.WriteString("<pre class=\"text-monospace\" style=\"font-size:1rem;\">")
	for i := 0; i < len(bits)*8; i++ {
		if invalidLen {
			if BitAtVector(bits, i) {
				buf.WriteString("1")
			} else {
				buf.WriteString("0")
			}
		} else {
			val := validators[i]
			if BitAtVector(bits, i) {
				buf.WriteString(fmt.Sprintf("<span title=\"Validator %[1]d\" data-validx=\"%[1]d\">1</span>", val))
			} else {
				buf.WriteString(fmt.Sprintf("<span title=\"Validator %[1]d\" data-validx=\"%[1]d\">0</span>", val))
			}
		}

		if (i+1)%64 == 0 {
			buf.WriteString("\n")
		} else if (i+1)%8 == 0 {
			buf.WriteString(" ")
		}
	}
	buf.WriteString("</pre>")
	return template.HTML(buf.String())
}

func FormatParticipation(v float64) template.HTML {
	return template.HTML(fmt.Sprintf("<span>%.2f %%</span>", v*100.0))
}

func FormatAmountFormatted(amount *big.Int, unit string, digits int, maxPreCommaDigitsBeforeTrim int, fullAmountTooltip bool, smallUnit bool, newLineForUnit bool) template.HTML {
	return formatAmount(amount, unit, digits, maxPreCommaDigitsBeforeTrim, fullAmountTooltip, smallUnit, newLineForUnit)
}
func FormatAmount(amount *big.Int, unit string, digits int) template.HTML {
	return formatAmount(amount, unit, digits, 0, true, false, false)
}
func FormatBigAmount(amount *hexutil.Big, unit string, digits int) template.HTML {
	return FormatAmount((*big.Int)(amount), unit, digits)
}
func FormatBytesAmount(amount []byte, unit string, digits int) template.HTML {
	return FormatAmount(new(big.Int).SetBytes(amount), unit, digits)
}
func formatAmount(amount *big.Int, unit string, digits int, maxPreCommaDigitsBeforeTrim int, fullAmountTooltip bool, smallUnit bool, newLineForUnit bool) template.HTML {
	// define display unit & digits used per unit max
	displayUnit := " " + unit
	var unitDigits int
	if unit == "ETH" || unit == "Ether" {
		unitDigits = 18
	} else if unit == "GWei" {
		unitDigits = 9
	} else {
		displayUnit = " ?"
		unitDigits = 0
	}

	// small unit & new line for unit handling
	{
		unit = displayUnit
		if newLineForUnit {
			displayUnit = "<BR />"
		} else {
			displayUnit = ""
		}
		if smallUnit {
			displayUnit += `<span style="font-size: .63rem;`
			if newLineForUnit {
				displayUnit += `color: grey;`
			}
			displayUnit += `">` + unit + `</span>`
		} else {
			displayUnit += unit
		}
	}

	trimmedAmount, fullAmount := trimAmount(amount, unitDigits, maxPreCommaDigitsBeforeTrim, digits, false)
	tooltip := ""
	if fullAmountTooltip {
		tooltip = fmt.Sprintf(` data-toggle="tooltip" data-placement="top" title="%s"`, fullAmount)
	}

	// done, convert to HTML & return
	return template.HTML(fmt.Sprintf("<span%s>%s%s</span>", tooltip, trimmedAmount, displayUnit))
}

func trimAmount(amount *big.Int, unitDigits int, maxPreCommaDigitsBeforeTrim int, digits int, addPositiveSign bool) (trimmedAmount, fullAmount string) {
	// Initialize trimmedAmount and postComma variables to "0"
	trimmedAmount = "0"
	postComma := "0"
	proceed := ""

	if amount != nil {
		s := amount.String()
		if amount.Sign() > 0 && addPositiveSign {
			proceed = "+"
		} else if amount.Sign() < 0 {
			proceed = "-"
			s = strings.Replace(s, "-", "", 1)
		}
		l := len(s)

		// Check if there is a part of the amount before the decimal point
		if l > int(unitDigits) {
			// Calculate length of preComma part
			l -= unitDigits
			// Set preComma to part of the string before the decimal point
			trimmedAmount = s[:l]
			// Set postComma to part of the string after the decimal point, after removing trailing zeros
			postComma = strings.TrimRight(s[l:], "0")

			// Check if the preComma part exceeds the maximum number of digits before the decimal point
			if maxPreCommaDigitsBeforeTrim > 0 && l > maxPreCommaDigitsBeforeTrim {
				// Reduce the number of digits after the decimal point by the excess number of digits in the preComma part
				l -= maxPreCommaDigitsBeforeTrim
				if digits < l {
					digits = 0
				} else {
					digits -= l
				}
			}
			// Check if there is only a part of the amount after the decimal point, and no leading zeros need to be added
		} else if l == unitDigits {
			// Set postComma to part of the string after the decimal point, after removing trailing zeros
			postComma = strings.TrimRight(s, "0")
			// Check if there is only a part of the amount after the decimal point, and leading zeros need to be added
		} else if l != 0 {
			// Use fmt package to add leading zeros to the string
			d := fmt.Sprintf("%%0%dd", unitDigits-l)
			// Set postComma to resulting string, after removing trailing zeros
			postComma = strings.TrimRight(fmt.Sprintf(d, 0)+s, "0")
		}

		fullAmount = trimmedAmount
		if len(postComma) > 0 {
			fullAmount += "." + postComma
		}

		// limit floating part
		if len(postComma) > digits {
			postComma = postComma[:digits]
		}

		// set floating point
		if len(postComma) > 0 {
			trimmedAmount += "." + postComma
		}
	}
	return proceed + trimmedAmount, proceed + fullAmount
}