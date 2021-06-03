package md5

import (
	"fmt"
	"testing"
)

func TestGetMd5(t *testing.T) {
	password := "kkirby"
	fmt.Println(GetMd5(password))
}
