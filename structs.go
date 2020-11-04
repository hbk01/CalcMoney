package main

import "sync"

var (
	defines    []Define
	orders     []Order
	UserMoney  sync.Map
	AdminMoney sync.Map
)

type Define struct {
	key   string
	value string
}

type Order struct {
	title string
	money float64
	user  []string
	admin string
}

func PutDefine(key string, value string) {
	for _, define := range defines {
		if define.key == key {
			return
		}
	}
	defines = append(defines, Define{key, value})
}

func FindDefine(key string) (d Define) {
	for _, define := range defines {
		if define.key == key {
			return define
		}
	}
	return
}
