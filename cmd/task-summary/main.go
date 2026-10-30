package main

import "fmt"

// statusLabel превращает код статуса в понятную человеку подпись.
func statusLabel(status string) string {
	// TODO: обработайте new, in_progress, done через switch.
	// TODO: предусмотрите ветку default для незнакомого статуса.
	var status_rus string
	switch status {
	case "new":
		status_rus = "Новая"
	case "in_progress":
		status_rus = "В работе"
	case "done":
		status_rus = "Завершена"
	default:
		status_rus = "Неизвестный статус"
	}
	return status_rus
}

// countStatus считает количество элементов с заданным статусом.
func countStatus(statuses []string, target string) int {
	// TODO: пройдите по срезу через for range.
	// TODO: сравните каждый элемент с target и верните счётчик.
	var count int = 0
	for _, status := range statuses {
		if status == target {
			count++
		}
	}
	return count
}

func main() {
	statuses := []string{"new", "in_progress", "done", "done"}
	// TODO: добавьте ещё один статус new через append.
	statuses = append(statuses, "new")
	fmt.Printf("Задач всего: %d\n", len(statuses))
	for i, status := range statuses {
		fmt.Printf("%d. %s\n", i+1, statusLabel(status))
	}
	fmt.Printf("Завершено: %d\n", countStatus(statuses, "done"))

	// TODO: создайте map[string]int для подсчёта статусов.
	// TODO: выведите new, in_progress, done в фиксированном порядке.
	counts := make(map[string]int)
	for _, status := range statuses {
		counts[status]++
	}
	fmt.Printf("Статистика: new=%d, in_progress=%d, done=%d\n", counts["new"], counts["in_progress"], counts["done"])
	fmt.Println(statusLabel("blocked"))
	fmt.Println(countStatus([]string{}, "done"))
}
