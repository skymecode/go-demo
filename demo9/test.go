package main

import "context"

type MsgModel interface {
	Persist(context context.Context, msg interface{}) bool
	PersistOnSensitive(context context.Context, session_type, level, SensitiveStatus int32, msg interface{}) bool
}

// 定义一个struct用来实现接口类型
type msgModelImpl struct{}

// 定义一个变量MsgModelImpl等于msgModelImpl，相当于可以通过MsgModelImpl来调用msgModelImpl的成员
var MsgModelImpl = msgModelImpl{}

func main() {

}
