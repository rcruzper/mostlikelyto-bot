package bootstrap

import (
	"database/sql"
	tb "gopkg.in/tucnak/telebot.v2"
	"mostlikelyto/src/actions"
	"mostlikelyto/src/app/handler"
	"mostlikelyto/src/infrastructure/persistence/mysql/repository"
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

	appHandler := handler.New(teleBot, app.initActions())
	teleBot.Handle("/poll", appHandler.CreatePoll)

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

func (app *App) initActions() *handler.Actions {
	questionRepository := repository.NewMysqlQuestionRepository(app.conn)
	createPoll := actions.NewCreatePoll(questionRepository)

	return &handler.Actions{
		CreatePoll: createPoll,
	}
}
