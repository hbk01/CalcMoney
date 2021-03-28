package main

import (
	"strconv"
)

var (
	observer     = make(map[string][]IObserver)
	defaultAdmin User
	user         []User
	item         []Item
)

// observer
type IObserver interface {
	UpdateAll(Item)
}

// implement IItem
type Item struct {
	title string
	money string
	users []User
	admin User
}

// implement IObserver
type User struct {
	name      string
	code      string
	moneyPull map[string]string
	moneyPush map[string]string
	money     map[string]string
}

func (i Item) AddObs(obs IObserver) {
	observer[i.title] = append(observer[i.title], obs)
}

func (i Item) NotifyAll() {
	for _, o := range observer[i.title] {
		o.UpdateAll(i)
	}
}

func (u User) UpdateAll(i Item) {
	userNum := len(i.users)                       // 买的人数
	sMoney := Div(i.money, strconv.Itoa(userNum)) // 平摊每个人多少钱
	if u.name != i.admin.name {
		// 更新当前用户的购物清单数据库
		AutoAdd(u.money, i.title, sMoney)
		// 更新当前用户的给钱数据库
		AutoAdd(u.moneyPush, i.admin.name, sMoney)
		// 更新买单者的收钱数据库
		AutoAdd(i.admin.moneyPull, u.name, sMoney)
	} else {
		// 更新当前用户的购物清单数据库
		AutoAdd(u.money, i.title, sMoney)
		// 更新当前用户的给钱数据库
		AutoAdd(u.moneyPush, i.admin.name, sMoney)
	}
}

// p: what db?
// name: who?
// money: money
func AutoAdd(db map[string]string, name string, money string) {
	if _, ok := db[name]; ok {
		m := db[name]
		db[name] = Add(m, money)
	} else {
		db[name] = money
	}
}
