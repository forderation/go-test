package main

import (
	"fmt"
	"strings"
	"testing"
	"testing/quick"
)

type RomanNumerals []RomanNumeral

func (r RomanNumerals) ValueOf(symbols ...byte) int {
	symbol := string(symbols)
	for _, s := range r {
		if s.Symbol == symbol {
			return s.Value
		}
	}
	return 0
}

type RomanNumeral struct {
	Value  int
	Symbol string
}

var allRomanNumerals = RomanNumerals{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(arabic int) string {
	var result strings.Builder
	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}
	return result.String()
}

func ConvertoToArabic(roman string) int {
	total := 0
	for i := 0; i < len(roman); i++ {

		// symbol is byte of string
		symbol := roman[i]

		// look ahead to next symbol if we can and, the current symbol is base 10 (only valid subtractors)
		if couldBeSubtractive(i, symbol, roman) {
			nextSymbol := roman[i+1]

			// build the two character string
			// potentialNumber := []byte{symbol, nextSymbol}

			if value := allRomanNumerals.ValueOf(symbol, nextSymbol); value != 0 {
				total += value
				i++
			} else {
				total += allRomanNumerals.ValueOf(symbol)
			}
		} else {
			total += allRomanNumerals.ValueOf(symbol)
		}
	}
	return total
}

func couldBeSubtractive(index int, currentSymbol uint8, roman string) bool {
	isSubtractiveSymbol := currentSymbol == 'I' || currentSymbol == 'X' || currentSymbol == 'C'
	return index+1 < len(roman) && isSubtractiveSymbol
}

type Numeral struct {
	Arabic int
	Roman  string
}

var cases []Numeral = []Numeral{
	{1, "I"},
	{2, "II"},
	{3, "III"},
	{4, "IV"},
	{5, "V"},
	{20, "XX"},
	{21, "XXI"},
	{40, "XL"},
	{47, "XLVII"},
	{49, "XLIX"},
	{95, "XCV"},
	{412, "CDXII"},
	{1984, "MCMLXXXIV"},
	{1006, "MVI"},
	{2014, "MMXIV"},
}

func TestRomanNumerals(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%v gets converted to %v", test.Arabic, test.Roman), func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)
			if got != test.Roman {
				t.Errorf("got %q want %q", got, test.Roman)
			}
		})
	}
}

func TestConvertingToArabic(t *testing.T) {
	for _, test := range cases[:7] {
		t.Run(fmt.Sprintf("%q gets converted to %d", test.Roman, test.Arabic), func(t *testing.T) {
			got := ConvertoToArabic(test.Roman)
			if got != test.Arabic {
				t.Errorf("got %d want %d", got, test.Arabic)
			}
		})
	}
}

func TestProptertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {
		if arabic > 3999 {
			return true
		}
		roman := ConvertToRoman(int(arabic))
		fromRoman := ConvertoToArabic(roman)
		return fromRoman == int(arabic)
	}

	if err := quick.Check(assertion, &quick.Config{
		MaxCount: 1000,
	}); err != nil {
		t.Error("failed checks", err)
	}
}
