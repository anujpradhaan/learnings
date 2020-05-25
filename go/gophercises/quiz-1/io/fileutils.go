package io

import (
	"bufio"
	"encoding/csv"
	"os"
	"quiz-1/errorhandler"
	"strings"
)

func GetScannerForFile(path string) (*os.File, *bufio.Scanner) {
	file, err := os.Open(path)
	errorhandler.CheckError(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	return file, scanner
}

func GetQuestionAnswerFromLine(line string) (string, string) {
	r := csv.NewReader(strings.NewReader(line))
	r.Comma = ','
	records, err := r.Read()
	errorhandler.CheckError(err)
	return records[0], records[1]
}
