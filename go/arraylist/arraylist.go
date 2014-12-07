package arraylist

import "errors"

const DEFAULT_CAPACITY int = 10

type List interface {
	Size() int
	Get(index int) (interface{}, error)
	Set(index int, newVal interface{}) error
	Insert(index int, val interface{}) error
	Append(val interface{})
	Remove(index int) error
	Clear()
}

type ArrayList struct {
	dataStore []interface{}
}

func New() *ArrayList {
	return &ArrayList{
		dataStore: make([]interface{}, 0, DEFAULT_CAPACITY),
	}
}

func (list *ArrayList) Size() int {
	return len(list.dataStore)
}

func (list *ArrayList) Get(index int) (interface{}, error) {
	if index < 0 || index >= list.Size() {
		return nil, errors.New("Index out of range.")
	}
	return list.dataStore[index], nil
}

func (list *ArrayList) Set(index int, newVal interface{}) error {
	if index < 0 || index >= list.Size() {
		return errors.New("Index out of range.")
	}
	list.dataStore[index] = newVal

	return nil
}

func (list *ArrayList) Insert(index int, val interface{}) error {
	if index < 0 || index >= list.Size() {
		return errors.New("Index out of range.")
	}
	list.ensureCapacity()

	for i := list.Size() - 1; i > index; i-- {
		list.dataStore[i] = list.dataStore[i-1]
	}
	list.dataStore[index] = val

	return nil
}

func (list *ArrayList) Append(val interface{}) {
	list.dataStore = append(list.dataStore, val)
}

func (list *ArrayList) Remove(index int) error {
	if index < 0 || index >= list.Size() {
		return errors.New("Index out of range.")
	}
	beforeIndex := list.dataStore[:index]
	afterIndex := list.dataStore[index+1:]

	list.dataStore = make([]interface{}, 0, cap(list.dataStore))
	list.dataStore = append(list.dataStore, beforeIndex...)
	list.dataStore = append(list.dataStore, afterIndex...)

	return nil
}

func (list *ArrayList) Clear() {
	list.dataStore = make([]interface{}, 0, DEFAULT_CAPACITY)
}

func (list *ArrayList) ensureCapacity() {
	if list.Size() == cap(list.dataStore) {
		newDS := make([]interface{}, list.Size()+1, 2*list.Size())
		copy(newDS, list.dataStore)
		list.dataStore = newDS
	}
}
