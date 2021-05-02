package hmap

import (
	"fmt"
	"testing"
)

type Elements struct {
	value string
}

func (e Elements) GetKey() string {
	return e.value
}

func TestSortMap(t *testing.T) {
	fmt.Println("Starting test...")
	ml := NewMapList()
	var a, b, c Keyer
	a = &Elements{"Alice"}
	b = &Elements{"Bob"}
	c = &Elements{"Conrad"}
	ml.Push(a)
	ml.Push(b)
	ml.Push(c)
	cb := func(data Keyer) {
		fmt.Println(ml.dataMap[data.GetKey()].Value.(*Elements).value)
	}
	fmt.Println("Print elements in the order of pushing:")
	ml.Walk(cb)
	fmt.Printf("Size of MapList: %d \n", ml.Size())
	ml.Remove(b)
	fmt.Println("After removing b:")
	ml.Walk(cb)
	fmt.Printf("Size of MapList: %d \n", ml.Size())
}
