package codef

import (
	"fmt"
	"testing"
)

func TestIso2Utf8(t *testing.T) {
	result := UTF8Convert([]byte("I P 1 . N 5p��ψh . A I - 1 4 0"))
	fmt.Println(result)
}
