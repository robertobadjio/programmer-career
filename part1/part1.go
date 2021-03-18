package part1

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// Time complexity O(n)
// Spatial complexity O(2)
func StringComparison(a, b string) bool {
	if len(a) == len(b) {
		var flag = false
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				if flag {
					return false
				}
				flag = true
			}
		}
	} else {
		if math.Abs(float64(len(a)-len(b))) > 1 {
			return false
		}

		var flag = false
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				if flag {
					return false
				}

				if len(a) > len(b) {
					b = b[:i] + fmt.Sprintf("%c", a[i]) + b[i:]
				} else {
					a = a[:i] + fmt.Sprintf("%c", b[i]) + a[i:]
				}

				flag = true
			}
		}
	}

	return true
}

func StringCompression(str string) string {
	for i := 0; i < len(str); i++ {
		if i+1 >= len(str) {
			break
		}

		var j = 0
		for j = i + 1; j < len(str); j++ {
			if str[j] == str[i] {
				continue
			}
			break

		}

		str = str[:i+1] + strconv.Itoa(j-i) + str[j:]
		i++
	}

	return str
}

// Use 2x brute force
// Time complexity O(n^3logn)
// Spatial complexity O(1)
func Test(matrix [5][5]int) [5][5]int {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if matrix[i][j] == 0 {
				matrix[i][j] = -1
			}
		}
	}

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if matrix[i][j] == -1 {
				for k := 0; k < 5; k++ {
					matrix[i][k] = 0
					matrix[k][j] = 0
				}
			}
		}
	}

	return matrix
}

func PrintMatrix(matrix [5][5]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Print("\n")
	}
}

func GenerateRandomMatrix() [5][5]int {
	matrix := [5][5]int{}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			matrix[i][j] = rand.Intn(10)
		}
	}

	return matrix
}

// Use replace in place
// Time complexity O(n^2)
// Spatial complexity O(1)
func ReplaceSpacesInPlace(cArray []string) []string {
	for i := 0; i < len(cArray); i++ {
		if cArray[i] == " " {
			if i+1 < len(cArray) {
				for j := i + 1; j < len(cArray); j++ {
					if cArray[j] != " " {
						break
					}
					if j == len(cArray)-1 {
						return cArray[:i]
					}
				}
			}
			cArray[i] = "%20"
		}
	}

	return cArray
}

// Use brute force
// Time complexity O(n^2)
// Spatial complexity O(2)
func ReplaceSpacesBruteForce(c string) string {
	var output = ""
	for i := 0; i < len(c); i++ {
		if fmt.Sprintf("%c", c[i]) == " " {
			if i+1 < len(c) {
				for j := i + 1; j < len(c); j++ {
					if fmt.Sprintf("%c", c[j]) != " " {
						break
					}
					if j == len(c)-1 {
						return output
					}
				}
			}
			output += "%20"
		} else {
			output += fmt.Sprintf("%c", c[i])
		}
	}

	return output
}

// Use brute force
// Time complexity O(n^2)
// Spatial complexity O(1)
func PermutationStringBruteForce(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			if a[i] == b[j] {
				b = strings.Replace(b, fmt.Sprintf("%c", a[i]), "", 1)
				break
			}
		}
	}

	if b != "" {
		return false
	}

	return true
}

func PermutationSort(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	// Sort a
	// Sort b

	if a == b {
		return true
	}

	return false
}