package m

import (
	"bufio"
	"cybernetica/Golang/internal"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Введи два числа: ")

	reader := bufio.NewReader(os.Stdin)
	inp, _ := reader.ReadString('\n')

	// Парсим то, что ввел пользователь используя функцию, которая лежит в папке internal
	a, n := internal.ParseInput(inp)

	res := internal.Pow(a, n)
	fmt.Printf("%d в степени %d будет равно %d\n", a, n, res)
}
