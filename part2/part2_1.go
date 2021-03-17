package part2

import (
	"container/list"
	"fmt"
	"math/rand"
	"time"
)

// O(n^2)
// O(1)
func RemoveDuplicateElement() {
	l := list.New()
	rand.Seed(time.Now().UnixNano())
	for j := 0; j < 5; j++ {
		l.PushFront(rand.Intn(10))
	}

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%d ", e.Value)
	}
	fmt.Println("")

	for element := l.Front(); element != nil; element = element.Next() {
		var next = element.Next()
		for elementI := element.Next(); elementI != nil; elementI = next {
			next = elementI.Next()
			if elementI.Value.(int) == element.Value.(int) {
				l.Remove(elementI)
			}
		}
	}

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%d ", e.Value)
	}
	fmt.Println("")
}

// O(n)
// O(2)
func RemoveDuplicateElementByHashMap() {
	l := generateList()
	printList(l)

	var hashMap = make(map[int]bool)
	var next *list.Element
	for element := l.Front(); element != nil; element = next {
		next = element.Next()
		if hashMap[element.Value.(int)] {
			l.Remove(element)
		} else {
			hashMap[element.Value.(int)] = true
		}
	}

	printList(l)
}

// Not working
// https://overcoder.net/q/3838572/%D1%83%D0%B4%D0%B0%D0%BB%D0%B5%D0%BD%D0%B8%D0%B5-%D0%B4%D1%83%D0%B1%D0%BB%D0%B8%D0%BA%D0%B0%D1%82%D0%BE%D0%B2-%D0%B8%D0%B7-%D1%81%D0%B2%D1%8F%D0%B7%D0%B0%D0%BD%D0%BD%D0%BE%D0%B3%D0%BE-%D1%81%D0%BF%D0%B8%D1%81%D0%BA%D0%B0-%D0%B2-java-%D0%B1%D0%B5%D0%B7-%D0%B8%D1%81%D0%BF%D0%BE%D0%BB%D1%8C%D0%B7%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F
func RemoveDuplicateElementsByDoublePointers() {
	l := generateList()
	printList(l)

	var prev = l.Front()
	var current = l.Front().Next()
	for prev.Next() != nil {
		if current != nil {
			fmt.Printf("%d %d\n", prev.Value, current.Value)
			if prev.Value == current.Value {
				l.Remove(prev)
			}
			prev = current
			current = current.Next()
		} else {
			prev = prev.Next()
			current = prev.Next()
		}
	}

	printList(l)
}