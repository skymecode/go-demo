package main

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"github.com/vmihailenco/msgpack"
)

var ErrNotFountCache = errors.New("not found")

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis 服务器地址
	})
	ctx := context.Background()
	key := "mydata" // Redis 键名
	var a interface{} = "hello"
	marshal, _ := msgpack.Marshal(a)
	// 存储一个字符串到 Redis
	err := redisClient.Set(ctx, key, marshal, 0).Err()
	if err != nil {
		fmt.Println(err)
		return
	}

	// 从 Redis 获取数据
	data, err := redisClient.Get(ctx, key).Bytes()

	if err != nil {
		fmt.Println(err)
		return
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	err = test()
	fmt.Println(err == ErrNotFountCache)
	var order string
	test2(data, &order)
	fmt.Println(order)
}
func test2(data []byte, obj interface{}) {
	err := msgpack.Unmarshal(data, obj)
	if err != nil {
		fmt.Println(err)
		fmt.Println("出错了")
		return
	}
}

func test() error {
	return ErrNotFountCache
}
