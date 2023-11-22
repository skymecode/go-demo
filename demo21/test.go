package main

import (
	"fmt"
	"time"
)

func main() {
	//var a []byte
	//fmt.Println(len(a))
	//fmt.Println(time.Now().Unix())
	//i := int(time.Now().Unix() - 1698910935)
	//fmt.Println(94608000 - i)

	utcTimeString := "2023-11-11 10:05:03"
	utcTime, err := time.ParseInLocation("2006-01-02 15:04:05", utcTimeString, time.UTC)
	if err != nil {
		fmt.Println("时间解析错误:", err)
		return
	}

	// 加载上海的时区
	shanghaiLocation, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("时区加载错误:", err)
		return
	}

	// 转换为上海时区的时间
	shanghaiTime := utcTime.In(shanghaiLocation)
	fmt.Printf("UTC 时间：%v\n", utcTime)
	fmt.Printf("上海时区时间：%v\n", shanghaiTime)
	fmt.Println(utcTime.UnixNano() / 1e9)
	fmt.Println(shanghaiTime.UnixNano() / 1e9)

}
