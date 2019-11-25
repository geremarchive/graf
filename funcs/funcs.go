package funcs

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
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

func DispBar(nums []int, char string) {
	rn := rand.New(rand.NewSource(time.Now().UnixNano()))
	for _, elem := range nums {
		var (
			r int = rn.Intn(255)
			g int = rn.Intn(255)
			b int = rn.Intn(255)
		)

		for i := 0; i < elem; i++ {
			fmt.Print(escape.Vint(38, 2, r, g, b) + char + escape.Vint(0))
		}
	}

	fmt.Println("")
}

func ScaleBar(data []int, size float64, char string) {
	total := Total(data)
	var fractions []float64

	for _, elem := range data {
		fractions = append(fractions, (float64(elem)/float64(total))*size)
	}

	DispBar(IntArr2(fractions), char)
}
