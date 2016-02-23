package custom_sort

import (
  // . "fmt"
)

func Insert(arr []int, index int) []int {
  if index==0 {
    return arr
  }

  // Insert the first value of unsorted array
  temp := arr[index]
  for i:=index-1;i>=0;i-- {
    if arr[i] > temp {
      arr[i+1] = arr[i]
      if i==0 {
        arr[i] = temp
      }
    } else {
      arr[i+1] = temp
      break
    }
  }
  return arr
}

func InsertionSort(arr []int) []int {
  for i:=1;i<len(arr);i++ {
    Insert(arr, i)
  }
  return arr
}

// func main() {
//   arr := []int{5, 5, 7, 4, 5, 6, 7, 10, 21, 5, 6, 9}
//   Println("arr", arr)
//   InsertionSort(arr)
//   Println("arr", arr)
// }
