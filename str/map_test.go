package str

import (
	"fmt"
	"testing"
)

func TestSingleStructJsonTagToStrArray(t *testing.T) {

	type Student struct {
		Name  string `json:"name"`
		Age   int    `json:"age"`
		Id    int    `gorm:"column:id;not null;primaryKey;type:bigserial;commnet:'主键'"`
		Score string
	}

	ss, err := SingleStructJsonTagToStrArray(Student{})
	fmt.Println(ss, err)
}
