package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

type Data struct {
	Code        int    `json:"code"`
	Name        string `json:"name"`
	Source      string `json:"source"`
	PackageName string `json:"packagename"`
	Data        []Item `json:"data"`
}

type Item struct {
	Id     int64  `json:"id"`
	Source string `json:"source"`
}

func main() {
	filepath.Walk("./data", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			output := info.Name()[0 : len(info.Name())-5]
			f, _ := os.Open(path)
			b, _ := ioutil.ReadAll(f)
			d := Data{}
			json.Unmarshal(b, &d)
			for index, v := range d.Data {
				fmt.Printf("%v-%v......%v\n", info.Name(), index, v.Source)
				err := download(v.Source, fmt.Sprintf("%v/%v.jpg", output, v.Id))
				if err != nil {
					fmt.Sprintf("%v\t%v\n", v.Source, err.Error())
				}
			}
		}
		return nil
	})
}

func download(source string, dest string) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", source, nil)
	if err != nil {
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	f, err := os.Create("./down/" + dest)
	if err != nil {
		return err
	}
	io.Copy(f, resp.Body)
	return nil
}
