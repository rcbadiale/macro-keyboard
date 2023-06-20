package buttons

import (
	"fmt"
	"machine"
	"macro-keyboard/internal/utils"
	"strings"
	"time"
)

type Button struct {
	Name        string
	ActionChain []string
	LastCall    time.Time
}

func (b *Button) Pin() machine.Pin {
	return utils.PinFromString(b.Name)
}

func (b *Button) String() (out string) {
	action_chain := "[]"
	if len(b.ActionChain) > 0 {
		action_chain = fmt.Sprintf("%s", strings.Join(b.ActionChain, "$$"))
	}
	out = fmt.Sprintf("%s::%s", b.Name, action_chain)
	return out
}
