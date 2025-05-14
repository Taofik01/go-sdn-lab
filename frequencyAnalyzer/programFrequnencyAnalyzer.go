package frequencyAnalyzer

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func FrequencyAnalyzer() string {
	// This is a simple Go program that prints "Hello, World!" to the console.
	reader := bufio.NewReader(os.Stdin)
	var input string
	fmt.Println("Enter a word: ...")

	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)

	
	// Create a map to store the frequency of each character
	frequency := make(map[rune]int)

	for _, char := range input {
		frequency[char]++
		fmt.Printf("%c: %d\n", char, frequency[char])
	}

	return "Frequency analysis completed successfully"

}
