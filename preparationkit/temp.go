package main

import (
	"fmt"
	"strings"
)


func main() {
	x := func(s string) string {
			fmt.Println(strings.ToLower(s))
			return s
	}("RacheL")
	fmt.Printf("%T \n", x)
}