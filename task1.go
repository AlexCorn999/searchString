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

func merge(array1 [firstArray]int, array2 [secondArray]int) (thirdArray [finishArray]int) {
	i := 0

	for first, _ := range array1 {
		thirdArray[i] = array1[first]
		i++
	}

	for second, _ := range array2 {
		thirdArray[i] = array2[second]
		i++
	}
	return
}

func main() {
	Array1 := [firstArray]int{2, 5, 8, 12}
	Array2 := [secondArray]int{1, 3, 6, 7, 9}

	thirdArray := merge(Array1, Array2)
	bubbleSort(&thirdArray)

	fmt.Printf("После соединения - %v\n", thirdArray)

}
