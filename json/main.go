package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type VideoItem struct {
	Id         int64      `json:"id"`
	Url        string     `json:"url"`
	Size       int64      `json:"size"`
	Duration   float64    `json:"duration"`
	Title      string     `json:"title"`
	Width      int64      `json:"width"`
	Height     int64      `json:"height"`
	UpdateTime int64      `json:"update_time"`
	Views      int64      `json:"views"`
	Starts     int64      `json:"starts"`
	Cover      VideoCover `json:"cover"`
	Download   int64      `json:"download"`
	Cate       string     `json:"cate"`
	Src        string     `json:"src"`
}

type VideoCover struct {
	Id     int64  `json:"id"`
	Url    string `json:"url"`
	Size   int64  `json:"size"`
	Width  int64  `json:"width"`
	Height int64  `json:"height"`
	Src    string `json:"src"`
}

func main() {
	f, err := os.Open("./data/source.json")
	if err != nil {
		log.Fatal(err.Error())
	}
	data := make([]VideoItem, 0)
	buffer, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = json.Unmarshal(buffer, &data)
	if err != nil {
		log.Fatal(err.Error())
	}

	f, _ = os.Open("./data/from")
	buffer, _ = ioutil.ReadAll(f)
	lines := string(buffer)
	from := make(map[string]string)
	for _, v := range strings.Split(lines, "\n") {
		parts := strings.Split(v, ",")
		if len(parts) == 2 {
			from[parts[0]] = parts[1]
		}
	}

	filterData := make([]VideoItem, 0)
	for _, v := range data {
		if cate, ok := from[v.Src]; ok {
			v.Cate = cate
			v.UpdateTime = time.Now().Unix() + randomInt(0, 86400*7)
			filterData = append(filterData, v)
		}
	}

	output, _ := json.Marshal(filterData)
	ioutil.WriteFile("./data/videoTh.json", output, os.ModePerm)
}

func randomInt(min, max int64) int64 {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Int63n(max) // [0~max)
}
