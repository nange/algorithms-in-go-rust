package queue

type Queue struct {
	dataStore []interface{}
	theSize   int
}

func New() *Queue {
	q := new(Queue)
	q.Clear()
	return q
}

func (q *Queue) Enqueue(element interface{}) {
	q.dataStore = append(q.dataStore, element)
	q.theSize++
}

func (q *Queue) Dequeue() interface{} {
	if q.Size() == 0 {
		return nil
	}
	d := q.dataStore[0]
	if q.Size() > 1 {
		q.dataStore = q.dataStore[1:q.theSize]
	}
	q.theSize--

	return d
}

func (q *Queue) Size() int {
	return q.theSize
}

func (q *Queue) IsEmpty() bool {
	return q.theSize == 0
}

func (q *Queue) Font() interface{} {
	if q.Size() == 0 {
		return nil
	}
	return q.dataStore[0]
}

func (q *Queue) End() interface{} {
	if q.Size() == 0 {
		return nil
	}

	return q.dataStore[q.theSize-1]
}

func (q *Queue) Clear() {
	q.dataStore = make([]interface{}, 0)
	q.theSize = 0
}
