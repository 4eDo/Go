package main

import (
	"fmt"
	"strings"
)

func main (){
	out := comma("-111234567890.123456")
	fmt.Println(out)
}

func comma (s string) string {
	n := len(s)
	parts := strings.Split(s, ".")
	partsCount := len(parts)
	tail := 0
	sign := 0
	minus:= ""
	if partsCount > 2 {
		return "Это не число"
	}
	if partsCount > 1 {
		tail = len(parts[1]) + 1
	}
	if strings.HasPrefix(s, "-") {
		sign = 1
		minus = "-"
	}
	
	head := len(parts[0])
	
	if head <= 3 {
		return s
	}
	return minus + comma(s[sign:head-3]) + "," + s[head-3:head] + s[n-tail:]
}