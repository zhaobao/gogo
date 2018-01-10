package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	items := walkData("./data")
	var output string
	output += fmt.Sprintf("%v#%v#%v#%v#%v#%v#%v#%v#%v#%v#%v\n",
		"country", "category", "create_time", "devices", "channels",
		"logo", "name", "offer_id", "payout", "permission", "preview_url")
	for _, v := range items {
		//country := v.Country
		//country_str := fmt.Sprintf("%v", country)
		// open 应该代表打开
		// Required to apply 应该是下载
		//if strings.ToLower(v.Permission) == "open" &&
		//	(strings.Contains(country_str, "GH") ||
		//		len(country) == 0 ||
		//		strings.Contains(country_str, "All")) {
		//	fmt.Printf("%v,%v\n", v.Name, v.PreviewUrl)
		//}
		//if strings.ToLower(v.Permission) == "open" {
		var channels []string
		for _, v := range v.Channel {
			channels = append(channels, fmt.Sprintf("%v", v))
		}
		output += fmt.Sprintf("%v#%v#%v#%v#%v#%v#%v#%v#%v#%v#%v\n",
			strings.Join(v.Country, "|"), strings.Join(v.Category, "|"),
			v.CreateTime, strings.Join(v.Device, "|"), strings.Join(channels, "|"),
			v.Logo, v.Name, v.OfferId, v.Payout, v.Permission, v.PreviewUrl)
		//}
	}
	ioutil.WriteFile("output.txt", []byte(output), os.ModePerm)
}

type Body struct {
	D Data   `json:"data"`
	F string `json:"flag"`
	M string `json:"msg"`
}

type Data struct {
	L []Item `json:"list"`
}

type Item struct {
	Country    []string      `json:"country"`
	Category   []string      `json:"category"`
	CreateTime string        `json:"create_time"`
	Device     []string      `json:"device"`
	Channel    []interface{} `json:"channel"`
	Logo       string        `json:"logo"`
	Name       string        `json:"name"`
	OfferId    int64         `json:"offer_id"`
	Payout     interface{}   `json:"payout"`
	Permission string        `json:"permission"`
	PreviewUrl string        `json:"preview_url"`
}

func walkData(root string) []Item {
	data := make([]Item, 0)
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && strings.HasSuffix(info.Name(), "json") {
			f, _ := os.Open(path)
			buffer, _ := ioutil.ReadAll(f)
			d := Body{}
			json.Unmarshal(buffer, &d)
			data = append(data, d.D.L...)
		}
		return nil
	})
	return data
}
