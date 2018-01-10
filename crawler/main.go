package main

import (
	"github.com/gocolly/colly"
	"fmt"
	"strconv"
	"regexp"
	"os"
	"log"
	"time"
)

func main() {
	f, err := os.OpenFile("./"+string(os.PathSeparator)+"log"+string(os.PathSeparator)+"output.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err == nil {
		log.SetFlags(log.Ldate | log.Ltime)
		log.SetOutput(f)
	}
	defer f.Close()
	c := colly.NewCollector(
		colly.AllowedDomains("www.mzitu.com"),
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36"),
	)
	c.Limit(&colly.LimitRule{
		DomainGlob: "*",
		Parallelism: 10,
		Delay: time.Second * 2,
		})
	c.OnHTML(`#pins a`, func(element *colly.HTMLElement) {
		href := element.Attr("href")
		if element.Request.URL.Path == "/" {
			go element.Request.Visit(href)
		}
	})
	c.OnHTML(`div.pagenavi`, func(element *colly.HTMLElement) {
		path := element.Request.URL.Path
		re := regexp.MustCompile(`^/(\d+)$`)
		if re.MatchString(path) {
			ch := element.DOM.Children()
			total := ch.Eq(ch.Length() - 2).Text()
			pages, err := strconv.Atoi(total)
			if err == nil {
				for i := 1; i <= pages; i++ {
					href := element.Request.URL.Path + "/" + fmt.Sprintf("%v", i)
					element.Request.Ctx.Put("_meta", "1")
					go element.Request.Visit(href)
				}
			}
		}
	})
	c.OnHTML(`div.main-image p img`, func(element *colly.HTMLElement) {
		if len(element.Request.Ctx.Get("_meta")) > 0 {
			re := regexp.MustCompile(`^/(\d+)`)
			item := make(map[string]string, 3)
			item["gid"] = re.FindStringSubmatch(element.Request.URL.Path)[1]
			item["img_url"] = element.Attr("src")
			item["source_url"] = element.Request.AbsoluteURL(element.Request.URL.String())
			log.Println(fmt.Sprintf("%v %v %v", item["gid"], item["img_url"], item["source_url"]))
		}
	})
	c.OnRequest(func(request *colly.Request) {
		fmt.Println("visiting >>>>>>>>>> " + request.URL.String())
	})
	c.Visit("http://www.mzitu.com/")
	c.Wait()
}