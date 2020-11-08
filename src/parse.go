package main

import (
	"strconv"
	"strings"
)

// grep every line.
func parse(text []string) {
	for num, line := range text {
		if strings.HasPrefix(line, "#define") {
			// run define task
			lineSplit := strings.Split(line, " ")
			key := lineSplit[1]
			value := lineSplit[2]
			Log("Set define " + key + " as " + value)
			u := User{
				name:      value,
				code:      key,
				moneyPush: make(map[string]string),
				moneyPull: make(map[string]string),
				money:     make(map[string]string),
			}
			user = append(user, u)
		} else if strings.HasPrefix(line, "#default") {
			defaultAdmin = parseUser(strings.Split(line, " ")[1])[0]
			Log("Set default admin as " + defaultAdmin.name)
		} else if strings.HasPrefix(line, "#") {
			// skip comment
			Log("Skip comment at #" + strconv.Itoa(num+1))
		} else if strings.TrimSpace(line) == "" {
			// skip blank line
			Log("Skip blank line at #" + strconv.Itoa(num+1))
		} else {
			if defaultAdmin.name == "" {
				defaultAdmin = user[0]
			}
			parseItem(line)
		}
	}
}

// item 21.1 a,c,d-b
func parseItem(text string) {
	// parse line for item
	line := strings.Split(text, " ")
	title := line[0]
	money := line[1]
	UserAndAdmin := line[2]
	if strings.Contains(UserAndAdmin, "-") {
		// 不使用默认 admin
		s := strings.Split(UserAndAdmin, "-")
		user := parseUser(s[0])
		admin := parseUser(s[1])[0]
		order := Item{
			title: title,
			money: money,
			users: user,
			admin: admin,
		}
		item = append(item, order)
	} else {
		u := parseUser(UserAndAdmin)
		order := Item{
			title: title,
			money: money,
			users: u,
			admin: defaultAdmin,
		}
		item = append(item, order)
	}
}

// a[,b,c,d,e]
func parseUser(text string) []User {
	var u []User
	if text == "full" {
		for _, v := range user {
			u = append(u, v)
		}
	} else if strings.Contains(text, ",") {
		split := strings.Split(text, ",")
		for _, s := range split {
			FindUser(s, func(a User) {
				u = append(u, a)
			})
		}
	} else {
		FindUser(text, func(a User) {
			u = append(u, a)
		})
	}
	return u
}

func FindUser(text string, success func(User)) {
	for _, u := range user {
		if u.code == text || u.name == text {
			success(u)
		}
	}
}
