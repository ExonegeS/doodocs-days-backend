package service

import (
	"fmt"
	"net/mail"
	"os"
	"strings"

	"github.com/exoneges/doodocs-days-backend/models"
	gomail "gopkg.in/mail.v2"
)

var (
	smtpHost     = "smtp.gmail.com"
	smtpPort     = 587
	smtpUser     = os.Getenv("DOODOCS_DAYS2_BACKEND_MAIL_USERNAME")
	smtpPassword = os.Getenv("DOODOCS_DAYS2_BACKEND_MAIL_PASSWORD")
)

func SendEmailWithAttachment(file models.FileWithMeta, recipientEmails string) error {
	// Parse the email list
	emails := strings.Split(recipientEmails, ",")
	for _, email := range emails {
		_, err := mail.ParseAddress(email)
		if err != nil {
			return fmt.Errorf("invalid email: %s", email)
		}
	}

	// Create a new email message
	message := gomail.NewMessage()
	message.SetHeader("From", smtpUser)
	message.SetHeader("To", emails...)
	message.SetHeader("Subject", "Here is your file!")
	message.SetBody("text/plain", "Please find the attached file.")

	// Attach the file
	message.AttachReader(file.Filename, file.File)

	// Set up the SMTP dialer
	// port, _ := strconv.Atoi(smtpPort) // Ensure smtpPort is converted to an integer
	port := smtpPort
	dialer := gomail.NewDialer(smtpHost, port, smtpUser, smtpPassword)

	// Send the email
	if err := dialer.DialAndSend(message); err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	return nil
}