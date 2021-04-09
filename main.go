package main

import (
	"fmt"
	"math"
	"programmerCareer/other"
	"programmerCareer/part1"
	"programmerCareer/part2"
	"programmerCareer/part4"
	"programmerCareer/part5"
	"programmerCareer/part8"
	"regexp"
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

	// 4.2.
	fmt.Print("\n4.2.\n")
	part4.BinarySearchTree()

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

	part8.BubbleSort()

	// Other
	fmt.Print("\nOther. BubbleZero\n")
	other.BubbleZero()

	// Hackerrank
	fmt.Print("\nHackerrank. Encryption\n")
	fmt.Println(Encryption())
	fmt.Print("\nHackerrank. Encryption2\n")
	fmt.Println(EncryptionSlowFastPointer())
}

func InversionBits(num int64) {
	fmt.Println("%d\n%08b\n%08b\n%d\n", num, num, (^num)+(1<<32))
}

// Encryption https://www.hackerrank.com/challenges/encryption/problem
// Сложность алгоритмическая O(n^2)
// Сложность пространственная O(n)
func Encryption() string {
	//var s = "have a nice day"
	//var s = "chill out"
	var s = "iffactsdontfittotheorychangethefacts"

	r := regexp.MustCompile("\\s+")
	s = r.ReplaceAllString(s, "")

	rootSquare := math.Sqrt(float64(len(s)))
	rows := int(math.Floor(rootSquare))
	columns := int(math.Ceil(rootSquare))

	for rows * columns < len(s) {
		rows++
	}

	array := make([][]string, rows)
	j := 0
	k := 0
	array[j] = make([]string, columns)
	for i := 0; i < len(s); i++ {
		array[j][k] = fmt.Sprintf("%c", s[i])
		k++
		if (i+1)%columns == 0 {
			j++
			k = 0
			if j >= rows {
				break
			}
			array[j] = make([]string, columns)
		}
	}

	output := ""
	for i := 0; i < columns; i++ {
		for j = 0; j < rows; j++ {
			output += array[j][i]
		}
		output += " "
	}

	return output
}

// EncryptionSlowFastPointer https://www.hackerrank.com/challenges/encryption/problem
// Сложность алгоритмическая O(n)
// Сложность пространственная O(n)
func EncryptionSlowFastPointer() string {
	var s = "have a nice day"

	r := regexp.MustCompile("\\s+")
	s = r.ReplaceAllString(s, "")

	rootSquare := math.Sqrt(float64(len(s)))
	rows := int(math.Floor(rootSquare))
	columns := int(math.Ceil(rootSquare))

	for rows * columns < len(s) {
		rows++
	}

	output := ""
	for i := 0; i < columns; i++ {
		output += fmt.Sprintf("%c", s[i])
		j := i
		for j + columns < len(s) {
			output += fmt.Sprintf("%c", s[j + columns])
			j += columns
		}

		output += " "
	}

	return output
}