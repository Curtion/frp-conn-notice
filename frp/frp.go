package frp

import (
	"io"
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

func IsOnline() func(string, string, string, string, string) (bool, error) {
	lastConns := 0
	return func(url string, user string, password string, frpname string, frptype string) (bool, error) {
		client := http.Client{}
		req, err := http.NewRequest("GET", url+"/api/proxy/"+frptype, nil)
		if err != nil {
			return false, err
		}
		req.SetBasicAuth(user, password)
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
			if name == frpname || name == "" {
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
