package utils

import (
	"bufio"
	"fmt"
	"os"
)

func GetInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return input[:len(input)-2]
}
