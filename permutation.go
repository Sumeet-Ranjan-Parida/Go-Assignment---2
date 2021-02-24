package main

import "fmt"

func swap(data []int, x int, y int) {
	data[x], data[y] = data[y], data[x]
}
func Permutation(data []int, i int, length int) {
	if length == i {
		fmt.Println(data)
		return
	}

	for j := i; j < length; j++ {
		swap(data, i, j)
		Permutation(data, i+1, length)
		swap(data, i, j)
	}
}

func main() {
	var data [3]int
	for i := 0; i < 3; i++ {
		data[i] = i
	}
	Permutation(data[:], 0, 3)
}

//Output:

// [0 1 2]
// [0 2 1]
// [1 0 2]
// [1 2 0]
// [2 1 0]
// [2 0 1]
