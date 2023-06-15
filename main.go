package main

import (
	"machine"
	"machine/usb"
	"macro-keyboard/configs"
	"macro-keyboard/internal/types"
	"strconv"
	"time"
)

var usbVID, usbPID string
var usbManufacturer, usbProduct string
var buttons = configs.Buttons
var config = configs.BaseConfig

/* Setup for HID. */
func init() {
	if usbVID != "" {
		vid, _ := strconv.ParseUint(usbVID, 0, 16)
		usb.VendorID = uint16(vid)
	}

	if usbPID != "" {
		pid, _ := strconv.ParseUint(usbPID, 0, 16)
		usb.ProductID = uint16(pid)
	}

	if usbManufacturer != "" {
		usb.Manufacturer = usbManufacturer
	}

	if usbProduct != "" {
		usb.Product = usbProduct
	}
}

func executeActionChain(actionChain []types.Action) {
	for _, action := range actionChain {
		action.Execute()
	}
}

/*
Function responsible for checking if it should execute the action chain.
*/
func processInputs(ch chan *types.Button) {
	for {
		btn := <-ch
		if time.Now().Sub(btn.LastCall) > config.RepeatDelay {
			machine.LED.Set(!machine.LED.Get())
			executeActionChain(btn.ActionChain)
			if config.AllowRepeat {
				btn.LastCall = time.Now()
			}
		}
		if !config.AllowRepeat {
			btn.LastCall = time.Now()
		}
	}
}

/*
Function responsible for polling the buttons state and placing them in the execution channel.
*/
func pollButtons(ch chan *types.Button) {
	for idx := range buttons {
		btn := &buttons[idx]
		if !btn.Pin.Get() {
			ch <- btn
		} else {
			btn.LastCall = btn.LastCall.Add(-config.RepeatDelay)
		}
	}
	time.Sleep(config.PollingRate)
}

func main() {
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for _, btn := range buttons {
		btn.Pin.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	}

	ch := make(chan *types.Button)
	go processInputs(ch)
	for {
		pollButtons(ch)
	}
}
