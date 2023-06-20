package buttons

import (
	"fmt"
	"machine"
	"macro-keyboard/internal/utils"
	"time"
)

type Button struct {
	Name        string
	ActionChain string
	LastCall    time.Time
}

func (b *Button) Pin() machine.Pin {
	return utils.PinFromString(b.Name)
}

func (b *Button) String() (out string) {
	out = fmt.Sprintf("%s::%s", b.Name, b.ActionChain)
	return out
}
