package wordandnumber

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func WordAndNumber() string {

	readers := bufio.NewReader(os.Stdin)
	var input string

	fmt.Println("Enter a word: ...")
	input, _ = readers.ReadString('\n')

	input = strings.TrimSpace(input)

	// create a map to store the frequency of each word
	frequency := make(map[string]int)

	//split the input into words
	words := strings.Fields(input)

	// iterate over each word and count its frequency

	for _, w := range words {
		w := strings.ToLower(w)
		frequency[w]++

		
	}

	// Print the frequency of each word
	fmt.Println("Word frequency: \n")

	for word, count := range frequency {
		fmt.Printf("%s: %d\n", word, count)
	}

	// Letter frequency
	fmt.Printf("\n")
	cleaned := strings.ReplaceAll(input, " ", "")
	cleaned = strings.ReplaceAll(cleaned, ",", "")

	
	fmt.Println("Letter frequency:", '\n')
	LetterFrequency := make(map[rune]int)
	for _, char := range cleaned {
		LetterFrequency[char]++
		fmt.Printf("%c: %d\n", char, LetterFrequency[char])
	}

	// return the total number of characters excluding spaces
	fmt.Printf("Total number of characters excluding spaces is %d\n", len(strings.ReplaceAll(input, " ", "")))
	fmt.Printf("Total number of words is %d\n", len(words))


	// return successfully completions

	return "Word and number analysis completed successfully"


}