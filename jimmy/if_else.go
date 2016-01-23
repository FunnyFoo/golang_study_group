package main

import "fmt"

func main() {
  if 7%2 == 0 {
    fmt.Println("even");
  } else {
    fmt.Println("odd")
  }
  if 8%4 == 0 {
    fmt.Println("ok")
  }
  if num := 9; num < 0 {
     fmt.Println(num)
  } else if num < 10 {
    fmt.Println(num)
  } else {
    fmt.Println(num)
  }
}
