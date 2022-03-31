package main

import (
	"FrpConnNotice/bark"
	"FrpConnNotice/frp"
	"log"
	"time"

	"github.com/go-ini/ini"
)

var check func(string, string, string, string, string) (bool, error)

var cfg *ini.File

func init() {
	check = frp.IsOnline()
	config, err := ini.Load("config.ini")
	if err != nil {
		log.Fatal("Fail to read file: ", err)
	}
	cfg = config
}

func main() {
	second, err := cfg.Section("config").Key("time").Int64()
	if err != nil {
		log.Fatal(err)
	}
	frpurl := cfg.Section("frp").Key("frp_dashboard").String()
	frpuser := cfg.Section("frp").Key("frp_user").String()
	frppassword := cfg.Section("frp").Key("frp_password").String()
	frpname := cfg.Section("frp").Key("frp_conn_name").String()
	frptype := cfg.Section("frp").Key("frp_conn_type").String()
	for range time.Tick(time.Duration(second) * time.Second) {
		status, err := check(frpurl, frpuser, frppassword, frpname, frptype)
		if err != nil {
			log.Print(err)
		}
		if status {
			barkurl := cfg.Section("bark").Key("url").String()
			bark.Notice(barkurl)
		} else {
			log.Print("无需通知")
		}
	}
}
