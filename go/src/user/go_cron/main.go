package main

import (
	"fmt"
	"os"
	"time"

	"github.com/robfig/cron"
)

func main() {
	c := cron.New()
	c.AddFunc("1 * * * * *", func() {
		f, _ := os.OpenFile("path.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
		f.WriteString(time.Now().String())
		f.Close()
		fmt.Println(time.Now().String())
	})
	c.Start()
	select {}
}
