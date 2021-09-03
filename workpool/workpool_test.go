package workpool

import (
	"fmt"
	"testing"
)

func TestWorkPool(t *testing.T) {
	num := 100
	wp := New(num)

	for i := 0; i < num; i++ {
		tmp := i
		wp.Do(func() error {
			fmt.Println(tmp)
			return nil
		})
	}

	err := wp.Wait()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("done")
}
