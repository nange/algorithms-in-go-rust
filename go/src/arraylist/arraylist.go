package arraylist

import (
	"errors"
	"fmt"
)

const DEFAULT_CAPACITY int = 10

type List interface {
	Size() int
	Get(index int) (interface{}, error)
	Set(index int, newVal interface{}) error
	Insert(index int, val interface{}) error
	Append(val interface{})
	Remove(index int) error
	Clear()
	String()
}

type Iterator interface {
	HasNext() bool
	Next() (interface{}, error)
	Remove()
}

type Iterable interface {
	Iterator() Iterator
}

type arrayListIterator struct {
	list    *ArrayList
	current int
}

type ArrayList struct {
	dataStore []interface{}
	theSize   int
}

func New() *ArrayList {
	list := new(ArrayList)
	list.Clear()
	return list
}

func (list *ArrayList) Size() int {
	return list.theSize
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

func (list *ArrayList) Append(val interface{}) {
	list.dataStore = append(list.dataStore, val)
	list.theSize++
}

func (list *ArrayList) Insert(index int, val interface{}) error {
	if index < 0 || index >= list.Size() {
		return errors.New("Index out of range.")
	}
	list.ensureCapacity()
	list.dataStore = list.dataStore[:list.Size()+1]

	for i := list.Size(); i > index; i-- {
		list.dataStore[i] = list.dataStore[i-1]
	}
	list.dataStore[index] = val
	list.theSize++

	return nil
}

func (list *ArrayList) Remove(index int) error {
	if index < 0 || index >= list.Size() {
		return errors.New("Index out of range.")
	}
	list.dataStore = append(list.dataStore[:index], list.dataStore[index+1:]...)
	list.theSize--
	return nil
}

func (list *ArrayList) Clear() {
	list.dataStore = make([]interface{}, 0, DEFAULT_CAPACITY)
	list.theSize = 0
}

func (list *ArrayList) ensureCapacity() {
	if list.Size() == cap(list.dataStore) {
		newDS := make([]interface{}, 0, 2*list.Size())
		copy(newDS, list.dataStore)
		list.dataStore = newDS
	}
}

func (list *ArrayList) String() string {
	return fmt.Sprint(list.dataStore)
}

func (list *ArrayList) Iterator() Iterator {
	iterator := new(arrayListIterator)
	iterator.current = 0
	iterator.list = list
	return iterator
}

func (it *arrayListIterator) HasNext() bool {
	return it.current < it.list.Size()
}

func (it *arrayListIterator) Next() (interface{}, error) {
	if !it.HasNext() {
		return nil, errors.New("No Such Element.")
	}
	v, err := it.list.Get(it.current)
	it.current++
	return v, err
}

func (it *arrayListIterator) Remove() {
	it.current--
	it.list.Remove(it.current)
}
