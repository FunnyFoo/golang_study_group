package main

import (
  . "fmt"
  "reflect"
)

// struct define
type DoublyLinkedList struct {
  head *Node
  tail *Node
  len int
}

type Node struct {
  data int
  next *Node
  prev *Node
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
func New() *DoublyLinkedList {
  l := new(DoublyLinkedList)
  l.head = nil
  l.tail = nil
  l.len = 0
  return l
}

func (l *DoublyLinkedList) Traverse(cb Callback) {
  current := l.head

  for current != nil {
    cb(current)
    current = current.next
  }
}

func (l *DoublyLinkedList) Add(data int) *Node {
  node := new(Node)
  node.data = data
  node.next = nil

  if(l.head == nil) {
    node.prev = nil
    l.head = node
    l.tail = node
  } else {
    l.tail.next = node
    node.prev = l.tail
    l.tail = node
  }
  l.len++

  return node
}

func (l *DoublyLinkedList) InsertAfter(data int, markNode *Node) *Node {

  node := new(Node)
  node.data = data

  if markNode != nil {
    if markNode == l.tail {
      node.next = nil
      node.prev = markNode
      l.tail = node
      markNode.next = node
    } else {
      node.next = markNode.next
      node.prev = markNode
      markNode.next.prev = node
      markNode.next = node
    }
    l.len++
  }

  return node
}

func (l *DoublyLinkedList) InsertAfterFromData(data int, markData int) []*Node {

  markNodes := l.Search(markData)
  results := make([]*Node, 0)

  for _, node := range markNodes {
    insertNode := l.InsertAfter(data, node)
    results = append(results, insertNode)
  }

  return results
}

func (l *DoublyLinkedList) Remove(node *Node) {

  l.Traverse(func(current *Node) {
    if current == node {
      switch {
      case node == l.head:
        l.head = l.head.next
        l.head.prev = nil
      case node == l.tail:
        l.tail = l.tail.prev
        l.tail.next = nil
      default:
        current.prev.next = current.next
        current.next.prev = current.prev
      }
      l.len--
    }
  })

  // free memory
  node.Reset()
  node = nil
}

func (l *DoublyLinkedList) RemoveFromData(data int) {
  searchResults := l.Search(data)

  for _, node := range searchResults {
    l.Remove(node)
  }
}

func (l *DoublyLinkedList) Search(data int) []*Node {
  results := make([]*Node, 0) // results := []*Node{}

  l.Traverse(func(current *Node) {
    if(current.data == data) {
      results = append(results, current)
    }
  })

  return results
}

func (l *DoublyLinkedList) Len() int {
  return l.len
}

func (l *DoublyLinkedList) Print() {
  arr := make([]int, 0)

  l.Traverse(func(current *Node) {
    arr = append(arr, current.data)
  })

  Println(arr)
}

func main() {
  var l = new(DoublyLinkedList)
  one := l.Add(1)
  two := l.Add(2)
  l.Add(2)
  l.Add(2)
  l.Add(2)
  l.Add(2)
  three := l.Add(3)
  l.Add(2)
  l.Print()
  l.Remove(three)
  l.Print()
  // l.RemoveFromData(2)
  l.Print()
  ten := l.InsertAfter(10, one)
  l.InsertAfterFromData(999, 2)
  l.Print()

  Println(l)
  Println(one, two, three, ten)
}
