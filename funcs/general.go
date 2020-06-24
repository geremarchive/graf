package funcs

import (
	"fmt"
	"os"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"bufio"
	"github.com/geremachek/escape"
)

func PercentBar(top int, bottom int) []int {
	return []int{top, bottom-top}
}

func ScaleBar(numbers []int, size int) (scaled []int) {
	total := sliceTotal(numbers)

	for _, num := range numbers {
		scaled = append(scaled, int((float32(num) / float32(total)) * float32(size)))
	}

	return
}

func sliceTotal(slice []int) (total int) {
	for _, num := range slice {
		total += num
	}

	return
}

func PrintBar(numbers []int) {
	var (
		r int
		g int
		b int
	)

	rand.Seed(time.Now().UnixNano())

	for _, num := range numbers {
		r = rand.Intn(256)
		g = rand.Intn(256)
		b = rand.Intn(256)

		for i := 0; i < num; i++ {
			fmt.Print(escape.Vint(48, 2, r, g, b) + " " + escape.Vint(0))
		}
	}

	fmt.Println()
}

func ConvertArgs(args []string) (out []int, err error) {
	for _, elem := range args {
		var num float64
		num, err = strconv.ParseFloat(elem, 64)

		if err != nil {
			return
		}

		out = append(out, int(num))
	}

	return
}

func GetStdin() (elems []string, err error) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		for _, elem := range strings.Split(scanner.Text(), " ") {
			elems = append(elems, elem)
		}
	}

	err = scanner.Err()

	return
}
