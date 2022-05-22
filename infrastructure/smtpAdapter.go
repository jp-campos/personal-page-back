package infrastructure

import (
	"fmt"
	"net/smtp"
	"personal-page-back/domain"
	"sync"
)

var wg sync.WaitGroup

type SmtpAdapter struct {
	hostName, from, password string
	auth                     smtp.Auth
}

func NewSmtpAdapter(hostName, from, password string) (s SmtpAdapter) {
	fmt.Println("SMTP created with", hostName, from)
	auth := smtp.PlainAuth("", from, password, hostName)
	return SmtpAdapter{auth: auth, hostName: hostName, from: from, password: password}
}

func (s SmtpAdapter) SendEmail(email domain.Email) error {
	
	msg := fmt.Sprintf("To: %s\r\n"+
		"Subject:%s\r\n"+
		"\r\n"+
		"%s \r\n", email.To, email.Subject, email.Body)
	byteMsg := []byte(msg)

	err := smtp.SendMail(s.hostName+":587", s.auth, s.from, []string{email.To}, byteMsg)
	return err

}
