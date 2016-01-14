package main

import (
	"fmt"
	"io/ioutil"
	// "log"
	"encoding/json"
	"net/http"
	"net/url"
	"os"
	"strings"
)

//Courier 列表配送员结构体
type Courier struct {
	Name    string `json:"courierName"`
	Tel     string `json:"courierTel"`
	Company string `json:"companyName"`
	Area    string `json:"workArea"`
}

//LandMark 地理坐标信息
type LandMark struct {
	Name      string `json:"name"`
	XZQNumber string `json:"xzqNumber"`
	XZQName   string `json:"xzqName"`
	CoreName  string `json:"corename"`
	City      string `json:"city"`
}

//CourierAround 列表结构体
type CourierAround struct {
	LandMark LandMark  `json:"landMark"`
	Colist   []Courier `json:"coList"`
	Message  string    `json:"message"`
}

func main() {
	//横纵双向遍历杭州地图, 步进0.01, 大约1公里范围
	for latitude, s1, count := 30.122561, 0, 0; s1 < 25; latitude, s1 = latitude+0.01, s1+1 {
		for longitude, s2 := 120.082626, 0; s2 < 30; longitude, s2 = longitude+0.01, s2+1 {
			count++
			fmt.Printf("%3d,latitude = %f,", count, latitude)
			fmt.Printf("longitude = %f\n", longitude)

		}
	}

	urlofKuaidi := "http://j.kuaidi100.com/searchapi.do"

	keyVal := url.Values{}
	keyVal.Set("json", "{longitude:120.2099324782554,latitude:30.20681568480997}")
	keyVal.Add("method", "courieraround")

	body := ioutil.NopCloser(strings.NewReader(keyVal.Encode())) //把form数据编下码
	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlofKuaidi, body)

	//这个一定要加，不加form的值post不过去，被坑了两小时
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	//看下发送的结构
	//fmt.Printf("%+v\n\n\n\n", req)

	//发送
	resp, _ := client.Do(req)
	//一定要关闭resp.Body
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(data))

	var courierList CourierAround
	err := json.Unmarshal(data, &courierList)
	fmt.Println(courierList.Colist[1].Name, err)

}

func SaveDataToFile(fileName string, data []byte) {
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
}
