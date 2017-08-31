package main

import (
	"fmt"
)

func main() {
	fmt.Println("Traveling sales person")
	xfunc()
}

func xfunc() {
	x := 5-10
	fmt.Println(x)
	if x < 0 {
		x = -x
	}
	fmt.Println(x)


}
