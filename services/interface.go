package services

// Mailer interface for sending mails
type Mailer interface {
	SendMail(to, emailSubject, emailBody string) error
}
