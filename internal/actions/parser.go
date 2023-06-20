package actions

import (
	"machine/usb/hid/keyboard"
	"machine/usb/hid/mouse"
	"macro-keyboard/internal/utils"
	"strconv"
	"strings"
	"time"
)

func parseKeycodeAction(ac_str string) (keycodes []keyboard.Keycode) {
	if !strings.Contains(ac_str, "##") {
		return keycodes
	}
	keycodes_str := strings.Split(ac_str, "##")[1:]
	for _, k := range keycodes_str {
		keycodes = append(keycodes, utils.KeycodeFromString(k))
	}
	return keycodes
}

func parseMouseAction(ac_str string) (coordinates [2]int, click mouse.Button) {
	if !strings.Contains(ac_str, "##") {
		return coordinates, click
	}
	m := strings.Split(ac_str, "##")[1:]
	switch m[2] {
	case "left":
		click = mouse.Left
	case "right":
		click = mouse.Right
	case "middle":
		click = mouse.Middle
	}
	x, err := strconv.Atoi(m[0])
	if err != nil {
		x = 0
	}
	y, err := strconv.Atoi(m[1])
	if err != nil {
		y = 0
	}
	coordinates = [2]int{x, y}
	return coordinates, click
}

func parseTextAction(ac_str string) (text string) {
	if !strings.Contains(ac_str, "##") {
		return ""
	}
	text = strings.SplitN(ac_str, "##", 2)[1]
	return text
}

func parseDelayAction(ac_str string) (delay time.Duration, err error) {
	if !strings.Contains(ac_str, "##") {
		return delay, err
	}
	dur_str := strings.Split(ac_str, "##")[1]
	delay, err = time.ParseDuration(dur_str)
	return delay, err
}
