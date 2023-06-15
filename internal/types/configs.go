package types

import (
	"time"
)

type Config struct {
	PollingRate time.Duration
	AllowRepeat bool
	RepeatDelay time.Duration
	Buttons     []Button
}
