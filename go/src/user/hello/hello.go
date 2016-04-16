package main

import (
	"fmt"
	"strconv"
	"strings"

	log "github.com/Sirupsen/logrus"
)

func main() {
	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 100,
	}).Fatal("The ice breaks!")
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
