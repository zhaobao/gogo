package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type MkPushResponse struct {
	XMLName  xml.Name `xml:"push-response"`
	Tid      string   `xml:"tid"`
	StatusId string   `xml:"status-id"`
}

func main() {
	w := MkPushResponse{}
	f, err := os.Open("./d.xml")
	if err != nil {
		log.Fatal(err.Error())
	}
	body, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = xml.Unmarshal(body, &w)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("%#v\n", w)
}
