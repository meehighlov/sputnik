package reminder

import (
	"time"

	"github.com/meehighlov/sputnik/internal/config"
	"github.com/robfig/cron"
)

func MustSetupCron(cfg *config.Config) *cron.Cron {
	location, err := time.LoadLocation(cfg.Timezone)
	if err != nil {
		panic(err.Error())
	}

	c := cron.NewWithLocation(location)

	return c
}
