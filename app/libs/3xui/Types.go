package xui

import (
	"encoding/json"
)

type VpnRequest interface {
	ToJson() []byte
	GetPath() string
}

type LoginRequest struct {
	Username string
	Password string
}

func (r *LoginRequest) getPath() string {
	return "/login"
}

func (r *LoginRequest) ToJson() []byte {
	json, _ := json.Marshal(*r)
	return json
}
