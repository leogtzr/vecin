package email

type EmailSender interface {
	Send(user, token string) error
}
