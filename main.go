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
	if strings.Contains(arg, "=") {
		arg = strings.Split(arg, "=")[0]
	}
	if strings.HasPrefix(arg, "--") {
		fmt.Printf("%s: %s\n", arg, data.GetDesc(arg))
		data.AddNestedArgs(arg, indent)
	} else if strings.HasPrefix(arg, "-") {
		arg = strings.TrimPrefix(arg, "-")
		for _, r := range arg {
			s := "-" + string(r)
			fmt.Printf("%s: %s\n", s, data.GetDesc(s))
			data.AddNestedArgs(s, indent)
		}
	}
}
