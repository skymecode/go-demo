package main

import (
	"fmt"
	"time"
)

const NewRetailTTL = 3600 * 24 * 365 * 3

func main() {
	// 要转换的字符串
	dateTimeString := "2023-11-14 3:00:00"

	// 使用Parse函数将字符串转换为time.Time类型
	parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeString)
	if err != nil {
		fmt.Println("转换失败:", err)
		return
	}
	shanghaiLocation, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return
	}
	in := parsedTime.In(shanghaiLocation)
	nowTime := time.Now().In(shanghaiLocation)
	fmt.Println(time.Duration(NewRetailTTL - (int)(nowTime.Unix()-in.Unix())))
	// 打印转换后的time.Time类型
	fmt.Println("转换后的时间:", in)
}
