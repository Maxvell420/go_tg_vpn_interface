package Services

import (
	"GO/app/domain/User/Repositories"
	"GO/app/outworld"
)

type TrafficService struct {
	VpnClientRepository *Repositories.VpnClientRepository
	OutworldFacade      *outworld.OutworldFacade
}

func (s *TrafficService) HandleTrafficUsage() {
	uuidMap, _ := s.VpnClientRepository.AllUuidMap()

	clients := s.OutworldFacade.GetInbounds()

	for _, client := range clients {
		vpnClient, ok := uuidMap[client.Uuid]
		if ok {
			vpnClient.Total = client.Total
			vpnClient.Remaining = client.Remaining
			vpnClient.LastOnline = client.LastOnline
			vpnClient.Enabled = client.Enabled
		}
	}
}
