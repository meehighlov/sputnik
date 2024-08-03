package reminder

import (
	"context"

	"github.com/meehighlov/sputnik/internal/config"
	"github.com/meehighlov/sputnik/pkg/telegram"
)

func addReminderEntry(event telegram.Event) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), config.Cfg().HandlerTmeout())
	defer cancel()

	msg := "Введи текст напоминания"

	event.Reply(ctx, msg)

	return 2, nil
}

func addReminderAccepTimestamp(event telegram.Event) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), config.Cfg().HandlerTmeout())
	defer cancel()

	event.GetChatContext().AppendText(event.GetMessage().Text)

	msg := `
	Когда напомнить?

	Формат временной метки: dd.mm.yyyy [hh.mm] [d,w,m,y]
	Часы и минуты - опционально
	`

	event.Reply(ctx, msg)

	return 3, nil
}

func AddReminderHandler() map[int]telegram.CommandStepHandler {
	return map[int]telegram.CommandStepHandler{
		1: addReminderEntry,
		2: addReminderAccepTimestamp,
	}
}
