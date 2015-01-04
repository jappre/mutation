package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	strOrigin := ""
	intNum := ""

	_, err1 := fmt.Scanln(&strOrigin)
	if nil == err1 {
		fmt.Printf("strOrigin=%s\n", strOrigin)
	}
	_, err2 := fmt.Scanln(&intNum)
	if nil == err2 {
		fmt.Printf("intNum=%s\n", intNum)
	}
	x, _ := strconv.Atoi(intNum)

	subStr := strings.Split(strOrigin, "")
	section := len(subStr) / x
	if len(subStr)%x != 0 {
		section++
	}
	fmt.Printf("x=%d\n", x)
	for i := 0; i < section; i++ {
		for j := 0; j < x/2; j++ {
			fmt.Printf("i = %d,j= %d \n", i, j)
			if x*(i+1)-j-1 < len(subStr) {
				a := subStr[x*(i+1)-j-1]
				fmt.Printf("a = %v \n", a)
				subStr[x*(i+1)-j-1] = subStr[x*i+j]
				subStr[x*i+j] = a
			}
		}
	}

	// subStrs := strings.SplitN(strOrigin, "", len(strOrigin)/x+1)
	fmt.Printf("x=%d\n", x)
	fmt.Printf("section=%d\n", section)
	fmt.Printf("subStr = %v\n", subStr)
}
