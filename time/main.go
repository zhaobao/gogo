package main

import (
	"fmt"
	"time"
)

func GetTimeDuration(from, to string) (int64, int64) {
	from_time, _ := time.Parse("2006/01/02 15:04:05", fmt.Sprintf("%v 00:00:00", from))
	to_time, _ := time.Parse("2006/01/02 15:04:05", fmt.Sprintf("%v 23:59:59", to))
	return from_time.Unix() - 8*3600, to_time.Unix() - 8*3600
}

func main() {
	start, to := GetTimeDuration("2017/10/19", "2017/10/20")
	//fmt.Printf("%v\t%v\n", start, to)
	//fmt.Println(time.Now().Format("2"), 30%30)
	fmt.Println(start, to)
}
