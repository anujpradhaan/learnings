package main

import (
	"flag"
	"fmt"
	"quiz-1/part1"
)

func main() {
	filename := flag.String("file", "problems.csv", "CSV File to parse")
	flag.Parse()
	fmt.Println("FileName = " + *filename)
	part1.Start(*filename)
}
