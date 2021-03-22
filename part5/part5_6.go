package part5

import (
	"fmt"
	"math/rand"
	"time"
)

func NumberOfDistinctBits() {
	rand.Seed(time.Now().UnixNano())
	var numOne = rand.Intn(10000)
	var numTwo = rand.Intn(10000)
	fmt.Printf("Number one %d %08b\n", numOne, numOne)
	fmt.Printf("Number two %d %08b\n", numTwo, numTwo)
	fmt.Printf("Num of bits %d\n", CalcNumOnes(xor6(numOne, numTwo)))
}

func xor6(numOne, numTwo int) int {
	return numOne ^ numTwo
}
