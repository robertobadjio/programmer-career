package part2

import (
	"container/list"
	"fmt"
	"math/rand"
	"time"
)

// O(n^2)
// O(2)
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
// O(3)
func RemoveDuplicateElementByHashMap() {
	l := list.New()
	rand.Seed(time.Now().UnixNano())
	for j := 0; j < 5; j++ {
		l.PushFront(rand.Intn(10))
	}

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%d ", e.Value)
	}
	fmt.Println("")

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

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%d ", e.Value)
	}
	fmt.Println("")
}