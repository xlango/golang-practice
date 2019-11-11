package algorithm

import "fmt"

func BubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			} else {
				continue
			}
		}
	}

	fmt.Println(arr)
}

func SelectSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			} else {
				continue
			}
		}
	}

	fmt.Println(arr)
}

func InsertSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}

	fmt.Println(arr)
}

func QuickSort(arr []int, left, right int) {
	if left > right {
		return
	}

	base := left
	i := left
	j := right

	for i != j {
		for arr[j] >= arr[base] && i < j {
			j--
		}

		for arr[i] <= arr[base] && i < j {
			i++
		}

		arr[i], arr[j] = arr[j], arr[i]

	}

	arr[base], arr[i] = arr[i], arr[base]

	QuickSort(arr, left, i-1)
	QuickSort(arr, i+1, right)

}
