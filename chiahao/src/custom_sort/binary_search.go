package main

import (
  . "fmt"
)

func BSearch(target int, arr []int, bot int, top int) (int, bool) {
  if bot > top {
    return -1, false
  }

  var middle = (top + bot)/2

  if arr[middle] > target {
    return BSearch(target, arr, bot, middle-1)
  } else if arr[middle] < target {
    return BSearch(target, arr, middle+1, top)
  } else {
    return middle, true
  }
}

func main() {
  arr := []int{1, 2, 3, 4, 4, 4, 4, 5, 10, 100, 101, 105, 120, 140}
  index, exist := BSearch(100, arr, 0, len(arr)-1)
  Println(arr)
  Println("isExist = ", exist, "index = ", index)
}
