package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	// "time"
)

func main() {
	for i := 39100; i < 39200; i++ {
		// time.Sleep(1 * time.Second)
		fmt.Println(i)
		// fmt.Println("http://mp.imdada.cn/center/dada/informs/" + strconv.Itoa(i) + "/")
		// resp, err := http.Get("http://mp.imdada.cn/center/dada/" + strconv.Itoa(i) + "/")
		resp, err := http.Get("http://mp.imdada.cn/center/dada/informs/" + strconv.Itoa(i) + "/")
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}
		if resp.StatusCode == http.StatusOK {
			fmt.Println(resp.StatusCode)
			fmt.Println("\n")
		}
		defer resp.Body.Close()

		buf := make([]byte, 1024)
		f, err1 := os.OpenFile("path.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm) //可读写，追加的方式打开（或创建文件）
		if err1 != nil {
			panic(err1)
			return
		}
		// defer f.Close()

		for {
			n, _ := resp.Body.Read(buf)
			if 0 == n {
				break
			}
			f.WriteString("\n" + string(buf[:n]) + "," + strconv.Itoa(i))
		}
		f.Close()
	}
}
