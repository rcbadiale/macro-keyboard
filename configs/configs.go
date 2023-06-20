package configs

import (
	"machine"
	"macro-keyboard/internal/buttons"
	"macro-keyboard/internal/storage"
	"time"
)

var Format = false
var ResetPin = machine.D2
var Console = machine.Serial
var Filesystem storage.Flash

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
		ActionChain: "mouse##5##0##$$delay##50ms$$mouse##-5##0##",
	},
	buttons.Button{
		Name:        "D2",
		ActionChain: "",
	},
	buttons.Button{
		Name:        "D3",
		ActionChain: "",
	},
	buttons.Button{
		Name:        "D4",
		ActionChain: "",
	},
	buttons.Button{
		Name:        "D5",
		ActionChain: "",
	},
	buttons.Button{
		Name:        "D6",
		ActionChain: "",
	},
	buttons.Button{
		Name:        "D7",
		ActionChain: "keycode##KeyModifierCtrl##KeyC",
	},
	buttons.Button{
		Name:        "D8",
		ActionChain: "keycode##KeyModifierCtrl##KeyV",
	},
	buttons.Button{
		Name:        "D9",
		ActionChain: "",
	},
	buttons.Button{
		Name:        "D10",
		ActionChain: "",
	},
}
