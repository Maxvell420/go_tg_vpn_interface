package outworld

import (
	"GO/app/core"
	"GO/app/libs/3xui"
	"GO/app/libs/telegram"
	"GO/app/outworld/services"
)

type Builder struct {
	Cntx *core.Context
}

func (f *Builder) BuildOutworldManager() services.OutworldManager {
	return services.OutworldManager{TelegramBot: f.BuildTelegramBotService(), XuiService: f.BuildXuiService()}
}

func (f *Builder) BuildXuiService() *services.XuiService {
	return &services.XuiService{Request: f.BuildXuiLib()}
}

func (f *Builder) BuildTelegramBotService() *services.TelegramBot {
	return &services.TelegramBot{Lib: f.BuildTelegramLib(), KeyboardService: f.BuildKeyboardService()}
}

func (f *Builder) BuildTelegramLib() *telegram.Request {
	return f.Cntx.GetTelegramRequest()
}

func (f *Builder) BuildKeyboardService() *services.KeyboardService {
	return &services.KeyboardService{}
}

func (f *Builder) BuildXuiLib() *xui.Request {
	return f.Cntx.Get3xuiRequest()
}
