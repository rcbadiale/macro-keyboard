package parser

import (
	"encoding/json"
	"fmt"
	"machine"
	"macro-keyboard/internal/models"
	"macro-keyboard/internal/types"
)

var AvailableButtons = map[string]machine.Pin{
	"D1":  machine.D1,
	"D2":  machine.D2,
	"D3":  machine.D3,
	"D4":  machine.D4,
	"D5":  machine.D5,
	"D6":  machine.D6,
	"D7":  machine.D7,
	"D8":  machine.D8,
	"D9":  machine.D9,
	"D10": machine.D10,
}

func ParseConfig(data []string) (output []types.Button) {
	var button models.ButtonModel
	for _, d := range data {
		err := json.Unmarshal([]byte(d), &button)
		if err == nil {
			fmt.Println(button)
		} else {
			fmt.Printf("Unmarshal err:", err)
			return
		}
		output = append(output, ParseButton(button))
	}
	return output
}
