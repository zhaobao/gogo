package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for i := 0; i < 100; i++ {
		rand.Seed(time.Now().UTC().UnixNano())
		get := rand.Intn(5555) + 1
		fmt.Printf("%v\n", get)
	}
}
