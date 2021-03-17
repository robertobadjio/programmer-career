package main

import (
	"fmt"
	"programmerCareer/part1"
	"programmerCareer/part2"
)

func main() {
	// 1.1.
	var x = "abbc"
	fmt.Printf("%s %t\n", x, part1.BruteForce(x))
	fmt.Printf("%s %t\n", x, part1.HashMap(x))
	fmt.Printf("%s %t\n", x, part1.BitVector(x))

	// 1.2.
	var a = "abbcdeac"
	var b = "edcbbaac"
	fmt.Printf("%t\n", part1.PermutationStringBruteForce(a, b))
	fmt.Printf("%t\n", part1.PermutationSort(a, b))

	// 1.3.
	var c = "Mr John Smith     "
	fmt.Printf("%s\n", part1.ReplaceSpacesBruteForce(c))
	var cArray = []string{"M", "r", " ", "J", "o", "h", "n", " ", "S", "m", "i", "t", "h", " ", " ", " ", " "}
	fmt.Println(part1.ReplaceSpacesInPlace(cArray))

	// 1.5.
	var e = "pale"
	var f = "ple"
	fmt.Println(part1.StringComparison(e, f))

	// 1.6.
	var d = "aaabbfbdddaaabb"
	fmt.Println(part1.StringCompression(d))

	// 1.8.
	var matrix = part1.GenerateRandomMatrix()
	part1.PrintMatrix(matrix)
	fmt.Println("")
	part1.PrintMatrix(part1.Test(matrix))

	// 2.1.
	fmt.Println("2.1.")
	part2.RemoveDuplicateElement()
	part2.RemoveDuplicateElementByHashMap()
	//part2.RemoveDuplicateElementsByDoublePointers() // Not working

	// 2.2.
	fmt.Println("2.2.")
	part2.RemoveKElementOfEnd()

	// 2.6.
	fmt.Println("")
	fmt.Println("2.2. Is palindrome?")
	fmt.Printf("%t\n", part2.IsPalindrome())
}
