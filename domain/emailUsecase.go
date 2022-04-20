package domain

import (
	"context"
	"os"
	"sync"
)

var emailRepo EmailGateWay

var wg sync.WaitGroup

const PERSONAL_MAIL = "PERSONAL_MAIL"
const SENDER_EMAIL = "SENDER_EMAIL"

func InitEmailRepository(repo EmailGateWay) {
	emailRepo = repo
}

func SendMail(ctx context.Context, email Email) {

	wg.Add(2)
	//Send email to self
	go func() {

		email.To = os.Getenv(PERSONAL_MAIL)
		email.Subject = "Mail from personal page from: " + email.From

		emailRepo.SendEmail(email)
		wg.Done()
	}()
	//Send confirmation email to sender
	go func() {

		var confirmationEmail = Email{
			From:    os.Getenv(SENDER_EMAIL),
			To:      email.From,
			Subject: "Thank you for contacting me!: " + email.From,
			Body: "This email is a confirmation that I have received your email!" +
				"\nThis is an automatically sent email, please contact me at " + os.Getenv(PERSONAL_MAIL) + "\n\n" +
				"Kind Regards"}

		emailRepo.SendEmail(confirmationEmail)

		wg.Done()
	}()

	wg.Wait()

}
