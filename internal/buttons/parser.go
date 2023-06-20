package buttons

import (
	"fmt"
	"strings"
)

func ParseButton(data string, button *Button) {
	fmt.Println("Input string:", data)
	parts := strings.Split(data, "::")
	button.Name = parts[0]
	button.ActionChain = strings.Split(parts[1], "$$")
	fmt.Println("Output Button:", button)
}
