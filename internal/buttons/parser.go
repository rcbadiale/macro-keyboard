package buttons

import (
	"fmt"
	"strings"
)

func ParseButton(data string, button *Button) {
	fmt.Println("Input string:", data)
	parts := strings.SplitN(data, "::", 2)
	button.Name = parts[0]
	button.ActionChain = parts[1]
	fmt.Println("Output Button:", button)
}
