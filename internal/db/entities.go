package db

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

type BaseFields struct {
	ID        string
	CreatedAt string
	UpdatedAt string
}

func (b *BaseFields) RefresTimestamps() (created string, updated string, _ error) {
	now := time.Now().Format("02.01.2006T15:04:05")
	if b.CreatedAt == "" {
		b.CreatedAt = now
	}
	b.UpdatedAt = now

	return b.CreatedAt, b.UpdatedAt, nil
}

func NewBaseFields() BaseFields {
	now := time.Now().Format("02.01.2006T15:04:05")
	return BaseFields{uuid.New().String(), now, now}
}

type User struct {
	// telegram user -> bot's user

	BaseFields

	TGId       int // id will be taken from telegram
	Name       string
	TGusername string
	ChatId     int // chatId - id of chat with user, bot uses it to send notification
}

func (user *User) GetTGUserName() string {
	if !strings.HasPrefix("@", user.TGusername) {
		return "@" + user.TGusername
	}
	return user.TGusername
}
