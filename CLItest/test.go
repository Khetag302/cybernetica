package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

func main() {
	var numb string
	var correct int = 0
	done := make(chan bool, 1)
	to := time.After(5 * time.Second)

	file, err := os.Open("problems.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ','
	go func() {
		defer fmt.Printf("Время вышло\n")

		for {
			record, e := reader.Read()
			if e != nil {
				fmt.Println(e)
				break
			}

			fmt.Println(record[0])

			select {
			case <-to:
				fmt.Print("https://www.youtube.com/watch?v=P37mn84nabg&ab_channel=baton1337\n")
				done <- true
				return
			default:
				fmt.Scanln(&numb)
			}

			if numb == record[1] {
				correct++
			}
		}
	}()

	<-done
	fmt.Print("Медленно")

	result(correct)
}

func result(correct int) {
	if correct <= 4 {
		fmt.Print("https://www.youtube.com/watch?v=P37mn84nabg&ab_channel=baton1337")
	} else if correct > 4 && correct <= 7 {
		fmt.Print("Я думал все намного хуже")
	} else if correct > 7 && correct <= 10 {
		fmt.Print("Мммм... неплохо")
	} else if correct > 10 {
		fmt.Print("А ты хорош")
	}
}
