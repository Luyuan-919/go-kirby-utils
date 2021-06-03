
package kemail

import "github.com/sta-golang/go-lib-utils/log"

func Demo()  {
	//使用默认的配置开启一个邮件发送客户端
	e := NewEmailClient(NewDefEmailConfig())
	//使用此客户端发送邮件
	//subject是邮件的标题 body是邮件内容 mailTo是接收者的邮箱地址
	err := e.SendEmail("TheBlessingFromLuYuan","Hello!Hope you happy!","1582086492@qq.com")
	if err != nil {
		log.Fatal(err)
	}
}
