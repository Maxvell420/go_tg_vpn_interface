package services

import "GO/app/libs/3xui"

// Этот сервис отвечает за взаимодействие с xui
type XuiService struct {
	Request *xui.Request
}

func (s *XuiService) GetInbounds() {
	var req xui.VpnRequest
	req = &xui.LoginRequest{}

	s.Request.SendPost(req)
}
