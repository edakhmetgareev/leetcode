package main

import (
	"fmt"
	"math"
)

func main() {
	i := 1463847412

	res := reverse(i)
	fmt.Printf("result: %d", res)
}

// как второий вариант решения - можно перевести в строку, потом в массив байтов, потом сделать реверс (учитывая знак -)
// строки, проверить на максимальное значение.
func reverse(x int) int {
	negative := x < 0
	if negative {
		x *= -1
	}

	if x > math.MaxInt {
		return 0
	}

	res := 0
	for x > 0 {
		res = res*10 + x%10
		x /= 10
	}

	if math.MaxInt32 <= res {
		return 0
	}

	if negative {
		res *= -1
	}

	return res
}
