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
			executeKeycodeAction(ac_str)
		case strings.HasPrefix(ac_str, "mouse"):
			executeMouseAction(ac_str)
		case strings.HasPrefix(ac_str, "text"):
			executeTextAction(ac_str)
		case strings.HasPrefix(ac_str, "delay"):
			executeDelayAction(ac_str)
		}
	}
}

func executeKeycodeAction(action string) {
	keycodes := parseKeycodeAction(action)
	kb := keyboard.Port()
	for _, key := range keycodes {
		kb.Down(key)
	}
	kb.Release()
}

func executeTextAction(action string) {
	text := parseTextAction(action)
	kb := keyboard.Port()
	kb.Write([]byte(text))
}

func executeDelayAction(action string) {
	delay, err := parseDelayAction(action)
	if err != nil {
		return
	}
	time.Sleep(delay)
}

func executeMouseAction(action string) {
	coordinates, click := parseMouseAction(action)
	mouse := mouse.Port()
	mouse.Move(coordinates[0], coordinates[1])
	if click != 0 {
		mouse.Click(click)
	}
}
