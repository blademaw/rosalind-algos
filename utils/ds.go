package utils

type Queue struct {
  elements []interface{}
}

func (q *Queue) Enqueue(item interface{}) {
  q.elements = append(q.elements, item)
}

func (q *Queue) Dequeue() interface{} {
  if len(q.elements) == 0 {
    return nil
  }
  item := q.elements[0]
  q.elements = q.elements[1:]
  return item
}

func (q *Queue) IsEmpty() bool {
  return len(q.elements) == 0
}


// Implements a generic stack interface
type Stack struct {
  elements []interface{}
}

func (s *Stack) Push(item interface{}) {
  s.elements = append([]interface{}{item}, s.elements...)
}

func (s *Stack) Pop() interface{} {
  if len(s.elements) == 0 {
    return nil
  }
  item := s.elements[0]
  s.elements = s.elements[1:]
  return item
}

func (s *Stack) IsEmpty() bool {
  return len(s.elements) == 0
}
