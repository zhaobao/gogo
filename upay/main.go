package main

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
	"zhaobao/upay/utils"
)

type Line struct {
	Used       bool   `json:"used"`
	App        string `json:"app"`
	Tid        string `json:"tid"`
	Phone      string `json:"phone"`
	Status     string `json:"status"`
	Channel    string `json:"channel"`
	ChannelAid string `json:"channel_aid"`
	CreateTime int64  `json:"create_time"`
	FreeTime   string `json:"free_time"`
}

const DB string = "zb:Zhaobao_123@tcp(127.0.0.1:3306)/subscribe?charset=utf8"
const SAVE_DIR string = "./dr"
const UP_SERVER_KEY string = "bqWIKi9XpIsvLsW5"
const UP_SERVER_API string = "http://125.212.226.157:22222/down"

var offers []string = []string{"vnd59", "vnd7", "vnd48", "vnd61"}

func init() {
	utils.InitDB(DB)
}

func main() {

	from := "20170719"
	to := "20170831"
	download(from, to)

	country := "vn"
	data := analysis(from, to)

	pmap := make(map[string]int)
	result := make(map[string]map[string]map[string]int)

	for target, v := range data {
		for _, line := range v {
			pmap[line.Phone] = 0
			sql := `select channel, channel_aid, create_time
			from user_history
			where phone = ? and country = ?
			order by id asc limit 1 offset ` + fmt.Sprintf("%v", pmap[line.Phone])
			err := utils.DB.QueryRow(sql, line.Phone, country).
				Scan(&line.Channel, &line.ChannelAid, &line.CreateTime)
			if err == nil {
				pmap[line.Phone]++
				line.Used = true
			} else {
				line.Channel = "-"
			}
			if line.Channel != "" {
				if _, ok := result[line.Channel]; !ok {
					result[line.Channel] = make(map[string]map[string]int)
				}
				if _, ok := result[line.Channel][target]; !ok {
					result[line.Channel][target] = make(map[string]int)
				}
				if _, ok := result[line.Channel][target][line.App]; !ok {
					result[line.Channel][target][line.App] = 0
				}
				result[line.Channel][target][line.App]++
			}
			sql = strings.Replace(sql, "\n", "", -1)
			sql = strings.Replace(sql, "	", " ", -1)
			sql = strings.Replace(sql, "?", "%v", -1)
			sql = fmt.Sprintf(sql, "\""+line.Phone+"\"", "\"vn\"")
			writeLog(sql, "sql.log")
		}
	}

	saveFile := "result.txt"
	writeCsv(fmt.Sprintf("%v,%v,%v,%v", "channel", "date", "app", "count"), saveFile)
	for channel, v := range result {
		for date, vv := range v {
			for app, count := range vv {
				writeCsv(fmt.Sprintf("%v,%v,%v,%v", channel, date, app, count), saveFile)
			}
		}
	}
}

func analysis(from, to string) map[string][]Line {
	data := make(map[string][]Line)
	for _, o := range offers {
		t, _ := time.Parse("20060102", from)
		for i := 0; i < 365; i++ {
			mts := t.AddDate(0, 0, i).Format("20060102")
			if _, ok := data[mts]; !ok {
				data[mts] = make([]Line, 0)
			}
			file := fmt.Sprintf("%v/%v-%v.txt", SAVE_DIR, o, mts)
			_, err := os.Open(file)
			if err != nil {
				continue
			}
			bt, _ := ioutil.ReadFile(file)
			lines := strings.Split(string(bt), "\n")
			for _, line := range lines {
				if len(line) > 0 {
					columns := strings.Split(line, "|")
					if len(columns) == 6 && strToInt(columns[5]) == 200 {
						data[mts] = append(data[mts], Line{FreeTime: columns[0][0:8], App: columns[1], Tid: columns[2], Phone: columns[3], Status: columns[5]})
					}
				}
			}
			if mts == to {
				break
			}
		}
	}
	return data
}

func download(from, to string) {
	var url, sig string
	for _, o := range offers {
		t, _ := time.Parse("20060102", from)
		for i := 0; i < 365; i++ {
			mts := t.AddDate(0, 0, i).Format("20060102")
			sig = md5Str(fmt.Sprintf("%s%s%s", o, mts, UP_SERVER_KEY))
			url = fmt.Sprintf(fmt.Sprintf("%s/%s/%s/%s.viettel", UP_SERVER_API, o, mts, sig))
			err, _ := downloadDrReport(url, fmt.Sprintf("%v-%v.txt", o, mts))
			if err != nil {
				fmt.Println(err.Error())
			}
			if mts == to {
				break
			}
		}
	}
}

func downloadDrReport(link, path string) (error, string) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, link, nil)
	if err != nil {
		return errors.New("http get send request error:" + err.Error()), ""
	}
	resp, err := client.Do(req)
	if err != nil {
		return errors.New("http get do error:" + err.Error()), ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New("http get read body error:" + err.Error()), ""
	}
	ioutil.WriteFile(SAVE_DIR+"/"+path, body, os.ModePerm)
	return nil, string(body)
}

func md5Str(message string) string {
	h := md5.New()
	h.Write([]byte(message))
	return fmt.Sprintf("%s", hex.EncodeToString(h.Sum(nil)))
}

func strToInt(p string) int {
	r, _ := strconv.Atoi(p)
	return r
}

func writeLog(message, file string) {
	go func() {
		f, err := os.OpenFile("logs"+string(os.PathSeparator)+file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err == nil {
			log.SetFlags(log.Ldate | log.Ltime)
			log.SetOutput(f)
			log.Println(message)
		}
		defer f.Close()
	}()
}

func writeCsv(message, file string) {
	f, err := os.OpenFile("logs"+string(os.PathSeparator)+file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err == nil {
		log.SetFlags(0)
		log.SetOutput(f)
		log.Println(message)
	}
	defer f.Close()
}
