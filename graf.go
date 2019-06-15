package main

import (
	"fmt"
	"os"
	"math/rand"
	"strconv"
	"time"
	"github.com/gookit/color"
	"strings"
)

var (
	dataArr []int
	dataTotal int
)

const bar = "\u2501"
const err = "Error: Invalid argument! Use the '-h' option for help"
const helpm = `Usage: graf [NUMBER]...
Graph data in a visually pleasing way

-h, --help: Display this information.
-s=n, --scale=n: Scale the bar to a certain size, the default is 30.`

func intArr(strNums []string) (out []int) {
	for _, elem := range strNums {
		num, _ := strconv.Atoi(elem)
		out = append(out, num)
	}
	return
}

func intArr2(floatNums []float64) (out []int) {
	for _, elem := range floatNums {
		out = append(out, int(elem))
	}
	return
}

func total(nums []int) (out int) {
	for _, elem := range nums {
		out = out + elem
	}
	return
}

func dispBar(nums []int, char string) {
	rn := rand.New(rand.NewSource(time.Now().UnixNano()))
	for _, elem := range nums {
		var (
			r uint8 = uint8(rn.Intn(255))
			g uint8 = uint8(rn.Intn(255))
			b uint8 = uint8(rn.Intn(255))
		)

		for i := 0; i < elem; i++ {
			color.RGB(r, g, b).Print(char)
		}
	}

	fmt.Println("")
}

func scaleBar(data []string, size float64) {
	nums := intArr(data)
	total := total(nums)
	var fractions []float64

	for _, elem := range nums {
		fractions = append(fractions, (float64(elem)/float64(total))*size)
	}

	dispBar(intArr2(fractions), bar)
}

func main() {
	if len(os.Args) == 1 {
		color.New(color.FgRed, color.Bold).Println(err)
	} else {
		if os.Args[1] == "-h" || os.Args[1] == "--help" {
			fmt.Println(helpm)
		} else if os.Args[1] == "-s" || os.Args[1] == "--scale" {
			scaleBar(os.Args[2:], 30)
		} else if strings.ContainsAny(os.Args[1], "-s=") || strings.ContainsAny(os.Args[1], "-scale=") {
			num, _ := strconv.Atoi(strings.Split(os.Args[1], "=")[1])
			scaleBar(os.Args[2:], float64(num))
		} else {
			dispBar(intArr(os.Args[2:]), bar)
		}
	}
}
