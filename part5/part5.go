package part5

func CalcNumOnes(num int) int {
	var count int
	for ; num > 0; count++ {
		num &= num - 1 // Убираем младший бит
	}

	return count
}
