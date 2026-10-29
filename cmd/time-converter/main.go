package main

import "fmt"

func main() {
	var minutes int
	fmt.Print("Введите количество минут: ")
	n, err := fmt.Scan(&minutes)
	if err != nil || n != 1 {
		fmt.Println("Ошибка чтения")
	} else {
		if minutes < 0 {
			fmt.Println("Число отрицательное")
		} else {
			fmt.Println("Полных часов: ", minutes/60)
			fmt.Printf("Оставшихся минут: %d", minutes-(minutes/60)*60)
		}
	}
}
