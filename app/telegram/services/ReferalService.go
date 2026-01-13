package services

import (
	"crypto/sha256"
	"encoding/base64"
	"os"
	"strconv"

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

func (s *ReferalService) generateLink(tg_user_id int) string {
	hash := s.generateUserHash(tg_user_id)
	link := s.ReferalLinkRepository.BuildModel(hash, tg_user_id)
	link, _ = s.ReferalLinkRepository.Persist(link)

	return link.GetHash()
}

func (s *ReferalService) generateUserHash(tg_user_id int) string {
	salt := os.Getenv("REF_LINK_SALT")
	data := strconv.FormatInt(int64(tg_user_id), 10) + salt

	hash := sha256.Sum256([]byte(data))

	encoded := base64.URLEncoding.EncodeToString(hash[:])

	if len(encoded) > 20 {
		encoded = encoded[:20]
	}

	return encoded
}

func (s *ReferalService) GetUserRefLink(tg_user_id int) string {
	link, err := s.ReferalLinkRepository.GetByTgId(tg_user_id)
	if err == nil {
		return link.GetHash()
	} else {
		return s.generateLink(tg_user_id)
	}
}
