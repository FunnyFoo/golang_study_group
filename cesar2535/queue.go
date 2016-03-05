package cesar2535

import "fmt"

type Node struct {
  Value int
}

func (n *Node) ToString() string {
  return fmt.Sprint(n.Value)
}

type Queue []*Node

func (q *Queue) Push(n *Node) {
  *q = append(*q, n)
}

func (q *Queue) Pop() *Node {
  n := (*q)[0]
  *q = (*q)[1:]
  return n
}

func (q *Queue) Len() int {
  return len(*q)
}

type CircularQueue struct {
  nodes []*Node
  size int
  head int
  tail int
  count int
}

func (q *CircularQueue) Push(n *Node) {
  if q.head == q.tail && q.count > 0 {
    nodes := make([]*Node, len(q.nodes) + q.size)
    copy(nodes, q.nodes[q.head:])
    copy(nodes[len(q.nodes) - q.head:], q.nodes[:q.head])
    q.head = 0
    q.tail = len(q.nodes)
    q.nodes = nodes
  }
  q.nodes[q.tail] = n
  q.tail = (q.tail + 1) % len(q.nodes)
  q.count++
}

func (q *CircularQueue) Pop() *Node {
  if q.count == 0 {
    return nil
  }
  node := q.nodes[q.head]
  q.head = (q.head + 1) % len(q.nodes)
  q.count--
  return node
}

func (q *CircularQueue) Len() int {
  return len(q.nodes)
}
