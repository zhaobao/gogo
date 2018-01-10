package main

import (
	"encoding/csv"
	"encoding/json"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func main() {

	f, _ := os.Open("./data/d1")
	buffer, _ := ioutil.ReadAll(f)
	data := make([]Item, 0)
	json.Unmarshal(buffer, &data)

	e, _ := os.Create("./output/aff.csv")
	w := csv.NewWriter(e)
	w.Write([]string{
		"id", "name", "url", "thumb", "cats", "contries", "epc", "conv", "earning", "preview",
		"description", "offer_countries", "default_payout", "currency", "scree_shot", "is_adult",
	})
	for _, v := range data {
		w.Write([]string{
			v.Id, v.Name, v.Url, v.Thumbnail, formatCatsAndCountries(v.Cats), formatCatsAndCountries(v.Countries),
			v.Epc, v.Conv, v.Earnings, v.PreviewUrl, "", splitAndJoin(v.OfferCountries), v.DefaultPayout,
			v.Currency, "http://www.affiliaxe.com/" + v.ScreenShotUrl, v.IsAdult,
		})
	}
	w.Flush()
}

func splitAndJoin(str string) string {
	return strings.Join(strings.Split(str, ","), "|")
}

func formatCatsAndCountries(cats string) string {
	var c string
	reg := regexp.MustCompile(` s='(.*?)'`)
	reg_span := regexp.MustCompile(`>(.*?)<`)
	if reg.Match([]byte(cats)) {
		c = reg.FindStringSubmatch(cats)[1]
	} else {
		c = reg_span.FindStringSubmatch(cats)[1]
	}
	return splitAndJoin(c)
}

type Item struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	Url            string `json:"url"`
	Thumbnail      string `json:"thumbnail"`
	Cats           string `json:"cats"`
	Countries      string `json:"countries"`
	Epc            string `json:"epc"`
	Conv           string `json:"conv"`
	Earnings       string `json:"earnings"`
	PreviewUrl     string `json:"preview_url"`
	Description    string `json:"description"`
	OfferCountries string `json:"offer_countries"`
	DefaultPayout  string `json:"default_payout"`
	Currency       string `json:"currency"`
	ScreenShotUrl  string `json:"screenshot_url"` // http://www.affiliaxe.com/
	IsAdult        string `json:"is_adult"`
}
