package part2

// Если длина связаного списка неизвестна
// O(n)
// O(1)
func RemoveKElementOfEnd() {
	l := generateList(0)
	printList(l)

	var numToBeDeleted = 3
	var numElements = 0
	for element := l.Front(); element != nil; element = element.Next() {
		numElements++
	}
	var i = 0
	for element := l.Front(); element != nil; element = element.Next() {
		if i == numElements-numToBeDeleted {
			l.Remove(element)
		}
		i++
	}

	printList(l)
}

// Если длина связаного списка известна
// O(n)
// O(1)
func RemoveKElementOfEndN() {
	l := generateList(0)
	printList(l)

	var numToBeDeleted = 3
	var i = 0
	for element := l.Front(); element != nil; element = element.Next() {
		if i == 10-numToBeDeleted {
			l.Remove(element)
		}
		i++
	}

	printList(l)
}

func RemoveKElementOfEndDoublePointer() {
	l := generateList(0)
	printList(l)

	for element := l.Front(); element != nil; element = element.Next() {

	}

	printList(l)
}
