package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {

	answers := map[string]int{
		"correct":   0,
		"incorrect": 0,
	}


	filepath := flag.String("file", "problem.csv", "file name")
	timerFlag := flag.Int("t", 5, "timer")
	flag.Parse()


	timer := time.NewTimer(time.Duration(*timerFlag) * time.Second)
	file, err := os.Open(*filepath)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	inp := make(chan string)
LOOP:
	for {
		record, e := reader.Read()
		if e != nil {
			fmt.Println(e)
			break
		}
		fmt.Println(record[0])

		go func() {
			var numb string
			fmt.Scanln(&numb)
			inp <- numb
		}()

		select {
		case <-timer.C:
			break LOOP
		case b := <-inp:
			if b == record[1] {
				answers["correct"]++
			} else {
				answers["incorrect"]++
			}
		}
	}

	PrintResult(answers)
}

func PrintResult(correct map[string]int) {
	fmt.Printf("Правельных: %d; Неправельных:%d", correct["correct"], correct["incorrect"])
}
