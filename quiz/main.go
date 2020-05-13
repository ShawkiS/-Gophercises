package main

import (
	"strings"
	"time"
	"encoding/csv"
	"fmt"
	"os"
	"flag"
)

func main()  {

	csvFileName := flag.String("csv", "problems.csv", "a csv in the format of question,answer")
	timeLimit :=   flag.Int("limit", 2, "The time limit for the quiz in seconds")

	flag.Parse()
	_ = csvFileName

	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFileName))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided csv file.")
	}
	
	problems := parseLines(lines)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	corret := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		answerCh := make(chan string)

		go func ()  {
			var answer string 
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <- timer.C:
			fmt.Println()
			return

		case answer := <-answerCh:
			if answer == p.a {
				corret++
			}
		}
	}
	fmt.Printf("you scored %d out of %d.\n ", corret, len(problems))
}

func parseLines(lines [][]string) []problem {
	ret :=  make([]problem, len(lines))

	for i, line := range lines {
	ret[i]=problem{
		q: line[0],
		a: strings.TrimSpace(line[1]),
	}
	}
	return ret;
}

type problem struct {
	q string
	a string
}

func exit(msg string)  {
	fmt.Println(msg)
	os.Exit(1)
}