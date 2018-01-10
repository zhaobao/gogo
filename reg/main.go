package main

import (
	"regexp"
	"fmt"
)

func main() {
	str := "/11211/123"
	re := regexp.MustCompile(`/(\d+)`)
	fmt.Println(re.MatchString(str))
	fmt.Println(re.FindStringSubmatch(str)[1])
}
