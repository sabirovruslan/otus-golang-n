package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	size     int
	headItem *ListItem
	tailItem *ListItem
}

func (l *list) Len() int {
	return l.size
}

func (l *list) Front() *ListItem {
	return l.headItem
}

func (l *list) Back() *ListItem {
	return l.tailItem
}

func (l *list) PushFront(v interface{}) *ListItem {
	i := &ListItem{v, nil, nil}

	if l.Len() == 0 {
		l.headItem, l.tailItem = i, i
	} else {
		i.Next, l.headItem.Prev = l.headItem, i
		l.headItem = i
	}
	l.size++

	return i
}

func (l *list) PushBack(v interface{}) *ListItem {
	i := &ListItem{v, nil, nil}

	if l.Len() == 0 {
		l.headItem, l.tailItem = i, i
	} else {
		i.Prev, l.tailItem.Next = l.tailItem, i
		l.tailItem = i
	}
	l.size++

	return i
}

func (l *list) Remove(i *ListItem) {
	switch {
	case i.Next == nil && i.Prev == nil:
		l.headItem, l.tailItem = nil, nil
	case i.Next == nil:
		i.Prev.Next = nil
		l.tailItem = i.Prev
	case i.Prev == nil:
		i.Next.Prev = nil
		l.headItem = i.Next
	default:
		i.Prev.Next, i.Next.Prev = i.Next, i.Prev
	}

	l.size--
}

func (l *list) MoveToFront(i *ListItem) {
	if (i.Next == nil && i.Prev == nil) || (i.Prev == nil) {
		return
	}

	l.Remove(i)
	i.Prev, i.Next, l.headItem.Prev = nil, l.headItem, i
	l.headItem = i
	l.size++
}

func NewList() List {
	return new(list)
}
