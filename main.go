package main

import (
	"io"
	"log"
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

type msg struct {
	logo    string
	content string
}

func (m msg) send() (string, error) {
	resp, err := http.Get("" + m.content + "?icon=" + m.logo)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func notice() {
	notice := msg{
		logo:    "https://www.minecraft.net/etc.clientlibs/minecraft/clientlibs/main/resources/img/menu/menu-buy.gif",
		content: "MC有玩家上线了",
	}
	res, err := notice.send()
	if err != nil {
		log.Print(err)
	}
	resJson := []byte(res)
	message := jsoniter.Get(resJson, "message").ToString()
	log.Print("发送通知:", message)
}

func isOnline() func() (bool, error) {
	lastConns := 0
	return func() (bool, error) {
		client := http.Client{}
		req, err := http.NewRequest("GET", "", nil)
		if err != nil {
			return false, err
		}
		req.SetBasicAuth("", "")
		resp, err := client.Do(req)
		if err != nil {
			return false, err
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return false, err
		}
		resJson := []byte(body)
		var conns int
		for i := 0; ; i++ {
			name := jsoniter.Get(resJson, "proxies", i, "name").ToString()
			if name == "mc" || name == "" {
				conns = jsoniter.Get(resJson, "proxies", 0, "cur_conns").ToInt()
				break
			}
		}
		if conns > lastConns {
			lastConns = conns
			return true, nil
		} else {
			lastConns = conns
			return false, nil
		}
	}
}

func main() {
	// notice()
	check := isOnline()
	check()
}
