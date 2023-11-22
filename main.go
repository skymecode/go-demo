package main

import "fmt"

func main() {
	/*
		add := mathClass.Add(10, 12)
			fmt.Println(add)
	*/
	var stockcode = 123
	var enddate = "2020-12-31"
	//%d表示整型,%s表示字符串
	var url = "Code=%d&endDate=%s"
	var target_url = fmt.Sprintf(url, stockcode, enddate) //这个是要返回字符串
	fmt.Println(target_url)
	fmt.Printf(url, stockcode, enddate) //根据格式化参数生成格式化的字符串并写入标准输出
}
