package sentenceanalyzer

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func SentenceAnalyzer() string {
	reader := bufio.NewReader(os.Stdin)
	var input string
	fmt.Println("Enter a sentence: ...!")
	input, _ = reader.ReadString('\n')

	input = input[:len(input)-1]

	frequency := make(map[string]int)

	word := strings.Fields(input)

	for _, word := range word {
		word := strings.ToLower(string(word))
		frequency[word]++
	}

	for word, count := range frequency {
		fmt.Printf("%s: %d\n", word, count)
	}

	return "Word frequency analysis completed successfully!"

}
