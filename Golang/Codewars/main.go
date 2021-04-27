package Codewars

import (
	"fmt"
	"strings"
)

func codeWars() {

	var z int

	fmt.Print("Задание:\n1.Первое;\n2.Второе;\n3.Третье;\n")
	fmt.Scan(&z)
	if z == 1 {
		TwiceAsOld()
	}
	if z == 2 {
		CountTheDivisorsOfaNumber()
	}
	if z == 3 {
		ReturningStrings()
	}

}

func TwiceAsOld() {
	var dadYearsOld, sonYearsOld int = 42, 21
	var result int

	if 2*sonYearsOld < dadYearsOld {
		result = (2*sonYearsOld - dadYearsOld) * -1
	} else {
		result = 2*sonYearsOld - dadYearsOld
	}
	fmt.Print(result)
}

func CountTheDivisorsOfaNumber() {
	var x int = 0
	var n int = 50000
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			x = x + 1
		}
	}
	fmt.Print(x)
}

func ReturningStrings() {
	name := ", Ryan "

	stringg := strings.Replace("Hello, how are you doing today?", ", ", name, -1)

	fmt.Printf(stringg)
}
