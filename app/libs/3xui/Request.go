package xui

import (
	"bytes"
	"io"
	"net/http"
	// "sync"
)

type Request struct {
	Cookie     string
	CookieTime int
	Url        string
}

func (r *Request) SendPost(req VpnRequest) string {
	path := req.GetPath()
	data := req.ToJson()
	url := r.Url + path
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	// тут преобразовывать в структуру
	return string(body)
}
