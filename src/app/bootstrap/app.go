package bootstrap

import (
	"database/sql"
	tb "gopkg.in/tucnak/telebot.v2"
	"os"
	"time"
)

type App struct {
	conn *sql.DB
}

func New(conn *sql.DB) *App {
	return &App{
		conn,
	}
}

func (app *App) Start() {
	teleBot := app.initTeleBot()

	teleBot.Start()
}

func (app *App) initTeleBot() *tb.Bot {
	teleBot, err := tb.NewBot(tb.Settings{
		Token:  os.Getenv("TELEGRAM_TOKEN"),
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		panic(err.Error())
	}

	return teleBot
}
