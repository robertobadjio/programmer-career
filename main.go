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

	fmt.Print("\nHackerrank. Forming a magic square\n")
	fmt.Println(FormingMagicSquare())
}

// FormingMagicSquare https://www.hackerrank.com/challenges/magic-square-forming/problem
// Сложность алгоритмическая O(n)
// Сложность пространственная O(1)
func FormingMagicSquare() int32 {
	//matrix := [3][3]int32{{4, 5, 8}, {2, 4, 1}, {1, 9, 7}} // 14
	//matrix := [3][3]int32{{2, 9, 8}, {4, 2, 7}, {5, 6, 7}} // 21
	//matrix := [3][3]int32{{7, 2, 9}, {6, 6, 7}, {5, 1, 2}} // 19
	//matrix := [3][3]int32{{4, 4, 7}, {3, 1, 5}, {1, 7, 9}} // 20
	//matrix := [3][3]int32{{2, 2, 7}, {8, 6, 4}, {1, 2, 9}} // 16
	matrix := [3][3]int32{{2, 5, 4}, {4, 6, 9}, {4, 5, 2}} // 16
	possibleMagicSquares := [][]int32{
		{4,9,2,3,5,7,8,1,6},
		{4,3,8,9,5,1,2,7,6},
		{2,9,4,7,5,3,6,1,8},
		{2,7,6,9,5,1,4,3,8},
		{8,1,6,3,5,7,4,9,2},
		{8,3,4,1,5,9,6,7,2},
		{6,7,2,1,5,9,8,3,4},
		{6,1,8,7,5,3,2,9,4},
	}

	cost := int32(81)
	temp := int32(0)
	for i := 0; i < 8; i++ {
		temp = int32(math.Abs(float64(matrix[0][0] - possibleMagicSquares[i][0])) + math.Abs(float64(matrix[0][1] - possibleMagicSquares[i][1])) +
			math.Abs(float64(matrix[0][2] - possibleMagicSquares[i][2])) + math.Abs(float64(matrix[1][0] - possibleMagicSquares[i][3])) +
			math.Abs(float64(matrix[1][1] - possibleMagicSquares[i][4])) + math.Abs(float64(matrix[1][2] - possibleMagicSquares[i][5])) +
			math.Abs(float64(matrix[2][0] - possibleMagicSquares[i][6])) + math.Abs(float64(matrix[2][1] - possibleMagicSquares[i][7])) +
			math.Abs(float64(matrix[2][2] - possibleMagicSquares[i][8])))
		if temp < cost {
			cost = temp
		}
	}

	return cost
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

	for rows*columns < len(s) {
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

	for rows*columns < len(s) {
		rows++
	}

	output := ""
	for i := 0; i < columns; i++ {
		output += fmt.Sprintf("%c", s[i])
		j := i
		for j+columns < len(s) {
			output += fmt.Sprintf("%c", s[j+columns])
			j += columns
		}

		output += " "
	}

	return output
}
