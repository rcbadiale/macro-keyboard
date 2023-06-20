package configs

import (
	"macro-keyboard/internal/buttons"
	"time"
)

var Format = false

type Config struct {
	PollingDelay time.Duration
	AllowRepeat  bool
	RepeatDelay  time.Duration
}

var BaseConfig = Config{
	PollingDelay: time.Millisecond * 50,  // define time in between polling
	AllowRepeat:  true,                   // define if the input will be repeated when held
	RepeatDelay:  time.Millisecond * 500, // define time in between repeats
}

var Buttons = []buttons.Button{
	buttons.Button{
		Name:        "D1",
		ActionChain: []string{"mouse##5##0##", "delay##50ms", "mouse##-5##0##"},
	},
	buttons.Button{
		Name:        "D2",
		ActionChain: []string{},
	},
	buttons.Button{
		Name:        "D3",
		ActionChain: []string{},
	},
	buttons.Button{
		Name:        "D4",
		ActionChain: []string{},
	},
	buttons.Button{
		Name:        "D5",
		ActionChain: []string{},
	},
	buttons.Button{
		Name:        "D6",
		ActionChain: []string{},
	},
	buttons.Button{
		Name:        "D7",
		ActionChain: []string{"keycode##KeyModifierCtrl##KeyC"},
	},
	buttons.Button{
		Name:        "D8",
		ActionChain: []string{"keycode##KeyModifierCtrl##KeyV"},
	},
	buttons.Button{
		Name:        "D9",
		ActionChain: []string{},
	},
	buttons.Button{
		Name:        "D10",
		ActionChain: []string{},
	},
}
