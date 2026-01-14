package services

import (
	"GO/app/outworld"
	"GO/app/telegram/entities"
)

type TrafficService struct {
	OutworldFacade *outworld.OutworldFacade
}

func (s *TrafficService) HandleTrafficUsage(job entities.Job) {
}
