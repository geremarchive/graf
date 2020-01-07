package main

import (
	"os"
	"fmt"
	fu "graf/funcs"
	flag "github.com/spf13/pflag"
)

const help string = `Usage: graf [OPTION] [NUMBER]...
Display colorful representations of data

--help, -h: Display this information
--size=[NUMBER], -s: Change the bar's size
--percent, -p: Generate a percent bar`

func main() {
	if len(os.Args) == 2 {
		if os.Args[1] == "--help" || os.Args[1] == "-h" {
			fmt.Println(help)
			os.Exit(0)
		}
	}

	var size int

	flag.IntVarP(&size, "size", "s", 1, "Change the bar's size")
	percent := flag.BoolP("percent", "p", false, "Generate a percent bar")

	flag.Parse()

	nargs, err := fu.ConvertArgs(flag.Args())

	if err != nil {
		panic(err)
	}

	if *percent {
		if len(nargs) == 2 {
			if size == 1 {
				fu.PrintBar(fu.PercentBar(nargs[0], nargs[1]))
			} else {
				fu.PrintBar(fu.ScaleBar(fu.PercentBar(nargs[0], nargs[1]), size))
			}
		} else {
			fmt.Println("Error: to many, or to little arguments")
		}
	} else {
		if size == 1 {
			fu.PrintBar(nargs)
		} else {
			fu.PrintBar(fu.ScaleBar(nargs, size))
		}

	}
}
