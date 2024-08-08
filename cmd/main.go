package main

import (
	"github.com/meehighlov/sputnik/internal/config"
	"github.com/meehighlov/sputnik/internal/db"
	"github.com/meehighlov/sputnik/internal/handlers/start"
	"github.com/meehighlov/sputnik/internal/handlers/events"
	"github.com/meehighlov/sputnik/internal/lib"
	"github.com/meehighlov/sputnik/pkg/telegram"
)

func main() {
	cfg := config.MustLoad()

	logger := lib.MustSetupLogging("sputnik.log", true, cfg.ENV)

	db.MustSetup("sputnik.db", logger)

	bot := telegram.NewBot(cfg.BotToken)

	bot.RegisterCommandHandler("/start", start.StartHandler)
	bot.RegisterCommandHandler("/events", events.EventsHandler)
	bot.RegisterCommandHandler("/add_event", telegram.FSM(events.AddEventHandler()))

	bot.StartPolling()
	logger.Info("Polling started.")
}
