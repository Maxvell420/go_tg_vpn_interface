package services

import (
	"GO/app/domain/User/Models"
	"GO/app/domain/User/Repositories"
	"GO/app/telegram/updates"
)

type MyChatMemberService struct {
	UserRepository Repositories.UserRepository
}

func (s *MyChatMemberService) HandleMyChatMemberUpdate(update updates.MyChatMember) {
	user_id := update.GetUser()
	user, err := s.UserRepository.GetByTgID(user_id)
	if err != nil {
		user = &Models.User{Tg_id: &user_id, User_name: update.From.Username, Kicked: "no", Is_admin: "no"}
	}

	user.UpdateStatus(update.GetNewStatus())

	s.UserRepository.Persist(user)
}
