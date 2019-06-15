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
-s=n, --scale=n: Scale the bar to a certain size, the default is 30.
-p, --percent: Get a percentage graph.
-ps=n, --percentscale=n`

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

func scaleBar(data []int, size float64) {
	total := total(data)
	var fractions []float64

	for _, elem := range data {
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
			scaleBar(intArr(os.Args[2:]), 30)
		} else if os.Args[1] == "-p" || os.Args[1] == "--percent" {
			if len(os.Args[2:]) > 2 || len(os.Args[2:]) <= 1 {
				color.New(color.FgRed, color.Bold).Println("Error: Only two arguments needed")
			} else {
				nums := intArr(os.Args[2:])
				dispBar([]int{nums[0], nums[1]-nums[0]}, bar)
			}
		} else if os.Args[1] == "-ps" || os.Args[1] == "--percentscale"{
			if len(os.Args[2:]) > 2 || len(os.Args[2:]) <= 1 {
				color.New(color.FgRed, color.Bold).Println("Error: Only two arguments needed")
			} else {
				nums := intArr(os.Args[2:])
				scaleBar([]int{nums[0], nums[1]-nums[0]}, 30)
			}
		} else if strings.Contains(os.Args[1], "-ps=") || strings.Contains(os.Args[1], "-percentscale=") {
			if len(os.Args[2:]) > 2 || len(os.Args[2:]) <= 1 {
				color.New(color.FgRed, color.Bold).Println("Error: Only two arguments needed")
			} else {
				nums := intArr(os.Args[2:])
				size, _ := strconv.Atoi(strings.Split(os.Args[1], "=")[1])
				scaleBar([]int{nums[0], nums[1]-nums[0]}, float64(size))
			}
		} else if strings.Contains(os.Args[1], "-s=") || strings.Contains(os.Args[1], "-scale=") {
			size, _ := strconv.Atoi(strings.Split(os.Args[1], "=")[1])
			scaleBar(intArr(os.Args[1:]), float64(size))
		} else {
			dispBar(intArr(os.Args[1:]), bar)
		}
	}
}
