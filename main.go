package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
)

type qa struct {
	question string
	answer   string
}

func main() {
	file, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal("failed to open csv file: %w", err)
	}
	reader := csv.NewReader(file)

	lines, err := reader.ReadAll()
	if err != nil {
		log.Fatal("error reading file", err)
	}

	quiz := make([]qa, 0, len(lines))
	for _, line := range lines {
		quiz = append(quiz, qa{
			question: line[0],
			answer:   line[1],
		})
	}

	timer := time.NewTimer(10 * time.Second)

	answer := ""
	count := 0
	fmt.Println("Questions:")
	for i, q := range quiz {
		fmt.Printf("#%v: %s: ", i+1, q.question)

		answerCh := make(chan string)
		go func() {
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println("")
			fmt.Println("Time is up!")
			return
		case ans := <-answerCh:
			if ans == q.answer {
				count++
			}
		}

	}

	fmt.Printf("You got %v out of %v questions correct\n", count, len(quiz))

}
