package domain

import (
	"context"
	"os"

	"golang.org/x/sync/errgroup"
)

var emailRepo EmailGateWay


const (
	personalMailKey = "PERSONAL_MAIL"
	senderEmailKey = "SENDER_EMAIL"
)


func InitEmailRepository(repo EmailGateWay) {
	emailRepo = repo
}

func SendMail(ctx context.Context, email Email) error{

	errs, ctx := errgroup.WithContext(ctx)
	
	//Send email to self
	errs.Go(
		 func() error{
		
		email.To =  os.Getenv(personalMailKey)
		email.Subject = "Mail from personal page from: " + email.From

		return emailRepo.SendEmail(email)

	})
	//Send confirmation email to sender
	errs.Go(func() error {
	
		var confirmationEmail = Email{
			From:    os.Getenv(senderEmailKey),
			To:      email.From,
			Subject: "Thank you for contacting me!: " + email.From,
			Body: "This email is a confirmation that I have received your email!" +
				"\nThis is an automatically sent email, please contact me at " + os.Getenv(personalMailKey) + "\n\n" +
				"Kind Regards"}

		return emailRepo.SendEmail(confirmationEmail)

	})

	return errs.Wait()

}
