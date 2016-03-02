package main

import (
	. "fmt"
	"os"
)

func main() {
	arg1 := os.Args
	Println(arg1[2])
}
