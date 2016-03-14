package main

import (
  . "fmt"
  "reflect"
)

// struct define
type SinglyLinkedList struct {
  head *Node
  tail *Node
  len int
}

type Node struct {
  data int
  next *Node
}

// func define
type Callback func(*Node)

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

func (l *SinglyLinkedList) Traverse(cb Callback) {
  current := l.head

  for current != nil {
    cb(current)
    current = current.next
  }
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

func (l *SinglyLinkedList) InsertAfter(data int, markNode *Node) *Node {

  node := new(Node)
  node.data = data

  if markNode != nil {
    if markNode == l.tail {
      markNode.next = node
      node.next = nil
      l.tail = node
    } else {
      node.next = markNode.next
      markNode.next = node
    }
    l.len++
  }

  return node
}

func (l *SinglyLinkedList) InsertAfterFromData(data int, markData int) []*Node {

  markNodes := l.Search(markData)
  results := make([]*Node, 0)

  for _, node := range markNodes {
    insertNode := l.InsertAfter(data, node)
    results = append(results, insertNode)
  }

  return results
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
  node = nil
}

func (l *SinglyLinkedList) RemoveFromData(data int) {
  searchResults := l.Search(data)

  for _, node := range searchResults {
    l.Remove(node)
  }
}

func (l *SinglyLinkedList) Search(data int) []*Node {
  results := make([]*Node, 0) // results := []*Node{}

  l.Traverse(func(current *Node) {
    if(current.data == data) {
      results = append(results, current)
    }
  })

  return results
}

func (l *SinglyLinkedList) Len() int {
  return l.len
}

func (l *SinglyLinkedList) Print() {
  arr := make([]int, 0)

  l.Traverse(func(current *Node) {
    arr = append(arr, current.data)
  })

  Println(arr)
}

func main() {
  var l = new(SinglyLinkedList)
  l.Print()
  one := l.Add(1)
  two := l.Add(2)
  l.Add(2)
  three := l.Add(3)
  l.Add(2)
  l.Add(2)
  l.Print()
  l.Remove(three)
  l.Print()
  l.RemoveFromData(2)
  l.Print()
  four := l.Add(4)
  two = l.Add(2)
  l.Add(2)
  l.Add(2)
  l.Print()
  Println(l.Len())
  l.InsertAfterFromData(9, 2)
  l.Print()

  Println(one, two, three, four)
}