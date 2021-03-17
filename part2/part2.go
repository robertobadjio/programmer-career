package part2

import (
	"container/list"
	"fmt"
	"math/rand"
	"time"
)

func generateList(n int) *list.List {
	if n == 0 {
		n = 10
	}
	l := list.New()
	rand.Seed(time.Now().UnixNano())
	for j := 0; j < n; j++ {
		l.PushFront(rand.Intn(10))
	}

	return l
}

func printList(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%d ", e.Value)
	}
}