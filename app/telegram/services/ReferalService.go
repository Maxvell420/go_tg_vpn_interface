package services

import (
	"GO/app/domain/User/Repositories"
)

type ReferalService struct {
	ReferalLinkRepository *Repositories.ReferalLinkRepository
	ReferalUserRepository *Repositories.ReferalUserRepository
	UserRepository        *Repositories.UserRepository
}

func (s *ReferalService) HandleStartReferal(tg_user_id int, referal_link_hash string) {
	link, err := s.ReferalLinkRepository.GetByHash(referal_link_hash)
	if err != nil {
		return
	}

	referal_user, err := s.ReferalUserRepository.GetByTgId(link.GetTgId())
	if err == nil && *referal_user.GetID() == tg_user_id {
		return
	}

	user, err := s.UserRepository.GetByTgID(link.GetTgId())
	if err == nil {
		model := s.ReferalUserRepository.BuildModel(*user.GetTgId(), link.GetTgId())
		_, err = s.ReferalUserRepository.Persist(model)
	}
}
