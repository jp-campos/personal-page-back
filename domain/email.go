package domain

type Email struct {
	From    string
	To      string
	Subject string
	Body    string
}

type EmailGateWay interface {
	SendEmail(email Email)
}
