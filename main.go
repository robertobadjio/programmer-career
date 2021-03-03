package main

import (
	"fmt"
)

func main() {
	// 1.1.
	var x = "abbc"
	fmt.Printf("%s %t\n", x, bruteForce(x))
	fmt.Printf("%s %t\n", x, hashMap(x))
	fmt.Printf("%s %t\n", x, bitVector(x))
}

// Use brute force
// Time complexity O(n^2)
// Spatial complexity O(1)
func bruteForce(x string) bool {
	for i := 0; i < len(x); i++ {
		for j := 0; j < len(x); j++ {
			if i != j && x[i] == x[j] {
				return false
			}
		}
	}
	return true
}

// Use map (hash table)
// Time complexity O(n)
// Space complexity O(n)
func hashMap(x string) bool {
	count := make(map[string]bool)
	for _, symbol := range x {
		if _, ok := count[fmt.Sprintf("%c", symbol)]; ok {
			return false
		}
		count[fmt.Sprintf("%c", symbol)] = true
	}
	return true
}

// Use bit vector
// See https://coderoad.ru/38556390/%D0%9F%D0%BE%D0%B8%D1%81%D0%BA-%D0%B4%D1%83%D0%B1%D0%BB%D0%B8%D0%BA%D0%B0%D1%82%D0%BE%D0%B2-%D1%8D%D0%BB%D0%B5%D0%BC%D0%B5%D0%BD%D1%82%D0%BE%D0%B2-%D1%81-%D0%BE%D0%B3%D1%80%D0%B0%D0%BD%D0%B8%D1%87%D0%B5%D0%BD%D0%BD%D0%BE%D0%B9-%D0%BF%D0%B0%D0%BC%D1%8F%D1%82%D1%8C%D1%8E
// See https://germangorelkin.blogspot.com/2017/05/blog-post_21.html
// Time complexity O(n)
// Space complexity O(1)
func bitVector(x string) bool {
	type IntSet struct {
		words []uint
	}
	var s IntSet

	for i := 0; i < len(x); i++ {
		var charCode = x[i]
		byteIndex, indexInByte := uint(charCode/64), uint(charCode%64)
		for byteIndex >= uint(len(s.words)) {
			s.words = append(s.words, 0)
		}

		var bit = s.words[byteIndex] & (1 << indexInByte) // TODO ?

		if bit > 0 {
			return false
		} else {
			s.words[byteIndex] |= 1 << indexInByte
		}
	}
	return true
}