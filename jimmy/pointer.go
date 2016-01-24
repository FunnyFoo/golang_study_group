// Go supports <em><a href="http://en.wikipedia.org/wiki/Pointer_(computer_programming)">pointers</a></em>,
// allowing you to pass references to values and records
// within your program.

package main

import "fmt"

func zeroval(ival int) {
    ival = 0
}

func zeroptr(iptr, aptr *int) {
    fmt.Println("aptr: ", &*aptr)
    *iptr = *aptr
}

func main() {
    i := 1
    a := *&i
    i = 2
    fmt.Println("i:", i)
    fmt.Println("i pointer:", &i)

    fmt.Println("a:", a)
    fmt.Println("a pointer:", &a)

    zeroval(i)
    // fmt.Println("zeroval:", i)

    // zeroptr(&i, &a)

    fmt.Println("i:", i)

    fmt.Println("pointer:", &i)
}
