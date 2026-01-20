package Services

import (
	"GO/app/domain/User/Repositories"
	"GO/app/outworld"
)

type TrafficService struct {
	InboundService      *InboundService
	OutworldFacade      *outworld.OutworldFacade
	InboundRepository   *Repositories.InboundRepository
	VpnClientRepository *Repositories.VpnClientRepository
}

func (s *TrafficService) HandleTrafficUsage() {
	values_inbounds := s.InboundService.GetInbounds()

	vpn_inbounds := s.OutworldFacade.GetInbounds()

	for _, vpn_inbound := range vpn_inbounds {
		values_inbound, ok := values_inbounds[vpn_inbound.Id]

		// Здесь нужно будет сравнить значения и если они не совпадают то нужно обновить значения в базе данных или сознать новые записи
		if ok {
			values_inbound.Total = vpn_inbound.Total
		}
	}
}
