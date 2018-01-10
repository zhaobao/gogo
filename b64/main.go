package main

import (
	"encoding/base64"
	"fmt"
	"net/url"
)

func main() {
	from := []byte("name=zhaobao&gender=1&redirect=http://www.baidu.com?name=zhaobao")
	to := base64.URLEncoding.EncodeToString(from)
	fmt.Println(to)

	ret, _ := base64.URLEncoding.DecodeString(to)
	fmt.Println(string(ret))

	values, _ := url.ParseQuery(string(ret))
	fmt.Println(values.Get("redirect"))
}
