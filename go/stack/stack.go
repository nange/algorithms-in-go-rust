package stack

type Stack struct {
	dataStore []interface{}
	theSize   int
}

func New() *Stack {
	s := new(Stack)
	s.dataStore = make([]interface{}, 0)
	s.theSize = 0

	return s
}

func (s *Stack) Size() int {
	return s.theSize
}

func (s *Stack) IsEmpty() bool {
	return s.theSize == 0
}

func (s *Stack) Push(element interface{}) {
	s.dataStore = append(s.dataStore, element)
	s.theSize++
}

func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	last := s.dataStore[s.Size()-1]
	s.dataStore = s.dataStore[:s.Size()-1]
	s.theSize--
	return last
}

func (s *Stack) Peek() interface{} {
	if s.IsEmpty() {
		return nil
	}
	return s.dataStore[s.Size()-1]
}

func (s *Stack) Clear() {
	s.dataStore = s.dataStore[:1]
	s.theSize = 0
}
