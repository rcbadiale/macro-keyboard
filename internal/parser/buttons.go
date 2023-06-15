package parser

import (
	"macro-keyboard/internal/models"
	"macro-keyboard/internal/types"
	"strings"
)

func ParseButton(model models.ButtonModel) types.Button {
	return types.Button{
		Pin:         AvailableButtons[strings.ToUpper(model.Name)],
		ActionChain: ParseActionChain(model.ActionChain),
	}
}
