package main

import (
  "fmt"
)

type Queue struct {
  arr []int
  front int
  end int
}

func createQueue(length int) *Queue{
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

func (q *Queue) pop() int {

  if q.front == q.end {
    fmt.Println("The queue is empty.")
    return -1
  }

  // pop
  q.front++
  return q.arr[q.front]

}

func (q *Queue) push(v int) bool {

  if (q.end-q.front) == len(q.arr) {
    fmt.Println("The queue is full.")
    return false
  }

  // reform condition
  if q.end == (len(q.arr) -1) {
    fmt.Println("Do reform")
    q.reform()
  }

  // push data
  q.end++
  q.arr[q.end] = v

  return true

}

func (q *Queue) reform() *Queue  {
  var i int
  length := q.end - q.front

  for i=0;i<length;i++ {
    q.arr[i] = q.arr[q.front + 1 + i]
    fmt.Println(q.arr)
  }

  q.front = -1
  q.end = q.front + length

  return q

}

func main()  {
  q := createQueue(3)

  q.pop()
  q.push(1)
  q.push(2)
  q.push(3)
  q.push(4)
  fmt.Println("Queue", q)
  fmt.Println("Pop", q.pop(), "Queue", q)
  fmt.Println("Pop", q.pop(), "Queue", q)
  q.push(5)
  fmt.Println("Queue", q)
  q.push(6)
  fmt.Println("Queue", q)
  fmt.Println("Pop", q.pop(), "Queue", q)
  fmt.Println("Pop", q.pop(), "Queue", q)
  fmt.Println("Pop", q.pop(), "Queue", q)
  fmt.Println("Pop", q.pop(), "Queue", q)

  fmt.Println("end")
}
