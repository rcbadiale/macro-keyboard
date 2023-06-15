package parser

import (
	"log"
	"machine/usb/hid/keyboard"
	"machine/usb/hid/mouse"
	"macro-keyboard/internal/types"
	"macro-keyboard/internal/utils"
	"strconv"
	"strings"
	"time"
)

func ParseActionChain(input []string) []types.Action {
	var output []types.Action
	for _, ac_str := range input {
		switch {
		case strings.HasPrefix(ac_str, "keycode"):
			output = append(output, parseKeycodeAction(ac_str))
		case strings.HasPrefix(ac_str, "mouse"):
			action, err := parseMouseAction(ac_str)
			if err != nil {
				log.Println("unable to parse mouse action")
				continue
			}
			output = append(output, action)
		case strings.HasPrefix(ac_str, "text"):
			output = append(output, parseTextAction(ac_str))
		case strings.HasPrefix(ac_str, "delay"):
			action, err := parseDelayAction(ac_str)
			if err != nil {
				log.Println("unable to parse delay action")
				continue
			}
			output = append(output, action)
		}
	}
	return output
}

func parseKeycodeAction(ac_str string) *types.KeycodeAction {
	keycodes_str := strings.Split(ac_str, "#")[1:]
	var keycodes []keyboard.Keycode
	for _, k := range keycodes_str {
		keycodes = append(keycodes, utils.KeycodeFromString(k))
	}
	return &types.KeycodeAction{Keycodes: keycodes}
}

func parseMouseAction(ac_str string) (*types.MouseAction, error) {
	m := strings.Split(ac_str, "#")[1:]
	var click mouse.Button
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
		return nil, err
	}
	y, err := strconv.Atoi(m[1])
	if err != nil {
		return nil, err
	}
	return &types.MouseAction{Coordinates: [2]int{x, y}, Click: click}, nil
}

func parseTextAction(ac_str string) *types.TextAction {
	text := strings.Split(ac_str, "#")[1]
	return &types.TextAction{Text: text}
}

func parseDelayAction(ac_str string) (*types.DelayAction, error) {
	dur_str := strings.Split(ac_str, "#")[1]
	d, err := time.ParseDuration(dur_str)
	if err != nil {
		return nil, err
	}
	return &types.DelayAction{Delay: d}, nil
}
