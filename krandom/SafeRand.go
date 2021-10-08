package krandom

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"math/big"
)

func CreateRandomNumber(len int)  string{
	var numbers = []byte{0,1,2,3,4,5,6,7,8,9}
	var container string
	length := bytes.NewReader(numbers).Len()
	x := big.NewInt(int64(length))
	for i:=1;i<=len;i++{
		random,_ := rand.Int(rand.Reader,x)
		container += fmt.Sprintf("%d",numbers[random.Int64()])
	}
	return container
}

func CreateRandomString(n int) string  {
	var container string
	var str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	length := len(str)
	bigInt := big.NewInt(int64(length))
	for i := 0;i < n ;i++  {
		randomInt,_ := rand.Int(rand.Reader,bigInt)
		container += string(str[randomInt.Int64()])
	}
	return container
}