package main

import (
	"bytes"
	"gopkg.in/russross/blackfriday.v2"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	input, err := ioutil.ReadFile("./source/test.md")
	if err != nil {
		log.Fatal(err.Error())
	}
	header, _ := ioutil.ReadFile("./layout/header.html")
	footer, _ := ioutil.ReadFile("./layout/footer.html")
	output := blackfriday.Run(input)
	ioutil.WriteFile("./output/test.html", mergeBytes(header, output, footer), os.ModePerm)
}

func mergeBytes(pBytes ...[]byte) []byte {
	return bytes.Join(pBytes, []byte(""))
}
