package main

import (
  "fmt"
)

type Queue struct {
  arr []int
  front int
  end int
}

func CreateQueue(length int) *Queue{
  q := new (Queue)
  q.arr = make([]int, length)
  q.front = -1
  q.end = -1

  // initial value to "-1"
  var i int
  for i=0;i<len(q.arr);i++ {
    q.arr[i] = -1
  }

  return q
}

func (q *Queue) Pop() int {

  if q.front == q.end {
    fmt.Println("The queue is empty.")
    return -1
  }

  // Pop
  q.front++
  return q.arr[q.front]

}

func (q *Queue) Push(v int) bool {

  if (q.end-q.front) == len(q.arr) {
    fmt.Println("The queue is full.")
    return false
  }

  // Reform condition
  if q.end == (len(q.arr) -1) {
    fmt.Println("Do Reform")
    q.Reform()
  }

  // Push data
  q.end++
  q.arr[q.end] = v

  return true

}

func (q *Queue) Reform() *Queue  {
  var i int
  length := q.Len()

  for i=0;i<length;i++ {
    q.arr[i] = q.arr[q.front + 1 + i]
  }

  q.front = -1
  q.end = q.front + length

  return q

}

func (q *Queue) GetQueue() []int {
  if(q.front == q.end) {
    return make([]int, 0)
  }

  // slice[a:b], mean a ~ b-1
  return q.arr[q.front+1:q.end+1]
}

func (q *Queue) Len() int {
  return q.end - q.front
}

func main()  {
  q := CreateQueue(3)
  q.Pop()
  fmt.Println("Push 1", q.Push(1), "Queue", q, "Length", q.Len(), "GetQueue", q.GetQueue())
  fmt.Println("Push 2", q.Push(2), "Queue", q, "Length", q.Len(), "GetQueue", q.GetQueue())
  fmt.Println("Push 3", q.Push(3), "Queue", q, "Length", q.Len(), "GetQueue", q.GetQueue())

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
