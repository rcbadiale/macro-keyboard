package utils

import (
	"machine"
	"strings"
)

func PinFromString(s string) machine.Pin {
	switch strings.ToUpper(s) {
	case "D1":
		return machine.D1
	case "D2":
		return machine.D2
	case "D3":
		return machine.D3
	case "D4":
		return machine.D4
	case "D5":
		return machine.D5
	case "D6":
		return machine.D6
	case "D7":
		return machine.D7
	case "D8":
		return machine.D8
	case "D9":
		return machine.D9
	case "D10":
		return machine.D10
	}
	return 0
}
