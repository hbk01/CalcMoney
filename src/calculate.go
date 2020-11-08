package main

import (
	"github.com/shopspring/decimal"
)

var (
	x   decimal.Decimal
	y   decimal.Decimal
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

//func Sub(a, b string) string {
//	x, err = decimal.NewFromString(a)
//	checkError()
//	y, err = decimal.NewFromString(b)
//	checkError()
//	return x.Sub(y).String()
//}
//
//func Mul(a, b string) string {
//	x, err = decimal.NewFromString(a)
//	checkError()
//	y, err = decimal.NewFromString(b)
//	checkError()
//	return x.Mul(y).String()
//}

func Div(a, b string) string {
	x, err = decimal.NewFromString(a)
	checkError()
	y, err = decimal.NewFromString(b)
	checkError()
	return x.Div(y).StringFixed(2)
}
