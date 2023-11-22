package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"time"
)

type RedisConfig struct {
	Alias          string   `json:"alias" mapstructure:"alias"`
	Address        []string `json:"address" mapstructure:"address"`
	Password       string   `json:"password" mapstructure:"password"`
	DB             int      `json:"db" mapstructure:"db"`
	ConnectTimeout int      `json:"connect_timeout" mapstructure:"connect_timeout"`
	ReadTimeout    int      `json:"read_timeout" mapstructure:"read_timeout"`
	WriteTimeout   int      `json:"write_timeout" mapstructure:"write_timeout"`
	Wait           bool     `json:"wait" mapstructure:"wait"`
	MaxIdle        int      `json:"max_idle" mapstructure:"max_idle"`
	IdleTimeout    int      `json:"idle_timeout" mapstructure:"idle_timeout"`
}
type Config struct {
	KukaRedirect    string         `json:"kuka_redirect" toml:"kuka_redirect" mapstructure:"kuka_redirect"`
	HTTPListen      string         `json:"http_listen" toml:"http_listen" mapstructure:"http_listen"`
	ProfPort        string         `json:"admin_address" toml:"admin_address" mapstructure:"admin_address" `
	KukaRateLimit   int            `json:"kuka_rate_limit" toml:"kuka_rate_limit" mapstructure:"kuka_rate_limit"`
	CacheEnable     bool           `json:"cache_enable" toml:"cache_enable" mapstructure:"cache_enable"`
	SdkHost         string         `json:"sdk_host" toml:"sdk_host" mapstructure:"sdk_host"`
	RedisConfigList []*RedisConfig `json:"redis" toml:"redis" mapstructure:"redis"`
}

var (
	gConfig Config
)

func main() {
	v := viper.New()
	v.SetConfigFile("F:\\gogo\\src\\GoProject\\demo25\\test.toml")
	v.SetConfigType("toml")

	err := v.ReadInConfig()
	if err != nil {
		return
	}
	// 尝试将配置解析到结构体
	if err := v.Unmarshal(&gConfig); err != nil {
		fmt.Println("Error unmarshalling config:", err)
		return
	}
	index := -1
	for i, conf := range gConfig.RedisConfigList {

		if conf.Alias == "cache" {
			index = i
			break
		}
	}
	if index == -1 {
		fmt.Println("redis读取错误")

		return
	}
	cfg := gConfig.RedisConfigList[index]
	client := redis.NewClusterClient(&redis.ClusterOptions{Addrs: cfg.Address})
	fmt.Println(cfg.Address)
	result, err := client.Ping().Result()
	fmt.Println(result)
	//key := fmt.Sprintf("%v:%v:%v", "tb", "test", "testone")
	//err = b.Set(key, []byte("id"), time.Duration(93865728)*time.Second)
	//if err != nil {
	//	return
	//}

	fmt.Println("hello")

}

type Base struct {
	KeyPrefix string
	*redis.ClusterClient
}

func (c Base) Set(key string, val []byte, expiration time.Duration) error {
	key = c.KeyPrefix + key

	cmd := c.ClusterClient.Set(key, val, expiration)
	return cmd.Err()
}
