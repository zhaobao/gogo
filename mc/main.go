package main

import (
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	"log"
)

var mcClient *memcache.Client
var appName string

func Init(appName string, mcs []string) {
	appName = appName
	mcClient = memcache.New(mcs...)
	err := mcClient.Set(&memcache.Item{Key: appName, Value: []byte("ok")})
	if err != nil {
		log.Fatal("[error] connect to memcache failed")
	}
}

func GetMc() *memcache.Client {
	return mcClient
}

func GetItem(key string) (*memcache.Item, error) {
	return mcClient.Get(appName + "_" + key)
}

func DeleteItem(key string) error {
	return mcClient.Delete(appName + "_" + key)
}

func SetItem(key string, value []byte, expr ...int32) {
	key = appName + "_" + key
	item := &memcache.Item{Key: key, Value: value}
	if len(expr) > 0 {
		item.Expiration = expr[0]
	}
	mcClient.Set(item)
}

func main() {
	var mcs []string = []string{"127.0.0.1:11211"}
	Init("test", mcs)
	SetItem("name", []byte("zhaobao"))
	i, _ := GetItem("name")
	fmt.Println(string(i.Value))
}
