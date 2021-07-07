package secret

import (
	"fmt"
	"testing"
)

func TestMd5EnCode(t *testing.T) {
	fmt.Println(Md5EnCode("123"))
}

func TestMd5Encryption(t *testing.T) {
	fmt.Println(Md5Encryption("123"))
}
