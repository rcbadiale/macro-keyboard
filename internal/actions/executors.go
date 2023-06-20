package actions

import (
	"machine/usb/hid/keyboard"
	"machine/usb/hid/mouse"
	"strings"
	"time"
)

func ExecuteActionChain(input []string) {
	for _, ac_str := range input {
		switch {
		case strings.HasPrefix(ac_str, "keycode"):
			ExecuteKeycodeAction(ac_str)
		case strings.HasPrefix(ac_str, "mouse"):
			ExecuteMouseAction(ac_str)
		case strings.HasPrefix(ac_str, "text"):
			ExecuteTextAction(ac_str)
		case strings.HasPrefix(ac_str, "delay"):
			ExecuteDelayAction(ac_str)
		}
	}
}

func ExecuteKeycodeAction(action string) {
	keycodes := parseKeycodeAction(action)
	kb := keyboard.Port()
	for _, key := range keycodes {
		kb.Down(key)
	}
	kb.Release()
}

func ExecuteTextAction(action string) {
	text := parseTextAction(action)
	kb := keyboard.Port()
	kb.Write([]byte(text))
}

func ExecuteDelayAction(action string) {
	delay, err := parseDelayAction(action)
	if err != nil {
		return
	}
	time.Sleep(delay)
}

func ExecuteMouseAction(action string) {
	coordinates, click := parseMouseAction(action)
	mouse := mouse.Port()
	mouse.Move(coordinates[0], coordinates[1])
	if click != 0 {
		mouse.Click(click)
	}
}
