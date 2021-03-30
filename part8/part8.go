package part8

import "fmt"

// O(n^2)
// O(1)
func BubbleSort() {
	var arr = []int{10, 3, 15, 16, 10, 1, 8}
	fmt.Println(arr)
	for i := len(arr) - 1; i >= 0; i-- {
		for j := i; j >= 0; j-- {
			if j-1 >= 0 && arr[i] < arr[j-1] {
				var temp = arr[j - 1]
				arr[j - 1] = arr[i]
				arr[i] = temp
			}
		}
	}
	fmt.Println(arr)
}
