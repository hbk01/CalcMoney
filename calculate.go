package main

import (
	"github.com/shopspring/decimal"
	"strconv"
)

var (
	x decimal.Decimal
	y decimal.Decimal
	err error
)

func checkError() {
	if err != nil {
		panic(err)
	}
}

func Add(a, b string) string {
	x, err = decimal.NewFromString(a)
	checkError()
	y, err = decimal.NewFromString(b)
	checkError()
	return x.Add(y).String()
}

func Sub(a, b string) string {
	x, err = decimal.NewFromString(a)
	checkError()
	y, err = decimal.NewFromString(b)
	checkError()
	return x.Sub(y).String()

}

func Mul(a, b string) string {
	x, err = decimal.NewFromString(a)
	checkError()
	y, err = decimal.NewFromString(b)
	checkError()
	return x.Mul(y).String()

}

func Div(a, b string) string {
	x, err = decimal.NewFromString(a)
	checkError()
	y, err = decimal.NewFromString(b)
	checkError()
	return x.Div(y).StringFixed(2)
}

// Int to String
func I2S(a int) string {
	return strconv.Itoa(a)
}

// Float(64) to String
func F2S(a float64) string {
	return strconv.FormatFloat(a, 'f', 2, 64)
}

// String to Int
func S2I(a string) int {
	result, err := strconv.Atoi(a)
	if err != nil {
		panic(err)
	}
	return result
}

// String to Float(64)
func S2F(a string) float64 {
	result, err := strconv.ParseFloat(a, 64)
	if err != nil {
		panic(err)
	}
	return result
}