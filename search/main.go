package main

import (
	"fmt"
)

func binarySearch(data []int, v int) (pos int) {
	var first, last, mid int
	var count int

	last = len(data) - 1

	if v < data[first] || v > data[last] {
		return -1
	}

	defer func() { fmt.Printf("%d times\n", count) }()

	//for loop
	for first <= last {
		count++
		mid = (first + last) / 2
		if data[mid] == v {
			return mid
		} else if data[mid] < v {
			first = mid + 1
		} else {
			last = mid - 1
		}
	}

	return -1
}

func linearSearch(data []int, v int) (pos int) {
	var count int

	defer func() { fmt.Printf("%d times\n", count) }()

	for i, d := range data {
		count++
		if d == v {
			return i
		}
	}

	return -1
}

func main() {
	data := make([]int, 20)
	for i := 0; i < 20; i++ {
		data[i] = i + 2
	}

	fmt.Printf("data: %v\n", data)
	fmt.Println("linear search:")
	linearSearch(data, 2)
	fmt.Println("binary search:")
	binarySearch(data, 2)

}

