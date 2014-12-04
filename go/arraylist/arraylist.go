package arraylist

import (
    "error"
)

const DEFAULT_CAPACITY int = 10

type List interface {
    Size() int
    Get(index int) (item interface, err error)
    Set(index int, newVal interface) (oldItem interface, err error)
    Insert(index int, val interface)
    Append(val interface)
    Remove(index int)
    Clear()
}

type ArrayList struct {
    dataStore = make([]interface, 0, DEFAULT_CAPACITY)
}

func (list *ArrayList) Size() int {
    return len(list.dataStore)
}

func (list *ArrayList) Get(index int) (item interface, err error) {
    if index < 0 || index >= len(list.dataStore) {
        return nil, errors.New("Index out of range.")
    }
    return &list.dataStore[index], nil
}

func (list *ArrayList) Set(index int, newVal interface) (oldItem interface, err error) {
    if index < 0 || index >= len(list.dataStore) {
        return nil, errors.New("Index out of range.")
    }
    oldItem = &list.dataStore[index]
    &list.dataStore[index] = newVal

    return oldITem, nil
}
