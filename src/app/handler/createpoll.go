package handler

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"time"
)

func (h *Handler) CreatePoll(m *tb.Message) {
	poll, _ := h.actions.CreatePoll.Execute()

	telegramPoll := &tb.Poll{
		Type:          tb.PollRegular,
		Question:      poll.Question,
		CloseUnixdate: time.Now().Unix() + 300,
		Anonymous: false,
	}
	telegramPoll.AddOptions("Dani", "Luis", "Julian", "Silvia", "Álvaro", "Alberto", "Raúl")

	_, _ = h.teleBot.Send(m.Chat, telegramPoll)
}
