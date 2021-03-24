package main

import (
	"fmt"
	"programmerCareer/part1"
	"programmerCareer/part2"
	"programmerCareer/part4"
	"programmerCareer/part5"
)

func main() {
	fmt.Println("1.1.")
	var str = "bca"
	fmt.Printf("%s %t\n", str, part1.BruteForce(str))
	fmt.Printf("%s %t\n", str, part1.HashMap(str))
	fmt.Printf("%s %t\n", str, part1.BitVector(str))

	fmt.Print("\n")
	fmt.Print("1.2.\n")
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

	// 4.1.
	fmt.Print("\n4.1.\n")
	part4.ExistsRoute()

	// 5.2.
	fmt.Print("\n5.2.\n")
	part5.RealNumber()

	// 5.3.
	fmt.Print("\n5.3. Determining max length sequence ones\n")
	part5.DeterminingLengthLongestSequenceOnes()

	// 5.4.
	fmt.Print("\n5.4. Determining nearest numbers\n")
	part5.DeterminingNearestNumbers()

	// 5.6.
	fmt.Print("\n5.6.\n")
	part5.NumberOfDistinctBits()

	// Test
	//fmt.Print("\nTest\n")
	//InversionBits(802743475)
}

func InversionBits(num int64) {
	fmt.Printf("%d\n%08b\n%08b\n%d\n", num, num, (^num)+(1<<32))
}
