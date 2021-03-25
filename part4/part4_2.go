package part4

import (
	"programmerCareer/part4/binary_tree"
)

func BinarySearchTree() {
 	var arr = initSliceWithUniqueElements()

	node := part4_binary_tree.New(arr[0])
	for i := 1; i < len(arr); i++ {
		node = insertSearch(node, arr[i])
	}

	node.Traverse()
}

// InsertSearch : adds a node to the search tree
func insertSearch(t *part4_binary_tree.Node, data int) *part4_binary_tree.Node {
	if t == nil {
		return part4_binary_tree.New(data)
	} else if data < t.Data {
		t.Left = insertSearch(t.Left, data)
	} else if data > t.Data {
		t.Right = insertSearch(t.Right, data)
	}

	return t
}

func initSliceWithUniqueElements() [6]int {
	arr := [6]int{2, 4, 6, 8, 10, 20}
	//rand.Seed(time.Now().UnixNano())
	//for i := 0; i < 10; i++ {
		//arr[i] = rand.Intn(10
		//arr[i] = i
	//}

	return arr
}
