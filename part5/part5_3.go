package part5

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func DeterminingLengthLongestSequenceOnes() {
	rand.Seed(time.Now().UnixNano())
	var num = rand.Intn(10000)
	fmt.Printf("Iterative %d %08b %d\n", num, num, iterative(num))
}

// Метод "грубой силы", проверяем все возможные варианты
// Временная сложность O(n)
// Пространственная сложность O(1)
func iterative(num int) int {
	var binary = fmt.Sprintf("%08b", num)
	length, _ := strconv.Atoi(fmt.Sprintf("%d", len(binary)))
	var max = 0

	for i := 0; i < length; i++ {
		// (1 << i) - маска, сдвигаем бит при каждой итерации влево на i, при это освободившиеся позиции устанавливаем 0
		// Орерация ИЛИ | используется для включения битов по маске (1 << i)
		var value = num | (1 << i)
		//fmt.Printf("%d %08b %08b\n", value, value, 1 << i)

		// Определяем длину самой длинной последовательности из единиц
		// Подсчитываем сколько раз мы можем выполнить "побитовое И" смещенное на 1 бит вправо, пока оно не станет нулевым
		var count int
		for ; value > 0; count++ {
			//fmt.Printf("%08b %08b ", value, value >> 1)
			value &= value >> 1
			//fmt.Printf("%08b\n", value)
		}

		if count > max {
			max = count
		}
		//fmt.Printf("\n%d\n", count)
	}

	return max
}
