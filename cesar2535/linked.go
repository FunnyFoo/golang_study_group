package cesar2535

import "fmt"

type Element struct {
  Value interface{}
  next, prev *Element
}

func (e *Element) Next() *Element {
  p := e.next
  return p
}

type SLinkedList struct {
  root Element
  len int
}

func (li *SLinkedList) Init() *SLinkedList {
  li.root.next = nil
  li.len = 0
  return li
}

func New() *SLinkedList {
  return new(SLinkedList).Init()
}

func (li *SLinkedList) Len() int {
  return li.len
}

func (li *SLinkedList) Front() *Element {
  if li.len == 0 {
    return nil
  }

  return li.root.next
}

func (li *SLinkedList) Back() *Element {
  if li.len == 0 {
    return nil
  }
  node := &li.root
  for node.next != nil {
    node = node.next
  }
  return node
}

func (li *SLinkedList) insert(elm *Element, at *Element) *Element {
  n := at.next
  at.next = elm
  elm.next = n
  li.len++
  return elm
}

func (li *SLinkedList) insertValue(val interface{}, at *Element) *Element {
  return li.insert(&Element{Value: val}, at)
}

func (li *SLinkedList) InsertBeginning(val interface{}) *Element {
  return li.insertValue(val, &li.root)
}

func (li *SLinkedList) InsertEnd(val interface{}) *Element {
  node := &li.root

  for node.next != nil {
    node = node.next
  }
  return li.insertValue(val, node)
}

func (li *SLinkedList) InsertAfter(val interface{}, at *Element) *Element {
  return li.insertValue(val, at)
}

func (li *SLinkedList) RemoveAfter(at *Element) *Element {
  target := at.next
  at.next = at.next.next
  target.next = nil
  return target
}

func (li *SLinkedList) search(val interface{}) *Element {
  node := &li.root
  for node != nil {
    if node.Value == val {
      return node
    }
    node = node.next
  }
  return nil
}
