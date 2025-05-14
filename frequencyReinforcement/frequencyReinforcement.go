package frequencyreinforcement

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func FrequencyReinforcement() string {
	reader := bufio.NewReader(os.Stdin)
	var input string
	fmt.Println("Enter a word: ...")

	input, _ = reader.ReadString('\n')

	input = strings.TrimSpace(input)

	words := strings.Fields(input)

	frequency := make(map[string]int)

	for _, w := range words {
		w := strings.ToLower(w)
		frequency[w]++
	}

	// Print the frequency of each word
	fmt.Printf("Number of unique word is %d \n", len(frequency))
	fmt.Printf("Total number of word is %d \n", len(words))

	var maxWord string
	maxCount := 0

	for word, count := range frequency {
		if count > maxCount {
			maxCount = count
			maxWord = word
		}	
	}

	if maxCount > 1 {
			fmt.Printf("The most frequent word is %s: %d\n " , maxWord, maxCount)
		
		} else {
			fmt.Printf("All words are unique\n")
		}
	
		return "Frequency reinforcement completed."

}
