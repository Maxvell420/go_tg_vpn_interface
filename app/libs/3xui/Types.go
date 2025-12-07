package xui

import (
	"encoding/json"
)

type VpnRequest interface {
	ToJson() []byte
	GetMethod() string
}

type LoginRequest struct{}

func (r *LoginRequest) ToJson() []byte {
	json, _ := json.Marshal(*r)
	return json
}

func (r *LoginRequest) GetMethod() string {
	return "login"
}
