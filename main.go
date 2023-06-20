package main

import (
	"fmt"
	"machine"
	"macro-keyboard/configs"
	"macro-keyboard/internal/actions"
	btn "macro-keyboard/internal/buttons"
	"macro-keyboard/internal/storage"
	"strings"
	"time"
)

/*
Configure all pins and read persistent storage.
*/
func init() {
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	configs.ResetPin.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	time.Sleep(time.Second * 2)
	configs.Filesystem = storage.New(configs.Buttons, configs.Format || !configs.ResetPin.Get())

	for idx := range configs.Buttons {
		configs.Filesystem.ReadButton(&configs.Buttons[idx])
		time.Sleep(time.Millisecond * 100)
		configs.Buttons[idx].Pin().Configure(
			machine.PinConfig{Mode: machine.PinInputPullup},
		)
	}
}

/*
Function responsible for checking if it should execute the action chain.
*/
func processInputs(ch chan *btn.Button) {
	for {
		b := <-ch
		if time.Now().Sub(b.LastCall) > configs.BaseConfig.RepeatDelay {
			machine.LED.Set(!machine.LED.Get())
			actions.ExecuteActionChain(b.ActionChain)
			if configs.BaseConfig.AllowRepeat {
				b.LastCall = time.Now()
			}
		}
		if !configs.BaseConfig.AllowRepeat {
			b.LastCall = time.Now()
		}
	}
}

/*
Function responsible for polling the buttons state and placing them in the execution channel.
*/
func pollButtons(ch chan *btn.Button) {
	for idx := range configs.Buttons {
		if !(&configs.Buttons[idx]).Pin().Get() {
			ch <- (&configs.Buttons[idx])
		} else {
			(&configs.Buttons[idx]).LastCall = (&configs.Buttons[idx]).LastCall.Add(-configs.BaseConfig.RepeatDelay)
		}
	}
	time.Sleep(configs.BaseConfig.PollingDelay)
}

/*
Function responsible for executing a command when it is sent through serial port.
*/
func runCommand(s string) {
	if !strings.HasPrefix(s, "add ") {
		fmt.Println("\nInvalid command.")
		return
	}
	if !strings.Contains(s, "::") {
		fmt.Println("\nInvalid syntax!\n  Should be: `add <name>::<actions>`")
		return
	}
	btn_data := strings.SplitN(s[4:], "::", 2)
	for idx := range configs.Buttons {
		if btn_data[0] == configs.Buttons[idx].Name {
			configs.Buttons[idx].ActionChain = btn_data[1]
			configs.Filesystem.WriteButton(&configs.Buttons[idx])
		}
	}
}

func main() {
	ch := make(chan *btn.Button)
	cmd := ""
	go processInputs(ch)
	for {
		// process console inputs before polling buttons
		if configs.Console.Buffered() > 0 {
			data, _ := configs.Console.ReadByte()
			switch data {
			case 8: // backspace
				if len(cmd) > 0 {
					cmd = cmd[:len(cmd)-1]
					configs.Console.Write([]byte{0x8, 0x20, 0x8})
				}
			case 13: // enter
				runCommand(cmd)
				cmd = ""
			default: // any other char
				fmt.Print(string(data))
				cmd += string(data)
			}
		}
		pollButtons(ch)
	}
}
