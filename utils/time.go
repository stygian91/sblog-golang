package utils

import (
	t "time"
)

func FormatTime(time t.Time) string {
	return time.In(t.UTC).Format("2006-01-02 15:04:05")
}
