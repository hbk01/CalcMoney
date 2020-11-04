package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"sync"
)

var wait sync.WaitGroup

func main() {
	file := os.Args[1]
	b, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	text := strings.Split(string(b), "\n")
	parse(text)
	fmt.Println("== defines")
	for _, define := range defines {
		fmt.Println(define.key, define.value)
	}
	fmt.Println()
	fmt.Println("== orders")
	fmt.Println("__________________________________________________")
	for _, order := range orders {
		fmt.Printf("| %8s | %.2f | %s-(%s) |\n", order.title, order.money, order.user, order.admin)
	}
	fmt.Println("--------------------------------------------------")
	fmt.Println()
	for _, define := range defines {
		wait.Add(1)
		go calc(define)
	}
	wait.Wait()
	fmt.Println("===========")
	UserMoney.Range(func (a, b interface{}) bool {
		fmt.Println(a.(string), b.(string))
		return true
	})
}

func calc(define Define) {
	for _, order := range orders {
		for _, user := range order.user {
			if user == define.value {
				// 几个人
				p := len(order.user)
				// 每个人多少钱
				money := Div(F2S(order.money), I2S(p))


				// 计算我要给谁多少钱
				d := make(map[string]string, p)
				if m, ok := d[order.admin]; ok {
					d[order.admin] = Add(d[order.admin], m)
				} else {
					d[order.admin] = money
				}
				fmt.Println("============", user)
				for k, v := range d {
					fmt.Println(k, v)
				}
				fmt.Println("============")

				// 只能计算每个 User 用了多少钱
				// TODO User: 我要给哪些 Admin 多少钱？
				// TODO Admin: 哪些 User 要还我多少钱？
				if m, ok := UserMoney.Load(user); ok {
					UserMoney.Store(user, Add(m.(string), money))
				} else {
					UserMoney.Store(user, money)
				}
			}
		}
	}
	wait.Done()
}

func parse(text []string) {
	for _, line := range text {
		if strings.HasPrefix(line, "# define") {
			// run define task
			lineSplit := strings.Split(line, " ")
			key := lineSplit[2]
			value := lineSplit[3]
			PutDefine(key, value)
		} else if strings.HasPrefix(line, "#") {
			// skip comment
		} else if strings.TrimSpace(line) == "" {
			// skip blank line
		} else {
			parseLine(line)
		}
	}
}

// item 21.1 a,c,d-b
func parseLine(text string) {
	// parse line for item
	line := strings.Split(text, " ")
	title := line[0]
	money, err := strconv.ParseFloat(line[1], 64)
	if err != nil {
		panic("can't parse money.")
	}
	userAdmin := line[2]
	var user []string
	var admin string
	if strings.Contains(userAdmin, "-") {
		UserAdmin := strings.Split(userAdmin, "-")
		user = parseUser(UserAdmin[0])
		admin = FindDefine(UserAdmin[1]).value
	} else {
		user = parseUser(userAdmin)
		admin = defines[0].value
	}
	order := Order{
		title: title,
		money: money,
		user:  user,
		admin: admin,
	}
	orders = append(orders, order)
}

// a,b,c,d,e
func parseUser(text string) (u []string) {
	if text == "full" {
		for _, define := range defines {
			u = append(u, define.value)
		}
	} else if strings.Contains(text, ",") {
		split := strings.Split(text, ",")
		for _, s := range split {
			u = append(u, FindDefine(strings.TrimSpace(s)).value)
		}
	} else {
		u = append(u, FindDefine(strings.TrimSpace(text)).value)
	}
	return u
}

