package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"github.com/geremachek/escape"
	fu "./funcs"
)

var (
	dataArr []int
	dataTotal int
)

const err = "Error: Invalid argument! Use the '-h' option for help"
const helpm = `Usage: graf [OPTIONS] [NUMBER]...
Graph data in a visually pleasing way

-h, --help: Display this information.
-s=n, --scale=n: Scale the bar to a certain size, the default is 30.
-p, --percent: Get a percentage graph.
-ps=n, --percentscale=n: Scale the percent bar.`

func main() {
	if len(os.Args) == 1 {
		fmt.Println(escape.Vint(31, 1) + err + escape.Vint(0))
	} else {
		if os.Args[1] == "-h" || os.Args[1] == "--help" {
			fmt.Println(helpm)
		} else if os.Args[1] == "-s" || os.Args[1] == "--scale" {
			fu.ScaleBar(fu.IntArr(os.Args[2:]), 30)
		} else if os.Args[1] == "-p" || os.Args[1] == "--percent" {
			if len(os.Args[2:]) > 2 || len(os.Args[2:]) < 2 {
				fmt.Println(escape.Vint(31, 1) + "Error: Only two arguments needed" + escape.Vint(0))
			} else {
				nums := fu.IntArr(os.Args[2:])
				fu.DispBar([]int{nums[0], nums[1]-nums[0]})
			}
		} else if os.Args[1] == "-ps" || os.Args[1] == "--percentscale"{
			if len(os.Args[2:]) > 2 || len(os.Args[2:]) < 2 {
				fmt.Println(escape.Vint(31, 1) + "Error: Only two arguments needed" + escape.Vint(0))
			} else {
				nums := fu.IntArr(os.Args[2:])
				fu.ScaleBar([]int{nums[0], nums[1]-nums[0]}, 30)
			}
		} else if strings.Contains(os.Args[1], "-ps=") || strings.Contains(os.Args[1], "-percentscale=") {
			if len(os.Args[2:]) > 2 || len(os.Args[2:]) < 2 {
				fmt.Println(escape.Vint(31, 1) + "Error: Only two arguments needed" + escape.Vint(0))
			} else {
				nums := fu.IntArr(os.Args[2:])
				size, _ := strconv.Atoi(strings.Split(os.Args[1], "=")[1])
				fu.ScaleBar([]int{nums[0], nums[1]-nums[0]}, float64(size))
			}
		} else if strings.Contains(os.Args[1], "-s=") || strings.Contains(os.Args[1], "-scale=") {
			size, _ := strconv.Atoi(strings.Split(os.Args[1], "=")[1])
			fu.ScaleBar(fu.IntArr(os.Args[1:]), float64(size))
		} else {
			fu.DispBar(fu.IntArr(os.Args[1:]))
		}
	}
}
