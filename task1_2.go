package main

import "fmt"

const firstArray = 4
const secondArray = 5
const finishArray = 9

func bubbleSort(array *[finishArray]int) {
	for i := 0; i < len(*array); i++ {
		for j := i; j < len(*array); j++ {
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
}

func main() {
	Array1 := [firstArray]int{2, 5, 8, 12}
	Array2 := [secondArray]int{1, 3, 6, 7, 9}
	var finishArray [finishArray]int
	n := 0
	n1 := 0

	for i := 0; i < len(finishArray); i++ {
		if i < firstArray {
			finishArray[i] = Array1[n]
			n++
		} else if n1 < secondArray {
			finishArray[i] = Array2[n1]
			n1++
		}
	}

	bubbleSort(&finishArray)

	fmt.Printf("Первый массив - %v\n", Array1)
	fmt.Printf("Второй массив - %v\n", Array2)
	fmt.Printf("После соединения - %v\n", finishArray)

}
