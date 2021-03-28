package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var Path = flag.String("file", "./src.txt", "要解析的账单文件")
var Debug = flag.Bool("debug", false, "开启输出详细信息")
var Parse = flag.Bool("parse", false, "仅显示文件解析结果")
var Order = flag.Bool("order", true, "输出详细账单")
var Help = flag.Bool("help", false, "输出此帮助")

func main() {
	flag.Parse()

	if *Help {
		flag.Usage()
		os.Exit(0)
	}

	text := readFile(*Path)
	parse(text)

	if *Parse {
		fmt.Println("Parse Results: ")
		fmt.Printf("%-25s%-8s%-35s%-12s\n", "商品名", "价格", "购买人", "付钱人")
		for _, i := range item {
			fmt.Printf("%-25s%-8s%-35s%-12s\n", i.title, i.money, GetString(i.users), i.admin.name)
		}
		fmt.Println()
		os.Exit(0)
	}

	// 注册观察者并通知
	for _, i := range item {
		for _, u := range i.users {
			i.AddObs(u)
		}
	}

	// 通知观察者更新状态
	for _, i := range item {
		i.NotifyAll()
	}

	// 输出结果
	for _, u := range user {
		totalPull := "0.0"
		totalPush := "0.0"
		fmt.Println(u.name)
		// 输出购物清单
		if *Order {
			for k, v := range u.money {
				fmt.Printf("\t%6s --- %s\n", v, k)
			}
		}
		// 输出花钱项目
		for k, v := range u.moneyPush {
			totalPush = Add(totalPush, v)
			fmt.Println("\t- " + k + " " + v + " RMB.")
		}
		// 输出收钱项目
		for k, v := range u.moneyPull {
			totalPull = Add(totalPull, v)
			fmt.Println("\t+ " + k + " " + v + " RMB.")
		}
		// 输出总计
		fmt.Println("\t- 总共花费 " + totalPush + " RMB.")
		fmt.Println("\t+ 总共收入 " + totalPull + " RMB.")
	}
}

// Read file from path.
func readFile(path string) []string {
	Log("Read " + path)
	b, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(b), "\n")
}

func Log(v interface{}) {
	if *Debug {
		fmt.Println(v)
	}
}

func GetString(u []User) string {
	var builder strings.Builder
	for _, user := range u {
		builder.WriteString(user.name)
		builder.WriteString(",")
	}
	s := builder.String()
	s = s[:len(s)-1]
	return s
}
