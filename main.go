package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s, s1 string
	_, _ = fmt.Scan(&s, &s1)
	fmt.Print(adding(s, s1))
}

func adding(s, s1 string) int64 {
	number1, err := strconv.Atoi(digitString(s))
	if err != nil {
		panic(err)
	}
	number2, err2 := strconv.Atoi(digitString(s1))
	if err2 != nil {
		panic(err2)
	}
	return int64(number1) + int64(number2)
}

func digitString(s string) string {
	newString := ""
	for i := 0; i < len(s); i++ {
		if rune(s[i]) >= 48 && rune(s[i]) <= 57 {
			newString += string(s[i])
		}
	}
	return newString
}
