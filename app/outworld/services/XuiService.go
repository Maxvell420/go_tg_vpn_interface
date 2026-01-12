package services

import (
	"GO/app/libs/3xui"
	// "github.com/davecgh/go-spew/spew"
)

// Этот сервис отвечает за взаимодействие с xui
type XuiService struct {
	Request *xui.Request
}

func (s *XuiService) GetInbounds() {
	// var req xui.VpnRequest
	// req = &xui.XuiRequest{Method: xui.Inbounds}

	// body := s.Request.SendPost(req)
	// response := body.(xui.ListResponse)
	// spew.Dump(response)
}
