package events

import (
	"context"
	"fmt"
	"strings"

	"github.com/meehighlov/sputnik/internal/config"
	"github.com/meehighlov/sputnik/pkg/telegram"
)

func EventsHandler(event telegram.Event) error {
	ctx, cancel := context.WithTimeout(context.Background(), config.Cfg().HandlerTmeout())
	defer cancel()

	msgTemplate := "Выбери опцию\n\n%s"

	options := []string{
		"/add_event",
		"/remove_event",
		"/list_event",
	}

	msg := fmt.Sprintf(msgTemplate, strings.Join(options, "\n"))

	event.Reply(ctx, msg)

	return nil
}
