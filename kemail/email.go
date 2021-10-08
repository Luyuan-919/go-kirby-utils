package kemail

import (
	"github.com/go-gomail/gomail"
)

const (
	fromName    = "From"
	sendTo      = "To"
	sendSubject = "subject"
	sendBody    = "text/html"
	defServerName = "sta-golang"
	defHost = "smtpdm.aliyun.com"
	defPort = 25
)

type EmailConfig struct {
	Host        string `json:"host" yaml:"host"`
	Port        int    `json:"port" yaml:"port"`
	Email       string `json:"kemail" yaml:"kemail"`
	Pwd         string `json:"pwd" yaml:"pwd"`
	ServerName  string `json:"serverName" yaml:"serverName"`
	ContentType string `json:"contentType" yaml:"contentType"`
}

type EmailClient struct {
	cfg    *EmailConfig
	helper *gomail.Dialer
}


func NewEmailConfig(Host, Email, Pwd string, Port int) (cfg *EmailConfig) {
	email :=  EmailConfig{
		Host: Host,
		Port: Port,
		Email: Email,
		Pwd: Pwd,
		ServerName: "",
		ContentType: "",
	}
	return &email
}

func NewEmailClient(cfg *EmailConfig) *EmailClient {
	if cfg.ServerName == "" {
		cfg.ServerName = defServerName
	} else if cfg.ContentType == "" {
		cfg.ContentType = sendBody
	}
	return &EmailClient{
		cfg:    cfg,
		helper: gomail.NewDialer(cfg.Host, cfg.Port, cfg.Email, cfg.Pwd),
	}
}

func (es *EmailClient) newMessage(subject, body string, mailTo ...string) *gomail.Message {
	message := gomail.NewMessage()
	message.SetHeader(fromName, es.cfg.Email, es.cfg.ServerName)
	message.SetHeader(sendTo, mailTo...)
	message.SetHeader(sendSubject, subject)
	message.SetBody(sendBody, body)
	return message
}

func (es *EmailClient) SendEmail(subject, body string, mailTo ...string) error {
	err := es.helper.DialAndSend(es.newMessage(subject, body, mailTo...))
	if err != nil {
		return err
	}
	return nil
}
