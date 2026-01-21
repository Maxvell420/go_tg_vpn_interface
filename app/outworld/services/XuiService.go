package services

import (
	"GO/app/libs/3xui"
)

// Этот сервис отвечает за взаимодействие с xui
type XuiService struct {
	Request *xui.Request
}

func (s *XuiService) GetInbounds() []xui.ListObj {
	var req xui.VpnRequest
	req = &xui.XuiRequest{Method: xui.Inbounds}

	body := s.Request.SendPost(req)
	response := body.(xui.ListResponse)
	return response.Obj
}

func (s *XuiService) DisableClient() {
	// закрывать клиента
	// var req xui.VpnRequest
	// req = &xui.XuiRequest{Method: xui.DisableClient, ClientId: clientId}

	// body := s.Request.SendPost(req)
	// response := body.(xui.DisableClientResponse)
	// return response.Success
}
