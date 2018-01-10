package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"reflect"
)

func Md5(message string) string {
	h := md5.New()
	h.Write([]byte(message))
	return fmt.Sprintf("%s", hex.EncodeToString(h.Sum(nil)))
}

//
//func GetTimeDuration(from, to string) (int64, int64) {
//	from_time, _ := time.Parse("2006/01/02 15:04:05", fmt.Sprintf("%v 00:00:00", from))
//	to_time, _ := time.Parse("2006/01/02 15:04:05", fmt.Sprintf("%v 23:59:59", to))
//	return from_time.Unix() - 8*3600, to_time.Unix() - 8*3600
//}
//
//const key string = "zhaobao"

//func getMissCount(key string, current int64) int64 {
//	if v, ok := missesMap[key]; ok {
//		for i := 0; i < len(v["c"]); i++ {
//			if v["c"][i] >= current {
//				if i == 0 {
//					return 0
//				} else {
//					return v["v"][i-1]
//				}
//			}
//		}
//		return v["v"][len(v["v"])-1]
//	}
//	return 0
//}
//
//var missesMap map[string]map[string][]int64
//
//func hit(max int, current int64) bool {
//	rand.Seed(time.Now().UTC().UnixNano())
//	get := rand.Intn(max) + 1
//	return int64(get) <= current
//}
//
//type NitItem struct {
//	Id string
//}
//type NitTest struct {
//	i NitItem
//}

// mobidea mobusi top

func toArray(x interface{}) []interface{} {
	v := reflect.ValueOf(x)
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		fmt.Println(v.Field(i))
		values[i] = v.Field(i).Interface()
	}
	return values
}

func main() {
	//param := map[string]interface{}{"vcode": 1, "app_id": 4}
	//p, _ := json.Marshal(param)
	//fmt.Println(string(p))
	//fmt.Println(Md5("1393ead961daa531d0880aace8c7f180" + "session" + string(p)))
	//fmt.Printf("http://125.212.226.157:22222/down/vnd48/20170720/%v.viettel\n", Md5("vnd4820170720bqWIKi9XpIsvLsW5"))
	//fmt.Println(time.Now().Unix())
	//today_date := time.Now().Format("2006/01/02")
	//fmt.Println(today_date)
	//fmt.Println(fmt.Sprintf("%v 23:59:59", today_date))
	//today_end, _ := time.Parse("2006/01/02 15:04:05", fmt.Sprintf("%v 23:59:59", today_date))
	//fmt.Println(today_end.Unix() - 8*3600)
	//from, to := GetTimeDuration("2017/07/25", "2017/07/25")
	//fmt.Printf("from %v now %v to %v\n", from, time.Now().Unix(), to)
	//var m map[string]string = map[string]string{
	//	key: "game",
	//}
	//fmt.Printf("%#v\n", m)

	//missesMap = make(map[string]map[string][]int64)
	//missesMap["vn-bl-vnd7"] = make(map[string][]int64)
	//missesMap["vn-bl-vnd7"]["c"] = []int64{100, 200, 300, 400, 500}
	//missesMap["vn-bl-vnd7"]["v"] = []int64{10, 15, 20, 25, 30}
	//missesMap["th-up-photo"] = make(map[string][]int64)
	//missesMap["th-up-photo"]["c"] = []int64{1}
	//missesMap["th-up-photo"]["v"] = []int64{1000}
	//
	//fmt.Println(hit(1000, getMissCount("th-up-photo", 888)))

	//fmt.Println(len(missesMap))
	//n := NitTest{}
	//fmt.Printf("%v", len(get_ips()))
	//fmt.Printf("%v", n.i == nil)

	//InStringAry("abc", get_ips())
	//fmt.Println(NitTest{}.i.Id == "")
	//fmt.Printf("weekday:%v %v\n", int(time.Now().Weekday()), time.Wednesday)

	//thumb := "https://i.ytimg.com/vi/tSjqx_0RarU/default.jpg"
	//if strings.HasSuffix(thumb, "default.jpg") {
	//	thumb = strings.Replace(thumb, "default.jpg", "hqdefault.jpg", -1)
	//}
	//fmt.Println(thumb)

	//fmt.Println(string(os.PathSeparator))

	// /content/th/mk/ugame/20170812/3
	//path := "/content/th/mk/ugame/20170812/3"
	//parts := strings.Split(path, "/")
	//fmt.Printf("%#v - %v", len(parts), parts[6])

	//holders := strings.Repeat("?,", 10)
	//fmt.Println(holders[0 : len(holders)-1])

	//x := struct {
	//	Name string
	//	Age  uint
	//}{"zhaobao", 12}
	//fmt.Println(toArray(x))

	fmt.Println("zhaobao"[0:3])

}

//func get_ips() []string {
//	return nil
//}
