package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	readCsv()
}

type IpLocation struct {
	From        int64
	To          int64
	CountryCode string
	CountryName string
}

func toInt64(str string) int64 {
	r, _ := strconv.ParseInt(str, 10, 64)
	return r
}

func readCsv() {
	f, _ := os.Open("./input/ip.CSV")
	r := csv.NewReader(f)
	result, err := r.ReadAll()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		for _, row := range result {
			line := IpLocation{From: toInt64(row[0]), To: toInt64(row[1]), CountryCode: row[2], CountryName: row[3]}
			if line.CountryCode != "-" {
				fmt.Printf("%v\n", line)
			}
		}
	}
}

func writeCsv() {
	f, _ := os.Create("x.csv")

	records := [][]string{
		{"first_name", "last_name", "username"},
		{"Rob", "Pike", "rob"},
		{"Ken", "Thompson", "ken"},
		{"Robert", "Griesemer", "gri"},
	}

	w := csv.NewWriter(f)

	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}

	// Write any buffered data to the underlying writer (standard output).
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}
