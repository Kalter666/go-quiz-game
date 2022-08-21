package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func randomize(problems []Problem) []Problem {
	for i := range problems {
    j := rand.Intn(i + 1)
    problems[i], problems[j] = problems[j], problems[i]
	}
	return problems
}

func readAnswer(question string) chan string {
	answerCh := make(chan string)
	go func() {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("%v: ", question)
		text, _ := reader.ReadString('\n')
		answerCh <- text
	}()
	return answerCh
}

func formatAnswer(s string) string {
	return strings.ToLower(strings.TrimSpace(s))
}

func isRight(answer Answer) bool {
	return formatAnswer(answer.givenAnswer) == formatAnswer(answer.problem.answer)
}

func countScore(answers []Answer) uint8 {
	var score uint8 = 0

	for _, answer := range answers {
		isCorrect := isRight(answer)
		if isCorrect {
			score += 1
		}
	}

	return score
}