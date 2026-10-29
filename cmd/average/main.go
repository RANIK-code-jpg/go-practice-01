package main

import "fmt"

func main() {
	var minut1 int = 20
	var minut2 int = 25
	var minut3 int = 35
	var summ int = minut1 + minut2 + minut3
	var average float64 = float64(summ) / 3
	fmt.Println("сумма: ", summ)
	fmt.Printf("Среднее значение: %.2f", average)
}
