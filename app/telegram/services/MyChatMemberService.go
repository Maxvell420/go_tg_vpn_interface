package services

import (
	"GO/app/telegram/repositories"
	"GO/app/telegram/updates"
)

type MyChatMemberService struct {
	UserRepository repositories.UserRepository
}

func (s *MyChatMemberService) HandleMyChatMemberUpdate(update updates.MyChatMember) {
	user_id := update.GetUser()
	user, err := s.UserRepository.GetByTgID(user_id)
	if err != nil {
		user = user.FromData(update.GetUser(), update.From.Username, "no", "no")
	}

	user.UpdateStatus(update.GetNewStatus())

	s.UserRepository.Persist(user)
}
