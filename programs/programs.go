package programs

import(
	"fmt"
)


func Programs() string {
	var num int

	fmt.Print("Enter a number: ")
	// Read input from the user

	_, err := fmt.Scanf("%d", &num)
	if err != nil {
		fmt.Println("Error reading input")
		return "Error occurred"
	}

	if num%2 == 0 {
		fmt.Printf("The number %d is even.\n", num)
	} else {
		fmt.Printf("The number %d is odd.\n", num)
	}

	return "Program completed successfully"
}
	