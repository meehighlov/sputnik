package events

import (
	"context"
	"log/slog"

	"github.com/meehighlov/sputnik/internal/config"
	"github.com/meehighlov/sputnik/pkg/telegram"
)

func addEventEntry(event telegram.Event) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), config.Cfg().HandlerTmeout())
	defer cancel()

	msg := "Что это за событие?"

	event.Reply(ctx, msg)

	return 2, nil
}

func addEventAccepTimestamp(event telegram.Event) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), config.Cfg().HandlerTmeout())
	defer cancel()

	event.GetChatContext().AppendText(event.GetMessage().Text)

	msg := "Введи временную метку"

	event.Reply(ctx, msg)

	return 3, nil
}

func addEventSave(event telegram.Event) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), config.Cfg().HandlerTmeout())
	defer cancel()

	timestamp := event.GetMessage().Text
	eventText := event.GetChatContext().GetTexts()[0]

	slog.Debug("timestamp", timestamp)
	slog.Debug("event text", eventText)

	msg := "Событие сохранено"
	event.Reply(ctx, msg)

	return -1, nil
}

func AddEventHandler() map[int]telegram.CommandStepHandler {
	return map[int]telegram.CommandStepHandler{
		1: addEventEntry,
		2: addEventAccepTimestamp,
		3: addEventSave,
	}
}
