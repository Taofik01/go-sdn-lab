package wordclassifier

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

const wordCountFormat = "%s: %d "

func WordClassifier() string {
	// This function classifies words based on their length in alphabet, numeric or alphanumeric and also gives the count of each type

	var input string
	var alphabetCount, numericCount, alphanumericCount = 0, 0, 0
	fmt.Println("Enter a sentence to classify: ...")
	reader := bufio.NewReader(os.Stdin)
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)

	sentence := strings.Fields(input);

	frequency := make(map[string]int);

	// alphamatched, _ := regexp.MatchString("^[a-zA-Z]+$", input);
	// numericmatched, _ := regexp.MatchString("^[0-9]+$", input);
	// alphanumericmatched, _ := regexp.MatchString("^[a-zA-Z0-9]+$", input);

	var words []string
	var nums []string
	var alphas []string

	for _, w := range sentence {
		w = strings.ToLower(w)
		frequency[w]++



		if matched, _ := regexp.MatchString("^[a-zA-Z]+$", w); matched && frequency[w] == 1 {
			alphabetCount++
			// storing the alphabetic words in a map
			word := fmt.Sprintf(w)

			
			words = append(words, word)
			

		} else if matched, _ := regexp.MatchString("^[0-9]+$", w); matched && frequency[w] == 1 {

			numericCount++
			// storing the numeric words in a map
			num := fmt.Sprintf(w)
			nums = append(nums, num)
		} else if matched, _ := regexp.MatchString("^[a-zA-Z0-9]+$", w); matched && frequency[w] == 1 {
			alphanumericCount++
			// storing the alphanumeric words in a map
			alpha := fmt.Sprintf(w)
			alphas = append(alphas, alpha)
		}
	}

	// printing the numbers and the mapped words together for alphabetic words
	fmt.Printf("Alphabetic words (%d): [%s] \n", alphabetCount, strings.Join(words, ", "))
	fmt.Printf("Numeric words (%d): [%s] \n", numericCount, strings.Join(nums, ", "))
	fmt.Printf("Alphanumeric words (%d): [%s] \n", alphanumericCount, strings.Join(alphas, ", "))
	




	return "Word classification completed successfully"
	

 }
