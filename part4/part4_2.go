package part4

import (
	"programmerCareer/part4/binary_tree"
)

func BinarySearchTree() {
 	var arr = initSliceWithUniqueElements()

 	node := createBalancedTree(arr, 0, len(arr) - 1)
	node.Traverse()
}

func createBalancedTree(arr [9]int, start, end int) *part4_binary_tree.Node {
	if end < start {
		return nil
	}
	middle := (start + end) / 2

	node := part4_binary_tree.New(arr[middle])

	node.Left = createBalancedTree(arr, start, middle - 1)
	node.Right = createBalancedTree(arr, middle + 1, end)

	return node
}

func initSliceWithUniqueElements() [9]int {
	arr := [9]int{1, 3, 4, 6, 7, 8, 10, 14, 13}
	//rand.Seed(time.Now().UnixNano())
	//for i := 0; i < 10; i++ {
		//arr[i] = rand.Intn(10
		//arr[i] = i
	//}

	return arr
}
