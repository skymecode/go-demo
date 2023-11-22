package main

import (
	"fmt"
	"reflect"
)

type User struct {
	username string `username:"username"`
	password string `password:"password"`
}

func GetTag(r interface{}) {
	elem := reflect.TypeOf(r).Elem()
	for i := 0; i < elem.NumField(); i++ {
		fmt.Println("执行")
		username := elem.Field(i).Tag.Get("username")
		password := elem.Field(i).Tag.Get("password")
		fmt.Println(username)
		fmt.Println(password)
	}
}

func main() {
	user := &User{username: "hello", password: "test"}
	GetTag(user)
	username := user.username
	fmt.Println(username)
}
