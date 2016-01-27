package main

import (
  "fmt"
  "math"
)

type Queue struct {
  arr []int
  front int
  end int
}

func CreateQueue(length int) *Queue{
  q := new (Queue)
  q.arr = make([]int, length)
  q.front = 0
  q.end = 0

  // initial value to "-1"
  var i int
  for i=0;i < len(q.arr);i++ {
    q.arr[i] = -1
  }

  return q
}

func (q *Queue) Pop() int {

  if q.front == q.end {
    fmt.Println("The queue is empty.")
    return -1
  }

  q.front = (q.front + 1) % len(q.arr)
  return q.arr[q.front]

}

func (q *Queue) Push(v int) bool {

  if q.front == (q.end + 1)%len(q.arr) {
    fmt.Println("The queue is full.")
    return false
  }

  // Push data
  q.end = (q.end + 1) % len(q.arr)
  q.arr[q.end] = v

  return true
}

func (q *Queue) GetQueue() []int {
  if(q.front == q.end) {
    return make([]int, 0)
  }

  var i int
  length := q.Len()
  arr := make([]int, length)
  for i=0;i<length;i++ {
    arr[i] = q.arr[(q.front + 1 + i) % len(q.arr)]
  }

  return arr
}

func (q *Queue) Len() int {
  if q.end >= q.front {
    return q.end - q.front
  } else {
    return len(q.arr) - (int)(math.Abs((float64)(q.front-q.end)))
  }
}

func main()  {
  q := CreateQueue(3)
  q.Pop()
  fmt.Println("Push 1", q.Push(1), "Queue", q, "Length", q.Len(), "GetQueue", q.GetQueue())
  fmt.Println("Push 2", q.Push(2), "Queue", q, "Length", q.Len(), "GetQueue", q.GetQueue())
  fmt.Println("Queue", q, "Length", q.Len(), "GetQueue", q.GetQueue())
  fmt.Println("Push 3", q.Push(3), "Queue", q, "Length", q.Len(), "GetQueue", q.GetQueue())
  fmt.Println("Queue", q, "Length", q.Len(), "GetQueue", q.GetQueue())
  fmt.Println("Pop:", q.Pop(), "Queue", q, "GetQueue", q.GetQueue())
  fmt.Println("Pop:", q.Pop(), "Queue", q, "GetQueue", q.GetQueue())

  fmt.Println("Push 4", q.Push(4), "Queue", q, "Length", q.Len(), "GetQueue", q.GetQueue())
  fmt.Println("Pop:", q.Pop(), "Queue", q, "GetQueue", q.GetQueue())

  fmt.Println("Push 5", q.Push(5), "Queue", q, "Length", q.Len(), "GetQueue", q.GetQueue())

  fmt.Println("Pop:", q.Pop(), "Queue", q, "GetQueue", q.GetQueue())

  fmt.Println("Push 7", q.Push(7), "Queue", q, "Length", q.Len(), "GetQueue", q.GetQueue())

  fmt.Println("Push 8", q.Push(8), "Queue", q, "Length", q.Len(), "GetQueue", q.GetQueue())


  fmt.Println("end")
}
