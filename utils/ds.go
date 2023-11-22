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
