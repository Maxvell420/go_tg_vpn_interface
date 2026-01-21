package Services

import (
	"time"

	"GO/app/domain/User/Models"
	"GO/app/domain/User/Repositories"
	xui "GO/app/libs/3xui"
	"GO/app/outworld"
)

type TrafficService struct {
	InboundService            *InboundService
	OutworldFacade            *outworld.OutworldFacade
	InboundRepository         *Repositories.InboundRepository
	VpnClientRepository       *Repositories.VpnClientRepository
	VpnTrafficUsageRepository *Repositories.VpnTrafficUsageRepository
}

func (s *TrafficService) HandleTrafficUsage() {
	values_inbounds := s.InboundService.GetInbounds()

	vpn_inbounds := s.OutworldFacade.GetInbounds()

	for _, vpn_inbound := range vpn_inbounds {
		values_inbound, ok := values_inbounds[vpn_inbound.Id]

		if !ok {
			continue
		}

		if values_inbound.Total == vpn_inbound.Total {
			continue
		}

		for _, vpn_client := range vpn_inbound.ClientStats {
			client, ok := values_inbound.Clients[vpn_client.Id]

			if !ok {
				continue
			}

			// Здесь посчитать остатки
			if s.needToDisableClient(client, vpn_client) {
				s.disableClient(client, vpn_client)
			}

			s.logTrafficUsage(client, vpn_client)
		}
	}
}

func (s *TrafficService) disableClient(client Models.VpnClient, vpn_client xui.ListObjClientStats) {
	// здесь будет закрытие клиента, вероятно мне нужно отправлять VO vpn обратно в либу
	// s.OutworldFacade.DisableClient(*client.GetID())
	// здесь будет обнуление остатка клиента
	to_update_client := s.VpnClientRepository.BuildModel(*client.GetID(), false, 0, 0, vpn_client.LastOnline, *client.Uuid, *client.Email, *client.InboundId, *client.UserId, *client.TimestampExpire)
	s.VpnClientRepository.Persist(to_update_client)
}

func (s *TrafficService) needToDisableClient(client Models.VpnClient, vpn_client xui.ListObjClientStats) bool {
	if client.IsUnlimited() {
		if client.IsExpired(time.Now().Unix()) {
			return true
		}
		return false

	}

	if client.IsExpired(time.Now().Unix()) {
		return true
	}

	if (vpn_client.AllTime - *client.Total) < *client.Remaining {
		return true
	}

	return false
}

func (s *TrafficService) logTrafficUsage(client Models.VpnClient, vpn_client xui.ListObjClientStats) {
	usage := vpn_client.AllTime - *client.Total
	to_log_usage := s.VpnTrafficUsageRepository.BuildModel(*client.GetID(), usage)
	s.VpnTrafficUsageRepository.Persist(to_log_usage)
}
