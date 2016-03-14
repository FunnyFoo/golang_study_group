package main

import (
  . "fmt"
  "reflect"
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

// Other methods
func clear(v interface{}) {
  p := reflect.ValueOf(v).Elem()
  p.Set(reflect.Zero(p.Type()))
}

// Node methods
func (node *Node) Reset() {
  clear(node)
}

// List methods
func New() *SinglyLinkedList {
  l := new(SinglyLinkedList)
  l.head = nil
  l.tail = nil
  l.len = 0
  return l
}

func (l *SinglyLinkedList) Add(data int) *Node {
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

  return node
}

func (l *SinglyLinkedList) InsertAfter(v int, markNode *Node) *Node {
  return nil
}

func (l *SinglyLinkedList) Remove(node *Node) {
  current := l.head
  previous := l.head

  for current != nil {
    if current == node {
      switch {
      case node == l.head:
        l.head = l.head.next
      case node == l.tail:
        l.tail = previous
        l.tail.next = nil
      default:
        previous.next = current.next
      }
      l.len--
    }

    previous = current
    current = current.next

  }

  // free memory
  node.Reset()

}

func (l *SinglyLinkedList) RemoveFromData(data int) {

}

func (l *SinglyLinkedList) Len() int {
  return l.len
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
  one := l.Add(1)
  two := l.Add(2)
  three := l.Add(3)
  l.Print()
  Println(one, two, three)
  l.Remove(three)
  l.Print()
  l.Remove(two)
  l.Print()
}
