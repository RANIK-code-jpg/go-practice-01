package main

import "fmt"

func main() {
	const course string = "Бэкенд-разработка на языке Go"
	var name string = "Алексей"
	group := "ЭФБО-01-26"
	completed := 1
	var interes_for_backend bool = true
	fmt.Println("Участник: ", name)
	fmt.Println("Группа", group)
	fmt.Printf("Курс: %s \n", course)
	fmt.Printf("Завершено работ: %d \n", completed)
	fmt.Println("Интерес к бэкенду: ", interes_for_backend)
}
