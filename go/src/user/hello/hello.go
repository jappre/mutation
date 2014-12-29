package main

import (
	"fmt"
	// "time"
	//"user/newmath"
	// "strconv"
	// "sort"
)

func main() {
	c := make(chan bool)
	go func() {
		fmt.Println("go go go !")
		<-c
		// close(c)
	}()
	c <- true
}
