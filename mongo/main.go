package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type UserTable struct {
	Uid             int64  `bson:"uid",json:"uid"`
	AppId           int    `bson:"app_id",json:"app_id"`
	CreateTime      int64  `bson:"create_time",json:"create_time"`
	SubscribeTime   int64  `bson:"subscribe_time",json:"subscribe_time"`
	UnSubscribeTime int64  `bson:"unsubscribe_time",json:"unsubscribe_time"`
	Country         string `bson:"country",json:"country"`
	Status          int    `bson:"status",json:"status"`
	Channel         string `bson:"channel",json:"channel"`
	ChannelUid      string `bson:"channel_uid",json:"channel_uid"`
	ChannelOfferId  string `bson:"channel_offer_id",json:"channel_offer_id"`
	ChannelPayout   string `bson:"channel_payout",json:"channel_payout"`
	ChannelAffId    string `bson:"channel_aff_id",json:"channel_aff_id"`
	ChannelMeta     string `bson:"channel_meta",json:"channel_meta"`
	Operator        string `bson:"operator",json:"operator"`
	ServiceId       string `bson:"service_id",json:"service_id"`
	Phone           string `bson:"phone",json:"phone"`
	AppName         string `bson:"app_name",json:"app_name"`
	PayChannel      string `bson:"pay_channel",json:"pay_channel"`
	PayTrackId      string `bson:"pay_track_id",json:"pay_track_id"`
	PaySubscribeId  string `bson:"pay_subscribe_id",json:"pay_subscribe_id"`
}

func main() {
	se, err := mgo.Dial("mongodb://subscribe:palmax123456@127.0.0.1:27017/subscribe")
	if err != nil {
		log.Fatal("connect to mongo error: " + err.Error())
	}
	defer se.Close()

	c := se.DB("subscribe").C("user")

	users := make([]UserTable, 0)
	user := UserTable{}
	//q := c.Find(bson.M{"status": 1, "country": "vn", "phone": bson.M{"$regex": "/67589823/"}}).Iter()
	q := c.Find(bson.M{"phone": bson.M{"$regex": "/67589823/"}}).Iter()
	for q.Next(&user) {
		fmt.Printf("%v\n", user)
		users = append(users, user)
	}

	//c.Update(bson.M{"uid": 21}, bson.M{"$inc": bson.M{"wap_count": 1}})
}
