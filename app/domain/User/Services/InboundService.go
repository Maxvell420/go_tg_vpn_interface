package Services

import (
	"GO/app/domain/User/Models"
	"GO/app/domain/User/Repositories"
	"GO/app/domain/User/Values"
)

type InboundService struct {
	InboundRepository   *Repositories.InboundRepository
	VpnClientRepository *Repositories.VpnClientRepository
}

func (s *InboundService) GetInbounds() map[int]values.Inbound {
	inbounds := s.InboundRepository.GetInbounds()
	clients := s.VpnClientRepository.All()

	values_inbounds := make(map[int]values.Inbound, 0)

	for _, client := range clients {
		inbound, ok := inbounds[*client.InboundId]
		if ok {

			values_inbound, ok := values_inbounds[*inbound.Id]

			if ok {
				values_inbound.Clients[*client.Id] = client
			} else {
				values_inbound = values.Inbound{
					Id:        *inbound.Id,
					Total:     *inbound.Total,
					CalcTotal: *inbound.CalcTotal,
					Protocol:  *inbound.Protocol,
					Tag:       *inbound.Tag,
					Clients:   map[int]Models.VpnClient{*client.Id: client},
				}
				values_inbounds[*inbound.Id] = values_inbound
			}

		}
	}

	return values_inbounds
}
