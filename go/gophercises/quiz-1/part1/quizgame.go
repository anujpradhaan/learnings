package part1

import (
	"bufio"
	"fmt"
	"os"
	"quiz-1/io"
)

func Start(path string) {
	correctanswers, wronganswers, sequence := 0, 0, 1
	file, scanner := io.GetScannerForFile(path)

	consolescanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		question, answer := io.GetQuestionAnswerFromLine(scanner.Text())
		fmt.Println("________________")
		fmt.Printf(" %d : What is the value for question = %s ", sequence, question)
		consolescanner.Scan()
		useranswer := consolescanner.Text()
		if useranswer == answer {
			correctanswers++
		} else {
			wronganswers++
		}
		fmt.Println("________________")
		sequence++
	}
	file.Close()
	fmt.Printf("Total questions %d, Correct answers %d, and incorrect answers %d", sequence-1, correctanswers, wronganswers)
}
