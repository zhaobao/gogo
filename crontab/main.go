package main

import (
	"github.com/robfig/cron"
	"fmt"
)

func main() {

	c := cron.New()
	c.AddFunc("*/5 * * * * *", func() {
		fmt.Print("crontab run every 5 seconds\n")
	})
	c.Start()

	select {}
}
