package main

import "fmt"

const size = 6

func bubbleSort(array *[size]int) {
	for i := 0; i < len(*array); i++ {
		for j := i; j < len(*array); j++ {
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
}

func main() {
	array := [size]int{4, 2, 55, 32, 1, 4}
	fmt.Printf("До сортировки - %v\n", array)
	bubbleSort(&array)
	fmt.Printf("После сортировки - %v\n", array)
}
