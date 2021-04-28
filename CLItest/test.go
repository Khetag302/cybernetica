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
	var swtime int
	var answer string

	timer2:=timer(swtime)
	done := make(chan bool, 1)
	to := time.After(time.Duration(timer2) * time.Second)

	file, err := os.Open("problem.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Print("Если вы готовы к тесту напишите \"да\"\n")
	fmt.Scan(&answer)

	if answer!="да"{
		os.Exit(0)
	}

	reader := csv.NewReader(file)
	reader.Comma = ','
	go func() {
		defer os.Exit(0)

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

			if correct+incorrect==14{
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

func timer(swtime int)int{
	var check int
fmt.Print("Сколько времени вы хотите:\n1)15 секунд\n2)30 секунд\n3)45 секунд\n5)60 секунд\n")
fmt.Scan(&check)

switch check{
case 1:
	swtime=15
case 2:
	swtime=30
case 3:
	swtime=45
case 4:
	swtime=60
default:
	swtime=30
}
return swtime
}