package main

import (
	"fmt"
)

type Item struct {
	weight int
	value  int
	number int
}

func knapsackBruteForce(capacity int, items []Item) (int, []Item) {
	return knapsackHelper(capacity, items, 0, 0, []Item{})
}

func knapsackHelper(capacity int, items []Item, index int, currentValue int, bag []Item) (int, []Item) {
	// Базовый случай: если достигнут конец списка предметов, возвращаем текущую стоимость
	if index == len(items) {
		return currentValue, bag
	}

	// Если вес текущего предмета превышает оставшуюся вместимость рюкзака, не берем его
	if items[index].weight > capacity {
		return knapsackHelper(capacity, items, index+1, currentValue, bag)
	}

	// Рекурсивно проверяем два случая: берем текущий предмет и не берем его
	new_bag := append(bag, items[index])
	include, bag1 := knapsackHelper(capacity-items[index].weight, items, index+1, currentValue+items[index].value, new_bag)
	exclude, bag2 := knapsackHelper(capacity, items, index+1, currentValue, bag)

	// Возвращаем максимальную стоимость из двух вариантов
	if include > exclude {
		return include, bag1
	}
	return exclude, bag2
}

func main() {
	capacity := 250
	items := []Item{
		{weight: 100, value: 10, number: 1},
		{weight: 100, value: 2, number: 2},
		{weight: 100, value: 3, number: 3},
		{weight: 110, value: 4, number: 4},
	}

	maxTotalValue, bag := knapsackBruteForce(capacity, items)
	fmt.Println("Максимальная сумма:", maxTotalValue)
	fmt.Println("Были взяты следующие предметы")
	for _, e := range bag {
		fmt.Printf("Вес: %d, Ценность: %d, Номер: %d\n", e.weight, e.value, e.number)
	}
}
