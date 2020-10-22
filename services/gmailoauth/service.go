package gmailoauth

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/omarghader/emailer/services"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

// oauthGmailService for gmail
type oauthGmailService struct {
	GmailService *gmail.Service
}

// NewOAuthGmailService Gmail service
func NewOAuthGmailService() services.Mailer {
	config := oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Endpoint:     google.Endpoint,
		RedirectURL:  "http://localhost",
	}
	if len(os.Getenv("GOOGLE_ACCESS_TOKEN")) <= 0 {
		panic("Access token is empty")
	}

	token := oauth2.Token{
		AccessToken:  os.Getenv("GOOGLE_ACCESS_TOKEN"),
		RefreshToken: os.Getenv("GOOGLE_REFRESH_TOKEN"),
		TokenType:    "Bearer",
		Expiry:       time.Now(),
	}

	var tokenSource = config.TokenSource(context.Background(), &token)

	srv, err := gmail.NewService(context.Background(), option.WithTokenSource(tokenSource))
	if err != nil {
		log.Printf("Unable to retrieve Gmail client: %v", err)
	}

	if srv != nil {
		fmt.Println("Email service is initialized")
	}
	return &oauthGmailService{
		GmailService: srv,
	}
}

// SendEmailOAUTH2 Send email with oauth2.0
func (srv oauthGmailService) SendMail(to, emailSubject, emailBody string) error {
	var message gmail.Message

	emailTo := "To: " + to + "\r\n"
	subject := "Subject: " + emailSubject + "\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	msg := []byte(emailTo + subject + mime + "\n" + emailBody)

	message.Raw = base64.URLEncoding.EncodeToString(msg)

	// Send the message
	_, err := srv.GmailService.Users.Messages.Send("me", &message).Do()
	if err != nil {
		log.Printf("Error : %+v\n", err)
		return err
	}
	return nil
}
