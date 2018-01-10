package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func Md5(message string) string {
	h := md5.New()
	h.Write([]byte(message))
	return fmt.Sprintf("%s", hex.EncodeToString(h.Sum(nil)))
}

const UP_SERVER_KEY string = "bqWIKi9XpIsvLsW5"
const UP_SERVER_API string = "http://125.212.226.157:22222/down"

var targetDate string = "20170829"

func main() {
	sig := Md5(fmt.Sprintf("%s%s%s", "vnd48", targetDate, UP_SERVER_KEY))
	fmt.Println(fmt.Sprintf(UP_SERVER_API+"/vnd48/"+targetDate+"/%v.viettel", sig))
	sig = Md5(fmt.Sprintf("%s%s%s", "vnd59", targetDate, UP_SERVER_KEY))
	fmt.Println(fmt.Sprintf(UP_SERVER_API+"/vnd59/"+targetDate+"/%v.viettel", sig))
	sig = Md5(fmt.Sprintf("%s%s%s", "vnd7", targetDate, UP_SERVER_KEY))
	fmt.Println(fmt.Sprintf(UP_SERVER_API+"/vnd7/"+targetDate+"/%v.viettel", sig))
}
