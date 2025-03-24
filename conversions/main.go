package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("enter review between 1 to 5")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	// fmt.Println(reader)
	fmt.Println("You entered: ", input)
	fmt.Printf("You entered: %T", input)

	numRating , err:= strconv.ParseFloat(strings.TrimSpace(input),64)

	// fmt.Println(numRating + 1.5)
	// fmt.Println(err)

	if err != nil {
		fmt.Println(err)
	} else {
		numRating = numRating + 1.5
		fmt.Println(numRating)
		fmt.Println("Thank you for your rating")
	}
}
