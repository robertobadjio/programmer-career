package part5

import (
	"fmt"
	"math/rand"
	"time"
)

func DeterminingNearestNumbers() {
	rand.Seed(time.Now().UnixNano())
	var num = rand.Intn(10000)
	var baseCount = CalcNumOnes(num)
	fmt.Printf("Iterative %d %08b %d\n", num, num, baseCount)
	iterative4(num, baseCount)
}

// Временная сложность O(n^2)
// Пространственная сложность O(1)
func iterative4(num, baseCount int) {
	for i := num - 1; i > 0; i-- {
		var count = CalcNumOnes(i)
		if baseCount == count {
			fmt.Printf("%d %08b %d\n", i, i, count)
			break
		}
	}

	for i := num + 1; ; i++ {
		var count = CalcNumOnes(i)
		if baseCount == count {
			fmt.Printf("%d %08b %d\n", i, i, count)
			break
		}
	}
}
