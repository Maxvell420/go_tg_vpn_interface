package services

import (
	"strconv"

	"GO/app/libs/telegram"
	"GO/app/outworld/types"
)

type KeyboardService struct{}

func (s *KeyboardService) GetStartKeyboard() telegram.ReplyMarkup {
	var buttons []telegram.Button
	var buttonsMap [][]telegram.Button
	text := string(telegram.Statistic)
	// Предполагаю что таким образом буду билдить callback data через функцию
	CallbackButtonData := s.buildCallbackData(types.Inbounds, 0)
	button := telegram.Button{Text: &text, Callback_data: CallbackButtonData}
	buttons = append(buttons, button)
	buttonsMap = append(buttonsMap, buttons)
	return s.buildReplyMarkUp(buttonsMap)
}

func (s *KeyboardService) buildCallbackData(action types.Command, data int) string {
	return "action:" + string(action) + ";data:" + strconv.Itoa(data)
}

func (s *KeyboardService) buildReplyMarkUp(buttonMap [][]telegram.Button) telegram.ReplyMarkup {
	return telegram.ReplyMarkup{InlineKeyboard: buttonMap}
}
