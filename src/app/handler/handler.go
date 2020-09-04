package handler

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"mostlikelyto/src/actions"
)

type Actions struct {
	CreatePoll *actions.CreatePoll
}

type Handler struct {
	teleBot *tb.Bot
	actions *Actions
}

func New(teleBot *tb.Bot, actions *Actions) *Handler {
	return &Handler{
		teleBot,
		actions,
	}
}
