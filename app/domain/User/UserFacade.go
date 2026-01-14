package user

import "GO/app/domain/User/Models"

type UserFacade struct {
	Builder *UserBuilder
}

func (f *UserFacade) HandleStartReferal(tg_user_id int, referal_link_hash string) {
	service := f.Builder.BuildReferalService()
	service.HandleStartReferal(tg_user_id, referal_link_hash)
}

func (f *UserFacade) GetUserRefLink(tg_user_id int) string {
	service := f.Builder.BuildReferalService()
	return service.GetUserRefLink(tg_user_id)
}

func (f *UserFacade) GetUserByTgId(tg_user_id int) (Models.UserModel, error) {
	repository := f.Builder.BuildUserRepository()
	return repository.GetByTgID(tg_user_id)
}
