package main

type Stack struct {
	Tail *Node
}

type Node struct {
	Value string
	Next  *Node
}

func (s *Stack) Push(value string) {
	node := Node{Value: value}

	if s.Tail != nil {
		node.Next = s.Tail
	}
	s.Tail = &node
}

func (s *Stack) Pop() string {
	if s.Tail == nil {
		return ""
	}
	node := s.Tail
	s.Tail = node.Next

	return node.Value
}
