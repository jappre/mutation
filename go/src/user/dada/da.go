package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	// "time"
)

//main 主函数
func main() {
	for i := 0; i < 107020; i++ {
		// time.Sleep(1 * time.Second)
		fmt.Println(i)
		// fmt.Println("http://mp.imdada.cn/center/dada/informs/" + strconv.Itoa(i) + "/")
		// resp, err := http.Get("http://mp.imdada.cn/center/dada/" + strconv.Itoa(i) + "/")

		// client := &http.Client{}
		// req, _ := http.NewRequest("GET", "http://mp.imdada.cn/center/dada/informs/"+strconv.Itoa(i)+"/", nil)
		// req.Header.Set("name", "value")
		// resp, err := client.Do(req)

		url := "http://mp.imdada.cn/center/dada/" + strconv.Itoa(i) + "/"

		req, _ := http.NewRequest("GET", url, nil)

		req.Header.Add("user-agent", "Mozilla/5.0 (Linux; Android 4.4.2; GT-I9300 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/30.0.0.0 Mobile Safari/537.36")
		req.Header.Add("host", "mp.imdada.cn")
		req.Header.Add("sdcard-id", "972d887ef0634fe3b07c15207207a02f")
		req.Header.Add("city-id", "1")
		req.Header.Add("os-version", "4.3.1")
		req.Header.Add("lng", "118.276521")
		req.Header.Add("network", "WIFI")
		req.Header.Add("location-time", "1437315740271")
		req.Header.Add("uuid", "ffffffff-b692-f40c-c996-9d1c00000001")
		req.Header.Add("user-token", "140175177e5a1590c73d5b03b4b8975b")
		req.Header.Add("channel-id", "B004")
		req.Header.Add("enable-gps", "1")
		req.Header.Add("city-code", "021")
		req.Header.Add("model", "HUAWEI-7")
		req.Header.Add("user-id", "38696")
		req.Header.Add("app-name", "a-dada")
		req.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
		req.Header.Add("accuracy", "75.0")
		req.Header.Add("lat", "28.298731")
		req.Header.Add("unique-id", "ffffffff-9e73-c20q-ffff-ffffc9969d1c")
		req.Header.Add("platform", "Android")
		req.Header.Add("location-provider", "lbs")
		req.Header.Add("app-version", "2.7")
		req.Header.Add("accept-encoding", "gzip,deflate")
		req.Header.Add("accept-language", "zh-CN,en-US;q=0.8")
		req.Header.Add("x-requested-with", "com.dada.mobile.android")

		resp, err := http.DefaultClient.Do(req)

		// resp, err := http.Get("http://mp.imdada.cn/center/dada/informs/" + strconv.Itoa(i) + "/")
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
