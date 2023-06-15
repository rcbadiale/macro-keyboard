package configs

import (
	"machine"
	"machine/usb/hid/keyboard"
	"macro-keyboard/internal/structures"
	"time"
)

var BaseConfig = structures.Config{
	PollingRate: time.Millisecond * 50,  // define time in between polling
	AllowRepeat: true,                   // define if the input will be repeated when held
	RepeatDelay: time.Millisecond * 500, // define time in between repeats
}

var Buttons = []structures.Button{
	structures.Button{
		Pin: machine.D1,
		ActionChain: []structures.Action{
			&structures.MouseAction{Coordinates: [2]int{5, 0}},
			&structures.DelayAction{Delay: time.Millisecond * 50},
			&structures.MouseAction{Coordinates: [2]int{-5, 0}},
		},
		LastCall: time.Now(),
	},
	structures.Button{
		Pin:         machine.D2,
		ActionChain: []structures.Action{},
		LastCall:    time.Now(),
	},
	structures.Button{
		Pin:         machine.D3,
		ActionChain: []structures.Action{},
		LastCall:    time.Now(),
	},
	structures.Button{
		Pin:         machine.D4,
		ActionChain: []structures.Action{},
		LastCall:    time.Now(),
	},
	structures.Button{
		Pin:         machine.D5,
		ActionChain: []structures.Action{},
		LastCall:    time.Now(),
	},
	structures.Button{
		Pin:         machine.D6,
		ActionChain: []structures.Action{},
		LastCall:    time.Now(),
	},
	structures.Button{
		Pin: machine.D7,
		ActionChain: []structures.Action{
			&structures.KeycodeAction{Keycodes: []keyboard.Keycode{keyboard.KeyModifierCtrl, keyboard.KeyC}},
		},
		LastCall: time.Now(),
	},
	structures.Button{
		Pin: machine.D8,
		ActionChain: []structures.Action{
			&structures.KeycodeAction{Keycodes: []keyboard.Keycode{keyboard.KeyModifierCtrl, keyboard.KeyV}},
		},
		LastCall: time.Now(),
	},
	structures.Button{
		Pin:         machine.D9,
		ActionChain: []structures.Action{},
		LastCall:    time.Now(),
	},
	structures.Button{
		Pin:         machine.D10,
		ActionChain: []structures.Action{},
		LastCall:    time.Now(),
	},
}
