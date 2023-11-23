package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

// set variables for correct responses and total responses that we'll print at the end
var correctCount = 0
var totalCount = 0

// function that calculates the score and percentage at the end of the quiz
func score() {
	if totalCount > 0 {
		percentage := float64(correctCount) / float64(totalCount) * 100
		fmt.Printf("You got %d out of %d questions correct. %.2f%%\n", correctCount, totalCount, percentage)
	}
}

func main() {
	// set the csv file into a variable and open it
	file, err := os.Open("quiz.csv")
	if err != nil {
		panic(err)
	}
	// close the file
	defer file.Close()

	// create a reader for the file and read it
	reader := csv.NewReader(file)
	// create a scanner to scan userInput
	scanner := bufio.NewScanner(os.Stdin)

	// read each line and break when we reach the end
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}
	// separate the columns in the csv file into variables
		question := record[0]
		answer := record[1]

		fmt.Println("Question:", question)
		fmt.Print("Your answer: ")

	// read user input with the scanner
		scanner.Scan()
		userInput := scanner.Text()

		// compare the scanned user input to the correct answer
		if strings.TrimSpace(strings.ToLower(userInput)) == strings.ToLower(answer) {
			fmt.Println("Correct!")
			// increment the correct answer count
			correctCount++
		} else {
			fmt.Println("Nope! The correct answer is:", answer)
		}
		// increment the total count
		totalCount++
	}
	// after the last question, output number of correct answers / number of total question and percentage
	score()
}