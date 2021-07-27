package panic

import "fmt"

func iWillPanic() {
	panic("something")
}

func catchPanic() {
	defer func() int {
		if r := recover(); r != nil {
			fmt.Printf("Panic because of: `%+v`\n", r)
		}
		return 0
	}()
	iWillPanic()
}
