package xui

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"sync"

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

func (r *Request) SendPost(req VpnRequest) string {
	if r.Cookie == "" {
		r.requestCookie()
	}
	// path := req.GetMethod()
	// data := req.ToJson()
	// url := r.Host + ":" + r.Port + "/" + r.Hash + "/" + path
	// spew.Dump(url)
	// resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	// spew.Dump(err)
	// if err != nil {
	// }

	// defer resp.Body.Close()
	// body, _ := io.ReadAll(resp.Body)
	// spew.Dump(string(body))
	// spew.Dump(resp.Cookies())
	// // тут преобразовывать в структуру
	// return string(body)
	return ""
}

func (r *Request) requestCookie() {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()
	if r.Cookie != "" {
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
	spew.Dump(err)
	if err != nil {
	}
	// Дописать либу распарсить куки и записывать их в структуру как и время протухания

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	spew.Dump(string(body))
	spew.Dump(resp.Cookies())
	r.Cookie = "1"
	// тут преобразовывать в структуру
}
