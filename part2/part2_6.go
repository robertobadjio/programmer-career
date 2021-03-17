package part2

import "container/list"

// Время O(n^2)
// Память O(n)
func IsPalindrome() bool {
	l := generateList(3)
	printList(l)

	l2 := list.New()
	for element := l.Back(); element != nil; element = element.Prev() {
		l2.PushBack(element.Value)
	}

	var i = 0
	for element := l.Front(); element != nil; element = element.Next() {
		var j = 0
		for element2 := l2.Front(); element2 != nil; element2 = element2.Next() {
			if i == j && element.Value != element2.Value {
				return false
			}
			j++
		}
		i++
	}

	return true
}
