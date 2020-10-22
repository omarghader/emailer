package main

import (
	"bytes"
	"flag"
	"html/template"
	"log"
	"strings"

	"github.com/omarghader/emailer/services/gmailoauth"
)

var (
	to = flag.String("to", "", "The recipient email address")
)

// MailModel is the struct to fill the template
type MailModel struct {
	To      string
	Subject string
	Message string
}

func main() {
	flag.Parse()

	tmpl, err := template.ParseFiles("./templates/basic.html")
	if err != nil {
		log.Fatalf("Error parsing the template: %s\n", err.Error())
	}

	if len(strings.TrimSpace(*to)) == 0 {
		log.Fatalf("to argument is madnatory !\nUsage : go run main.go -to=recipient_email_address@domain.com\n")
	}

	mailModel := MailModel{
		To:      *to,
		Subject: "My own subject",
		Message: `Hi, thank you for submitting an email using emailer.
This email is generated automatically.

Best regards.`,
	}

	// Creating a HTML template
	var body bytes.Buffer
	tmpl.Execute(&body, mailModel)

	gmailService := gmailoauth.NewOAuthGmailService()
	gmailService.SendMail(mailModel.To, mailModel.Subject, body.String())
	log.Println("Email was sent successfully")
}
