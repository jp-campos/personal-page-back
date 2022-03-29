package domain

import "context"

var emailRepo EmailGateWay

func InitEmailRepository(repo EmailGateWay) {
	emailRepo = repo
}

func SendMail(ctx context.Context, email Email) {
	emailRepo.SendEmail(email)
}
