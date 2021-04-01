package other

import "fmt"

func BubbleZero() {
	var arr = []int{2, 0, 3, 10, 4, 0, 1}
	fmt.Println(arr)
	fmt.Println(bubbleZeroIterable(arr))
	fmt.Println(bubbleZeroSlowAndFastPointers(arr))
}

// Реализация итеративным методом "in-place"
// Временная сложность O(n^2)
// Пространственная сложность O(1)
func bubbleZeroIterable(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			continue
		}
		for j := i; j < len(arr) - 1; j++ {
			if arr[j + 1] != 0 {
				var temp = arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = temp
			}
		}
	}

	return arr
}

// Реализация двумя указателями
// Временная сложность O(n)
// Пространственная сложность O(1)
func bubbleZeroSlowAndFastPointers(arr []int) []int {
	var slow = 0
	var fast = 0

	for ; fast < len(arr); {
		if arr[fast] != 0 {
			arr[slow] = arr[fast]
			slow++
		}
		fast++
	}

	for ; slow < len(arr); {
		arr[slow] = 0
		slow++
	}

	return arr
}