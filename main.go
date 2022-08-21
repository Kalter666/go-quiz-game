package main

import (
	"fmt"
	"time"
)

func displayScore(answers []Answer, total int) {
	score := countScore(answers)

	fmt.Printf("\nYour score is %v out of %d \n", score, total)
}

func main() {
	limit, filename := readArgs()

	problems := parseCsv(filename)
	problems = randomize(problems)

	var answers []Answer

	timer := time.NewTimer(time.Duration(limit) * time.Second)

	for _, problem := range problems {
		answerCh := readAnswer(problem.question)
		select {
		case answer := <- answerCh:
			answers = append(answers, Answer{
				problem: problem,
				givenAnswer: answer,
			})
			case <- timer.C:
				displayScore(answers, len(problems))
				return
		}
	}

	displayScore(answers, len(problems))
}