package part1

import "fmt"

// Временная сложность O(n^2)
// Пространственная сложность O(1)
func BruteForce(x string) bool {
	for i := 0; i < len(x); i++ {
		for j := 0; j < len(x); j++ {
			if i != j && x[i] == x[j] {
				return false
			}
		}
	}
	return true
}

// Временная сложность O(n)
// Пространственная сложность O(n)
func HashMap(x string) bool {
	count := make(map[string]bool)
	for _, symbol := range x {
		if _, ok := count[fmt.Sprintf("%c", symbol)]; ok {
			return false
		}
		count[fmt.Sprintf("%c", symbol)] = true
	}
	return true
}

// Временная сложность O(n)
// Пространственная сложность O(1)
func BitVector(str string) bool {
	type IntSet struct {
		words []uint
	}
	var s IntSet
	const sizeWord = 32 << (^uint(0) >> 63)

	for i := 0; i < len(str); i++ {
		var charCode = str[i]

		byteIndex, indexInByte := uint(charCode/sizeWord), uint(charCode%sizeWord)
		for byteIndex >= uint(len(s.words)) {
			s.words = append(s.words, 0)
		}

		var bit = s.words[byteIndex] & (1 << indexInByte)
		if bit > 0 {
			return false
		} else {
			s.words[byteIndex] |= 1 << indexInByte
		}
	}

	return true
}
