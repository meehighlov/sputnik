package config

import (
	"log"
	"os"

	"strings"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	ENV                   string `env:"ENV" env-default:"local"`
	BotToken              string `env:"BOT_TOKEN" env-required:"true"`
	Users                 string `env:"USERS" env-required:"true"`
	ReportChatId          string `env:"REPORT_CHAT_ID" env-required:"true"`
	HandlerExecTimeoutSec int    `env:"HANDLER_EXEC_TIMEOUT_SEC" env-default:"2"`
	Timezone              string `env:"TIMEZONE" env-default:"Europe/Moscow"`

	loaded bool `env-default:"false"`
}

func (cfg *Config) AuthList() []string {
	return strings.Split(cfg.Users, ",")
}

func (cfg *Config) HandlerTmeout() time.Duration {
	return time.Duration(cfg.HandlerExecTimeoutSec) * time.Second
}

var cfg Config

// loads config from .env
// panics on any read error
// also sets TZ env variable from according .env value
func MustLoad() *Config {
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		log.Fatal("Not found .env file")
	}

	err := cleanenv.ReadConfig(".env", &cfg)
	if err != nil {
		log.Fatal("Failed to read envs:", err.Error())
	}

	os.Setenv("TZ", cfg.Timezone)

	cfg.loaded = true

	return &cfg
}

func Cfg() *Config {
	if !cfg.loaded {
		panic("Accessing not loaded config. Exiting.")
	}

	return &cfg
}
