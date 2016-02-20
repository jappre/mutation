package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/fatih/set.v0"

	// "log"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

//全杭州 终点 120.395050,30.354519
// const LAT = 30.122561
// const LON = 120.082626

//主要城区 终点120.207974,30.338176
const LAT = 30.253608
const LON = 120.060795

//Courier 列表配送员结构体
type Courier struct {
	GUID     string     `json:"guid"`
	Name     string     `json:"courierName"`
	Tel      string     `json:"courierTel"`
	Company  string     `json:"companyName"`
	Area     string     `json:"workArea"`
	AreaList []LandMark `json:"landMarkList"`
}

//LandMark 地理坐标信息
type LandMark struct {
	GUID      string `json:"guid"`
	Name      string `json:"name"`
	XZQNumber string `json:"xzqNumber"`
}

//CourierAround 列表结构体
type CourierAround struct {
	LandMark LandMark  `json:"landMark"`
	Colist   []Courier `json:"coList"`
	Message  string    `json:"message"`
}

//CourierAround 列表结构体
type CourierDetail struct {
	Courier Courier `json:"courier"`
	Message string  `json:"message"`
}

var s = set.New()

func main() {
	//横纵双向遍历杭州地图, 步进0.01, 大约1公里范围
	for latitude, s1, count := LAT, 0, 0; s1 < 4; latitude, s1 = latitude+0.02, s1+1 {
		for longitude, s2 := LON, 0; s2 < 7; longitude, s2 = longitude+0.02, s2+1 {
			count++
			fmt.Printf("count = %3d\n", count)
			// fmt.Printf("longitude = %f\n", longitude)
			courierAround := GetAround(latitude, longitude)
			// time.Sleep(time.Second * 10)
			for i := 0; i < len(courierAround.Colist); i++ {
				courier, err := GetDetail(latitude, longitude, courierAround.Colist[i].GUID)
				if err != nil {
					fmt.Println("continue\n")
					continue
				}
				time.Sleep(time.Second * 2)
				if len(courier.AreaList) > 0 {
					SaveDataToFile("courier.txt", courier.Name+","+courier.Tel+","+courier.Company+","+courier.AreaList[0].Name+"\n")
				} else {
					SaveDataToFile("courier.txt", courier.Name+","+courier.Tel+","+courier.Company+","+"\n")
				}

			}
		}
	}
}

// makeJsonKeyValue 获得json的参数
func makeJsonKeyValue(latitude, longitude float64) string {
	return "{longitude:" + strconv.FormatFloat(longitude, 'G', -1, 32) + ",latitude:" + strconv.FormatFloat(latitude, 'G', -1, 32) + "}"
}

// makeJsonGuidKeyValue 获得json的参数
func makeJsonGuidKeyValue(latitude, longitude float64, guid string) string {
	return "{\"guid\":\"" + guid + "\",longitude:" + strconv.FormatFloat(longitude, 'G', -1, 32) + ",latitude:" + strconv.FormatFloat(latitude, 'G', -1, 32) + "}"
}

//GetAround 获取周边列表
func GetAround(latitude, longitude float64) CourierAround {
	keyVal := url.Values{}
	keyVal.Set("json", makeJsonKeyValue(latitude, longitude))
	keyVal.Add("method", "courieraround")
	data := GetKuaidiData(keyVal)

	var courierAround CourierAround
	json.Unmarshal(data, &courierAround)
	fmt.Printf("在%s附近, 有%d个快递员", courierAround.LandMark.Name, len(courierAround.Colist))
	fmt.Printf("\n")
	return courierAround
}

//GetDetail 获取详情
func GetDetail(latitude, longitude float64, guid string) (Courier, error) {
	if !s.Has(guid) {
		s.Add(guid)
	} else {
		return Courier{}, errors.New("already have")
	}
	keyVal := url.Values{}
	keyVal.Set("json", makeJsonGuidKeyValue(latitude, longitude, guid))
	keyVal.Add("method", "courierdetail")
	data := GetKuaidiData(keyVal)
	fmt.Println(string(data))

	var courierDetail CourierDetail
	json.Unmarshal(data, &courierDetail)
	fmt.Printf("快递员名字是%s, 电话是%s", courierDetail.Courier.Name, courierDetail.Courier.Tel)
	fmt.Printf("\n")
	return courierDetail.Courier, nil
}

// GetKuaidiData 获取快递100的数据, 需要传入参数
func GetKuaidiData(keyVal url.Values) []byte {
	urlofKuaidi := "http://j.kuaidi100.com/searchapi.do"

	body := ioutil.NopCloser(strings.NewReader(keyVal.Encode())) //把form数据编下码
	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlofKuaidi, body)

	//这个一定要加，不加form的值post不过去，被坑了两小时
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	//看下发送的结构
	// fmt.Printf("req = %+v\n\n\n\n", req)

	//发送
	resp, _ := client.Do(req)
	// fmt.Println("resp = ", resp)
	//一定要关闭resp.Body
	fmt.Printf("%p, %T", resp, resp)
	defer resp.Body.Close()
	respData, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(data))
	return respData
}

//SaveDataToFile 将string存入文件
func SaveDataToFile(fileName string, data string) {
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(err)
		return
	}
	f.WriteString(data)
}
