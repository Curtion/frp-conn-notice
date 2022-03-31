package bark

import (
	"io"
	"log"
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

func send(url string) (string, error) {
	resp, err := http.Get(url)
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

func Notice(url string) {
	res, err := send(url)
	if err != nil {
		log.Print(err)
	}
	resJson := []byte(res)
	message := jsoniter.Get(resJson, "message").ToString()
	log.Print("发送通知:", message)
}
