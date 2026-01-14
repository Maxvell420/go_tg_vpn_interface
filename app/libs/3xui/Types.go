package xui

import (
	"encoding/json"
)

type VpnRequest interface {
	ToJson() []byte
	GetMethod() XuiPath
}

type XuiRequest struct {
	Method XuiPath
}

func (r *XuiRequest) ToJson() []byte {
	json, _ := json.Marshal(*r)
	return json
}

func (r *XuiRequest) GetMethod() XuiPath {
	return r.Method
}

type VpnResponse interface{}

type XuiPath string

const (
	Inbounds XuiPath = "panel/api/inbounds/list"
)

type ListResponse struct {
	Success bool
	Obj     []ListObj `json:"obj"`
}

type ListObj struct {
	Id          int                  `json:"id"`
	Up          int                  `json:"up"`
	Down        int                  `json:"down"`
	Total       int                  `json:"total"`
	AllTime     int                  `json:"allTime"`
	Protocol    string               `json:"protocol"`
	Tag         string               `json:"tag"`
	ClientStats []ListObjClientStats `json:"clientStats"`
}

type ListObjClientStats struct {
	Email      string `json:"email"`
	Up         int    `json:"up"`
	Down       int    `json:"down"`
	AllTime    int    `json:"allTime"`
	LastOnline int    `json:"lastOnline"`
	Uuid       string `json:"uuid"`
	Total      int    `json:"total"`
	Id         int    `json:"id"`
	Enable     bool   `json:"enable"`
}
