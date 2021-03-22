package part5

import (
	"fmt"
	"math/rand"
	"time"
)

func DeterminingNearestNumbers() {
	rand.Seed(time.Now().UnixNano())
	var num = rand.Intn(10000)
	var baseCount = calcNumOnes(num)
	fmt.Printf("Iterative %d %08b %d\n", num, num, baseCount)
	iterative4(num, baseCount)
}

// Временная сложность O(n^2)
// Пространственная сложность O(1)
func iterative4(num, baseCount int) {
	for i := num - 1; i > 0; i-- {
		var count = calcNumOnes(i)
		if baseCount == count {
			fmt.Printf("%d %08b %d\n", i, i, count)
			break
		}
	}

	for i := num + 1; ; i++ {
		var count = calcNumOnes(i)
		if baseCount == count {
			fmt.Printf("%d %08b %d\n", i, i, count)
			break
		}
	}
}

func calcNumOnes(num int) int {
	var count int
	for ; num > 0; count++ {
		num &= num - 1 // Убираем младший бит
	}

	return count
}
