package main

import (
  . "fmt"
)

type SinglyLinkedList struct {
  head *Node
  tail *Node
  len int
}

type Node struct {
  data int
  next *Node
}

// List methods
func New() *SinglyLinkedList {
  l := new(SinglyLinkedList)
  l.head = nil
  l.tail = nil
  l.len = 0
  return l
}

func (l *SinglyLinkedList) Add(data int) {
  node := new(Node)
  node.data = data
  node.next = nil

  if(l.head == nil) {
    l.head = node
    l.tail = node
  } else {
    l.tail.next = node
    l.tail = node
  }
  l.len++
}

func (l *SinglyLinkedList) InsertAfter(v int, markNode *Node) *Node {
  return nil
}

func (l *SinglyLinkedList) Remove() {

}

func (l *SinglyLinkedList) Len() {

}

func (l *SinglyLinkedList) Print() {
  arr := make([]int, 0)
  current := l.head
  for {
    if current == nil {
      break
    } else {
      arr = append(arr, current.data)
    }
    current = current.next
  }

  Println(arr)
}

func main() {
  var l = new(SinglyLinkedList)
  l.Print()
  l.Add(1)
  l.Add(2)
  l.Add(3)
  l.Print()
}
