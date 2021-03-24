package part5

import (
	"fmt"
	"math"
	"strconv"
)

func RealNumber() {
	var number = 0
	var binary = ""
	for ; number >= 1; number = number / 2 {
		temp := strconv.Itoa(number % 2)
		binary = temp + binary
	}

	var numberFloat = 0.72
	binary += "."
	numberFloat *= 2
	var count = 0
	for ; math.Mod(numberFloat, 2) != 0; numberFloat = numberFloat * 2 {
		binary += strconv.Itoa(int(math.Mod(numberFloat, 2)))
		count++
		if count > 32 {
			break
		}
	}

	if count > 32 {
		fmt.Println("Ошибка! Для точного представления числа необходимо больше 32х разрядов")
	} else {
		fmt.Printf("%s %d\n", binary, count)
	}
}
