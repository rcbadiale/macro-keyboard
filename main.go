package main

import (
	"machine"
	"macro-keyboard/configs"
	"macro-keyboard/internal/actions"
	btn "macro-keyboard/internal/buttons"
	"macro-keyboard/internal/storage"
	"time"
)

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

func init() {
	time.Sleep(time.Second * 2)
	filesystem := storage.New(configs.Buttons, configs.Format)
	defer filesystem.Stop()

	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for idx := range configs.Buttons {
		filesystem.ReadButton(&configs.Buttons[idx])
		time.Sleep(time.Millisecond * 100)
		configs.Buttons[idx].Pin().Configure(
			machine.PinConfig{Mode: machine.PinInputPullup},
		)
	}
}

func main() {
	ch := make(chan *btn.Button)
	go processInputs(ch)
	for {
		pollButtons(ch)
	}
}
