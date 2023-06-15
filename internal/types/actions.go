package types

import (
	"fmt"
	"machine/usb/hid/keyboard"
	"machine/usb/hid/mouse"
	"macro-keyboard/internal/utils"
	"time"
)

/* Action interface, every action type must have an Execute method. */
type Action interface {
	Execute()
	String() string
}

/* Keycode action with Execute method. */
type KeycodeAction struct {
	Keycodes []keyboard.Keycode
}

func (ac *KeycodeAction) String() string {
	output := "keycode"
	for _, kc := range ac.Keycodes {
		output += fmt.Sprintf("#%v", utils.KeycodeToString(kc))
	}
	return output
}

func (ac *KeycodeAction) Execute() {
	kb := keyboard.Port()
	for _, key := range ac.Keycodes {
		kb.Down(key)
	}
	kb.Release()
}

/* Text typing action with Execute method. */
type TextAction struct {
	Text string
}

func (ac *TextAction) String() string {
	return fmt.Sprintf("text#%v", ac.Text)
}

func (ac *TextAction) Execute() {
	kb := keyboard.Port()
	kb.Write([]byte(ac.Text))
}

/* Delay action with Execute method. */
type DelayAction struct {
	Delay time.Duration
}

func (ac *DelayAction) String() string {
	return fmt.Sprintf("delay#%v", ac.Delay.String())
}

func (ac *DelayAction) Execute() {
	time.Sleep(ac.Delay)
}

/* Mouse action with Execute method. */
type MouseAction struct {
	Coordinates [2]int
	Click       mouse.Button
}

func (ac *MouseAction) String() string {
	click := ""
	switch ac.Click {
	case mouse.Left:
		click = "left"
	case mouse.Right:
		click = "right"
	case mouse.Middle:
		click = "middle"
	}
	return fmt.Sprintf("mouse#%v#%v#%v", ac.Coordinates[0], ac.Coordinates[1], click)
}

func (ac *MouseAction) Execute() {
	mouse := mouse.Port()
	mouse.Move(ac.Coordinates[0], ac.Coordinates[1])
	if ac.Click != 0 {
		mouse.Click(ac.Click)
	}
}
