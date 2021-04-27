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
	var incorrect int = 0
	done := make(chan bool, 1)
	to := time.After(30 * time.Second)

	file, err := os.Open("problems.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ','
	go func() {
		defer fmt.Printf("https://www.pngitem.com/pimgs/m/258-2583502_transparent-shiny-png-slowpoke-shiny-png-download.png\n")

		for {
			record, e := reader.Read()
			if e != nil {
				fmt.Println(e)
				break
			}			

			select {
			case <-to:				
				done <- true
				return
			default:
				fmt.Println(record[0])
				fmt.Scanln(&numb)			
			}

			if numb == record[1] {
				correct++
			}else{
				incorrect++
			}

			if correct+incorrect==13{
				result(correct)
			}
		}
		
	}()

	<-done
	fmt.Print("Ты:")


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
	os.Exit(0)
}
