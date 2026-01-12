package xui

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/davecgh/go-spew/spew"
)

type Request struct {
	Cookie     string
	CookieTime int
	Hash       string
	Port       string
	Host       string
	XuiUser    string
	XuiPass    string
	Mutex      sync.Mutex
}

func (r *Request) SendPost(req VpnRequest) (response VpnResponse) {
	if r.Cookie == "" || r.CookieTime < int(time.Now().Unix()) {
		r.requestCookie()
	}
	path := string(req.GetMethod())
	url := r.Host + ":" + r.Port + "/" + r.Hash + "/" + path
	http_req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// обработать
	}

	http_req.Header.Add("Cookie", "3x-ui="+r.Cookie)

	http_client := &http.Client{}
	resp, err := http_client.Do(http_req)
	if err != nil {
		// обработать
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	return r.buildResultBody(body, XuiPath(req.GetMethod()))
}

func (r *Request) buildResultBody(data []byte, request XuiPath) (response VpnResponse) {
	switch request {
	case Inbounds:
		var body ListResponse

		err := json.Unmarshal(data, &body)
		if err != nil {
			spew.Dump(err)
		}
		return body
	}

	panic(1)
}

func (r *Request) requestCookie() {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()
	if r.CookieTime != 0 && r.CookieTime < int(time.Now().Unix()) {
		return
	}
	data := map[string]string{
		"username": r.XuiUser,
		"password": r.XuiPass,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		// обработать ошибки
	}
	url := r.Host + ":" + r.Port + "/" + r.Hash + "/login"

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
	}

	defer resp.Body.Close()

	for _, cookie := range resp.Cookies() {
		// spew.Dump(cookie.Name)
		if cookie.Name == "3x-ui" {
			r.Cookie = cookie.Value
			r.CookieTime = int(cookie.Expires.Unix())
		}
	}
}
