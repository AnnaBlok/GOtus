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
	len   int
	front *ListItem
	back  *ListItem
}

func NewListItem(v interface{}, next, prev *ListItem) *ListItem {
	return &ListItem{v, next, prev}
}

func NewList() List {
	return new(list)
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.front
}

func (l *list) Back() *ListItem {
	return l.back
}

func (l *list) PushFront(v interface{}) *ListItem {
	if l.front != nil {
		l.front = NewListItem(v, l.front, nil)
	} else {
		l.front = NewListItem(v, l.back, nil)
	}
	if l.front.Next != nil {
		l.front.Next.Prev = l.front
	}
	l.len++
	if l.len == 2 && l.back == nil {
		l.back = l.front.Next
	}

	return l.front
}

func (l *list) PushBack(v interface{}) *ListItem {
	if l.back != nil {
		l.back = NewListItem(v, nil, l.back)
	} else {
		l.back = NewListItem(v, nil, l.front)
	}
	if l.back.Prev != nil {
		l.back.Prev.Next = l.back
	}
	l.len++
	if l.len == 2 && l.front == nil {
		l.front = l.back.Prev
	}

	return l.back
}

func (l *list) Remove(i *ListItem) {
	if i.Prev != nil {
		i.Prev.Next = i.Next
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	}
	if i == l.front {
		l.front = l.front.Next
	} else if i == l.back {
		l.back = l.back.Prev
	}
	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	l.Remove(i)
	l.PushFront(i.Value)
}
