package fix

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func (f *Fix) fixPrompt() error {
	fmt.Println(os.Args)
	// Create a scanner to read from standard input
	scanner := bufio.NewScanner(os.Stdin)

	// Use a StringBuilder to concatenate all input lines
	var builder strings.Builder

	for scanner.Scan() {
		builder.WriteString(scanner.Text())
		builder.WriteString("\n")
	}

	// Get the complete input as a string
	pipedInput := builder.String()
	fmt.Println(pipedInput)

	fmt.Println("fix -> action")

	return nil
}
