package start

import (
	"context"
	"fmt"

	"github.com/meehighlov/sputnik/internal/config"
	"github.com/meehighlov/sputnik/internal/db"
	"github.com/meehighlov/sputnik/pkg/telegram"
)

func StartHandler(event telegram.Event) error {
	ctx, cancel := context.WithTimeout(context.Background(), config.Cfg().HandlerTmeout())
	defer cancel()

	message := event.GetMessage()

	user := db.User{
		BaseFields: db.NewBaseFields(),
		Name:       message.From.FirstName,
		TGusername: message.From.Username,
		TGId:       message.From.Id,
		ChatId:     message.Chat.Id,
	}

	user.Save(ctx)

	hello := fmt.Sprintf(
		"Привет, %s 👋",
		message.From.Username,
	)

	event.Reply(ctx, hello)

	return nil
}
