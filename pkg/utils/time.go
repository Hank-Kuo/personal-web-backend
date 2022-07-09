package utils

import (
	"time"

	"github.com/Hank-Kuo/personal-web-backend/config"
)

func GetCurrentTime(cfg *config.Config) time.Time {
	t := time.Now()

	utc, err := time.LoadLocation(cfg.Server.Timezone)
	if err != nil {
		return t
	}

	return t.In(utc)
}
