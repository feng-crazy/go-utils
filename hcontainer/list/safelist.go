package list

import (
	"container/list"
	"sync"
)

type SafeList struct {
	sync.RWMutex
	L *list.List
}

func NewSafeList() *SafeList {
	return &SafeList{L: list.New()}
}

func (l *SafeList) PushFront(v interface{}) *list.Element {
	l.Lock()
	e := l.L.PushFront(v)
	l.Unlock()
	return e
}

func (l *SafeList) PushFrontBatch(vs []interface{}) {
	l.Lock()
	for _, item := range vs {
		l.L.PushFront(item)
	}
	l.Unlock()
}

func (l *SafeList) PopBack() interface{} {
	l.Lock()

	if elem := l.L.Back(); elem != nil {
		item := l.L.Remove(elem)
		l.Unlock()
		return item
	}

	l.Unlock()
	return nil
}

func (l *SafeList) PopBackBy(max int) []interface{} {
	l.Lock()

	count := l.len()
	if count == 0 {
		l.Unlock()
		return []interface{}{}
	}

	if count > max {
		count = max
	}

	items := make([]interface{}, 0, count)
	for i := 0; i < count; i++ {
		item := l.L.Remove(l.L.Back())
		items = append(items, item)
	}

	l.Unlock()
	return items
}

func (l *SafeList) PopBackAll() []interface{} {
	l.Lock()

	count := l.len()
	if count == 0 {
		l.Unlock()
		return []interface{}{}
	}

	items := make([]interface{}, 0, count)
	for i := 0; i < count; i++ {
		item := l.L.Remove(l.L.Back())
		items = append(items, item)
	}

	l.Unlock()
	return items
}

func (l *SafeList) Remove(e *list.Element) interface{} {
	l.Lock()
	defer l.Unlock()
	return l.L.Remove(e)
}

func (l *SafeList) RemoveAll() {
	l.Lock()
	l.L = list.New()
	l.Unlock()
}

func (l *SafeList) FrontAll() []interface{} {
	l.RLock()
	defer l.RUnlock()

	count := l.len()
	if count == 0 {
		return []interface{}{}
	}

	items := make([]interface{}, 0, count)
	for e := l.L.Front(); e != nil; e = e.Next() {
		items = append(items, e.Value)
	}
	return items
}

func (l *SafeList) BackAll() []interface{} {
	l.RLock()
	defer l.RUnlock()

	count := l.len()
	if count == 0 {
		return []interface{}{}
	}

	items := make([]interface{}, 0, count)
	for e := l.L.Back(); e != nil; e = e.Prev() {
		items = append(items, e.Value)
	}
	return items
}

func (l *SafeList) Front() interface{} {
	l.RLock()

	if f := l.L.Front(); f != nil {
		l.RUnlock()
		return f.Value
	}

	l.RUnlock()
	return nil
}

func (l *SafeList) Len() int {
	l.RLock()
	defer l.RUnlock()
	return l.len()
}

func (l *SafeList) len() int {
	return l.L.Len()
}

// SafeList with Limited Size
type SafeListLimited struct {
	maxSize int
	SL      *SafeList
}

func NewSafeListLimited(maxSize int) *SafeListLimited {
	return &SafeListLimited{SL: NewSafeList(), maxSize: maxSize}
}

func (l *SafeListLimited) PopBack() interface{} {
	return l.SL.PopBack()
}

func (l *SafeListLimited) PopBackBy(max int) []interface{} {
	return l.SL.PopBackBy(max)
}

func (l *SafeListLimited) PushFront(v interface{}) bool {
	if l.SL.Len() >= l.maxSize {
		return false
	}

	l.SL.PushFront(v)
	return true
}

func (l *SafeListLimited) PushFrontBatch(vs []interface{}) bool {
	if l.SL.Len() >= l.maxSize {
		return false
	}

	l.SL.PushFrontBatch(vs)
	return true
}

func (l *SafeListLimited) PushFrontViolently(v interface{}) bool {
	l.SL.PushFront(v)
	if l.SL.Len() > l.maxSize {
		l.SL.PopBack()
	}

	return true
}

func (l *SafeListLimited) RemoveAll() {
	l.SL.RemoveAll()
}

func (l *SafeListLimited) Front() interface{} {
	return l.SL.Front()
}

func (l *SafeListLimited) FrontAll() []interface{} {
	return l.SL.FrontAll()
}

func (l *SafeListLimited) Len() int {
	return l.SL.Len()
}
