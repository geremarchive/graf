package funcs

import (
	"fmt"
	"strconv"
	"github.com/geremachek/escape"
)

func IntArr(strNums []string) (out []int) {
	for _, elem := range strNums {
		num, _ := strconv.Atoi(elem)
		out = append(out, num)
	}
	return
}

func IntArr2(floatNums []float64) (out []int) {
	for _, elem := range floatNums {
		out = append(out, int(elem))
	}
	return
}

func Total(nums []int) (out int) {
	for _, elem := range nums {
		out = out + elem
	}
	return
}

func DispBar(nums []int) {
	for c, elem := range nums {
		for i := 0; i < elem; i++ {
			if c % 2 == 0 {
				fmt.Print("━" + escape.Vint(0))
			} else {
				fmt.Print("-" + escape.Vint(0))
			}
		}
	}

	fmt.Println("")
}

func ScaleBar(data []int, size float64) {
	total := Total(data)
	var fractions []float64

	for _, elem := range data {
		fractions = append(fractions, (float64(elem)/float64(total))*size)
	}

	DispBar(IntArr2(fractions))
}
