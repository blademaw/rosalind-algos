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


// Implements a min-heap-based priority queue for integers.
type PriorityQueue struct {
  items []int
  keys  map[int]int
}

func NewPriorityQueue() *PriorityQueue {
  return &PriorityQueue{
    items: []int{},           // Items
    keys:  make(map[int]int), // Items -> Keys
  }
}

func (pq *PriorityQueue) IsEmpty() bool {
  return len(pq.items) == 0
}

func (pq *PriorityQueue) Insert(item int, key int) {
  pq.items = append(pq.items, item)
  pq.keys[item] = key
  pq.upHeap(len(pq.items) - 1)
}

func (pq *PriorityQueue) PopMin() (int, int, bool) {
  if len(pq.items) == 0 {
    return 0, 0, false
  }

  minElem := pq.items[0]
  minKey  := pq.keys[minElem]

  pq.items[0] = pq.items[len(pq.items) - 1]
  pq.items = pq.items[:len(pq.items) - 1]
  delete(pq.keys, minElem)
  pq.downHeap(0)

  return minElem, minKey, true
}

func (pq *PriorityQueue) upHeap(index int) {
  for pq.keys[pq.items[index]] < pq.keys[pq.items[(index-1)/2]] {
    pq.items[index], pq.items[(index-1)/2] = pq.items[(index-1)/2], pq.items[index]
    index = (index - 1) / 2
  }
}

func (pq *PriorityQueue) downHeap(index int) {
  lastInd := len(pq.items) - 1
  l, r := 2*index + 1, 2*index + 2
  child := 0

  for l <= lastInd {
    if l == lastInd {
      child = l
    } else if pq.keys[pq.items[l]] < pq.keys[pq.items[r]] {
      child = l
    } else {
      child = r
    }

    if pq.keys[pq.items[index]] > pq.keys[pq.items[child]] {
      pq.items[index], pq.items[child] = pq.items[child], pq.items[index]
      index = child
      l, r  = 2*index + 1, 2*index + 2
    } else {
      return
    }
  }
}
