package structures

import (
	"machine"
	"machine/usb/hid/keyboard"
	"machine/usb/hid/mouse"
	"time"
)

type Config struct {
	PollingRate time.Duration
	AllowRepeat bool
	RepeatDelay time.Duration
}

type Button struct {
	Pin         machine.Pin
	ActionChain []Action
	LastCall    time.Time
}

/* Action interface, every action type must have an Execute method. */
type Action interface {
	Execute()
}

/* Keycode action with Execute method. */
type KeycodeAction struct {
	Keycodes []keyboard.Keycode
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

func (ac *TextAction) Execute() {
	kb := keyboard.Port()
	kb.Write([]byte(ac.Text))
}

/* Delay action with Execute method. */
type DelayAction struct {
	Delay time.Duration
}

func (ac *DelayAction) Execute() {
	time.Sleep(ac.Delay)
}

/* Mouse action with Execute method. */
type MouseAction struct {
	Coordinates [2]int
	Click       mouse.Button
}

func (ac *MouseAction) Execute() {
	mouse := mouse.Port()
	mouse.Move(ac.Coordinates[0], ac.Coordinates[1])
	if ac.Click != 0 {
		mouse.Click(ac.Click)
	}
}
