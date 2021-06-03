package md5

import (
	"crypto/md5"
	"fmt"
	"io"
)

func GetMd5(data string) string {
	h := md5.New()
	_, _ = io.WriteString(h,data)
	return fmt.Sprintf("%x",h.Sum(nil))
}