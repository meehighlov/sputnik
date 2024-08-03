package reminder

import (
	"context"
	"fmt"
	"strings"

	"github.com/meehighlov/sputnik/internal/config"
	"github.com/meehighlov/sputnik/pkg/telegram"
)

func ReminderHandler(event telegram.Event) error {
	ctx, cancel := context.WithTimeout(context.Background(), config.Cfg().HandlerTmeout())
	defer cancel()

	msgTemplate := "Выбери опцию:\n\n%s"

	options := []string{
		"/add_reminder",
		"/remove_reminder",
		"/list_reminder",
	}

	msg := fmt.Sprintf(msgTemplate, strings.Join(options, "\n"))

	event.Reply(ctx, msg)

	return nil
}
