package types

import (
	"machine"
	"time"
)

type Button struct {
	Pin         machine.Pin
	ActionChain []Action
	LastCall    time.Time
}
