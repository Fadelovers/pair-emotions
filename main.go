package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var number [80]int      // нам нужно только 80 элементов, не 8000
	var result [10]int      // 10 средних
	var totalOfAverages int // будет суммой всех 10 средних

	// Проходим по 10 выборкам
	for j := 0; j < 10; j++ {
		var sum int // сбрасываем накопитель суммы для каждой выборки

		// Генерируем 80 чисел и считаем их сумму
		for i := 0; i < 80; i++ {
			number[i] = rand.Intn(100)
			if number[i] < 4 {
				number[i] = 4
			}
			sum += number[i]
			// Если нужно видеть сами сгенерированные числа:
			fmt.Printf("%d ", number[i])
		}

		// Среднее по первой выборке — целочисленное
		avg := sum / 80
		result[j] = avg
		totalOfAverages += avg

		fmt.Printf("\nВыборка #%d: среднее = %d\n\n", j+1, avg)
	}

	// Среднее из 10 средних
	finalAvg := totalOfAverages / 10
	fmt.Printf("Итоговое среднее по 10 выборкам: %d\n", finalAvg)
}
