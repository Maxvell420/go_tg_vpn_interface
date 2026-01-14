package user

import (
	"GO/app/core"
	"GO/app/domain/User/Repositories"
	"GO/app/domain/User/Services"
)

type UserBuilder struct {
	Cntx *core.Context
}

func (b *UserBuilder) BuildUserRepository() *Repositories.UserRepository {
	return &Repositories.UserRepository{Db: b.Cntx.GetDb()}
}

func (b *UserBuilder) BuildReferalLinkRepository() *Repositories.ReferalLinkRepository {
	return &Repositories.ReferalLinkRepository{Db: b.Cntx.GetDb()}
}

func (b *UserBuilder) BuildReferalUserRepository() *Repositories.ReferalUserRepository {
	return &Repositories.ReferalUserRepository{Db: b.Cntx.GetDb()}
}

func (b *UserBuilder) BuildReferalService() *Services.ReferalService {
	return &Services.ReferalService{UserRepository: b.BuildUserRepository(), ReferalLinkRepository: b.BuildReferalLinkRepository(), ReferalUserRepository: b.BuildReferalUserRepository()}
}
