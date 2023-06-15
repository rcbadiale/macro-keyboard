package configs

import (
	"macro-keyboard/internal/parser"
	"macro-keyboard/internal/types"
	"strings"
	"time"
)

var BaseConfig = types.Config{
	PollingRate: time.Millisecond * 50,  // define time in between polling
	AllowRepeat: true,                   // define if the input will be repeated when held
	RepeatDelay: time.Millisecond * 500, // define time in between repeats
}

var buttons_config = `
{
	"name": "D1",
	"action_chain": [
		"mouse#0#5#",
		"delay#50ms",
		"mouse#0#-5#"
	]
}$$
{
	"name": "D2",
	"action_chain": []
}$$
{
	"name": "D3",
	"action_chain": []
}$$
{
	"name": "D4",
	"action_chain": []
}$$
{
	"name": "D5",
	"action_chain": []
}$$
{
	"name": "D6",
	"action_chain": []
}$$
{
	"name": "D7",
	"action_chain": ["keycode#KeyModifierCtrl#KeyC"]
}$$
{
	"name": "D8",
	"action_chain": ["keycode#KeyModifierCtrl#KeyV"]
}$$
{
	"name": "D9",
	"action_chain": []
}$$
{
	"name": "D10",
	"action_chain": []
}
`

var Buttons = parser.ParseConfig(strings.Split(buttons_config, "$$\n"))
