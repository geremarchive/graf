package main

import (
	"fmt"
	"os"
	"strconv"
	"math/rand"
	"time"
)

var (
	dataArr []int
	dataTotal int
)

const escape = "["
const bar = "━"
const err = escape + "31;1m" + "Error: Invalid argument! Use the '-h' option for help\n"
const help = "Usage: grapho [NUMBER]...\nGraph data in a visually pleasing way"

func intArr(strNums []string) (out []int) {
	for _, elem := range strNums {
		num, _ := strconv.Atoi(elem)
		out = append(out, num)
	}
	return
}

func main() {
	if len(os.Args) == 1 {
		fmt.Print(err)
	} else {
		if os.Args[1] == "--help" || os.Args[1] == "-h" {
			fmt.Println(help)
		} else {
			dataArr = intArr(os.Args[1:])

			rn := rand.New(rand.NewSource(time.Now().UnixNano()))

			for _, elem := range dataArr {

				var (
					r string = strconv.Itoa(rn.Intn(255))
					g string = strconv.Itoa(rn.Intn(255))
					b string = strconv.Itoa(rn.Intn(255))
				)

				for i := 0; i < elem; i++ {
					fmt.Print(escape + "38;2;" + r + ";" + g + ";" + b + "m" + bar)
				}
			}

			fmt.Println("")
		}
	}
}
