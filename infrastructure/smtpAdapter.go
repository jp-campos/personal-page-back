package infrastructure

import (
	"fmt"
	"net/smtp"
	"personal-page-back/domain"
)

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

	tempMsg := []byte("To:motitaromero123@hotmail.com\r\n" +
	"Subject:Información confidencial\r\n" +
	"\r\n" +
	"Test message\r\n")
	err := smtp.SendMail(s.hostName+":587", s.auth, s.from, []string{"motitaromero123@hotmail.com"}, tempMsg)
	err = smtp.SendMail(s.hostName+":587", s.auth, s.from, []string{"champyjp99@hotmail.com"}, tempMsg)

	if err != nil {
		fmt.Println(err)
	}

}