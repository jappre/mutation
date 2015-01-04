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
		fmt.Printf("strOrigin=%s", strOrigin)
	}
	_, err2 := fmt.Scanln(&intNum)
	if nil == err2 {
		fmt.Printf("intNum=%s", intNum)
	}
	x, _ := strconv.Atoi(intNum)

	subStrs := make([][]string, 10)

	for a := 0; a < x; a++ {
		t := strings.
	}

	// subStrs := strings.SplitN(strOrigin, "", len(strOrigin)/x+1)
	fmt.Println(subStrs)
}
