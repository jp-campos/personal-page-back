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
	fmt.Println(hostName, from, password)
	auth := smtp.PlainAuth("", from, password, hostName)
	return SmtpAdapter{auth: auth, hostName: hostName, from: from, password: password}
}

func (s SmtpAdapter) SendEmail(email domain.Email) {
	msg := fmt.Sprintf("To:motitaromero123@hotmail.com\r\n"+
		"Subject:Información confidencial\r\n"+
		"\r\n"+
		"%s \r\n", email.Body)
	byteMsg := []byte(msg)

	wg.Add(2)
	go func() {
		err := smtp.SendMail(s.hostName+":587", s.auth, s.from, []string{"jp.campos99@hotmail.com"}, byteMsg)
		if err != nil {
			fmt.Println(err)
		}
		wg.Done()
	}()
	go func() {
		smtp.SendMail(s.hostName+":587", s.auth, s.from, []string{"champyjp99@hotmail.com"}, byteMsg)
		wg.Done()
	}()
	wg.Wait()
}
