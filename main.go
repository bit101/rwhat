package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/bit101/rwhat/data"
)

func main() {
	args := os.Args
	for i := 1; i < len(args); i++ {
		splain(args[i])
	}
}

func splain(arg string) {
	indent := 3
	if strings.HasPrefix(arg, "--") {
		fmt.Printf("%s: %s\n", arg, data.GetDesc(arg))
		data.AddNestedArgs(arg, indent)
	} else if strings.HasPrefix(arg, "-") {
		a := strings.TrimPrefix(arg, "-")
		for _, r := range a {
			s := "-" + string(r)
			fmt.Printf("%s: %s\n", s, data.GetDesc(s))
			data.AddNestedArgs(s, indent)
		}
	}
}
