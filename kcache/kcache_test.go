package kcache

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"go-Kirby-utils/kMemory"
	"sync"
	"testing"
	"time"
)

type String string
type User struct {
	UserName string
	PassWord string
	Email string
}

func (s String) Size() int {
	return kMemory.SizeofMemoryInt(s)
}

var (
	Db *gorm.DB
	err error
	times1 int = 0
	times2 int = 0
)

func CheckUser(user *User) bool {
	pwd := user.PassWord
	var u User
	err := Db.Where("user_name = ? AND pass_word = ?",user.UserName,pwd).Find(&u)
	if err.Error != nil {
		return false
	}else {
		times1++
		return true
	}
}

func TestNewKcache(t *testing.T) {
	Db, err = gorm.Open( "mysql","root:123456@tcp(localhost:3306)/chat?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	now := time.Now()
	for i := 0; i < 1000; i++ {
		CheckUser(&User{
			UserName: "kirby3",
			PassWord: "drink",
		})
	}
	fmt.Println("mysql 10000次 用时：",time.Since(now),"成功次数：",times1)
	k := NewKcacheWithMaxMemory(64,func(s string, value Value) {
		fmt.Println("Kcache delete ",s,"value is ",value)
	})
	k.Add("kirby3",String("drink"))
	f := func() {
		if _,ok := k.Get("kirby3");ok{
			times2++
		}
	}
	now = time.Now()
	for i := 0; i < 100000; i++ {
		go f()
	}
	fmt.Println("kcache 100000次 用时：",time.Since(now),"成功次数：",times2)
}

func TestKcache_Add(t *testing.T) {
	k1 := NewKcache(nil)
	k1.Add("1",String(rune(1)))
	//mu := sync.Mutex{}
	w := sync.WaitGroup{}
	w.Add(10000)
	k := NewKcacheWithMaxMemory(64,func(s string, value Value) {
		fmt.Println("Kcache delete ",s,"value is ",value)
	})
	k.Add("kirby3",String("drink"))
	now := time.Now()
	for i := 0; i < 10000; i++ {
		go func() {
			defer w.Done()
			//mu.Lock()
			if _,ok := k.Get("kirby3") ;ok{
				times2++
			}
			//mu.Unlock()
		}()
	}
	w.Wait()
	fmt.Println("kcache 100000000次 用时：",time.Since(now),"成功次数：",times2)
}
