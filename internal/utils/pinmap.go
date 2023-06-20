package utils

import (
	"machine"
	"strings"
)

// var AvailableButtons = map[string]machine.Pin{
// 	"D1":  machine.D1,
// 	"D2":  machine.D2,
// 	"D3":  machine.D3,
// 	"D4":  machine.D4,
// 	"D5":  machine.D5,
// 	"D6":  machine.D6,
// 	"D7":  machine.D7,
// 	"D8":  machine.D8,
// 	"D9":  machine.D9,
// 	"D10": machine.D10,
// }

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

func PinToString(p machine.Pin) string {
	switch p {
	case machine.D1:
		return "D1"
	case machine.D2:
		return "D2"
	case machine.D3:
		return "D3"
	case machine.D4:
		return "D4"
	case machine.D5:
		return "D5"
	case machine.D6:
		return "D6"
	case machine.D7:
		return "D7"
	case machine.D8:
		return "D8"
	case machine.D9:
		return "D9"
	case machine.D10:
		return "D10"
	}
	return ""
}
