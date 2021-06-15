package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/SphinX-SiG/homework/module-2/task1/fibonacci"
)

func main() {
	var pos interface{}
	var err error
	var input string

	fmt.Print("Hello, please enter required pos from fibonacci sequence: ")
	fmt.Scanf("%s", &input)

	if strings.Contains(input, ".") {
		pos, err = strconv.ParseFloat(input, 64)
	}
	if strings.Contains(input, "-") {
		var i64 int64
		i64, err = strconv.ParseInt(input, 10, 32)
		pos = int(i64)
	}
	if !strings.ContainsAny(input, "-.") {
		var ui64 uint64
		ui64, err = strconv.ParseUint(input, 10, 64)
		pos = uint(ui64)
	}

	if err != nil {
		fmt.Printf("\nOperation not supported for %v.\nError:\n%v", input, err)
		return
	}

	fmt.Printf("On required pos: %v\n", fibonacci.Fibonacci(pos))
}
