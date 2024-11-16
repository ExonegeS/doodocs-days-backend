package service

import (
	"fmt"
	"net/mail"
	"os"
	"strings"

	"github.com/exoneges/doodocs-days-backend/internal/config"
	"github.com/exoneges/doodocs-days-backend/models"
	gomail "gopkg.in/mail.v2"
)

var (
	smtpHost = "smtp.gmail.com"
	smtpPort = 587
)

func SendEmailWithAttachment(file models.FileWithMeta, recipientEmails string) error {
	err := config.UpdateENV()
	if err != nil {
		return err
	}
	data, err := AnalyzeMailFile(file)
	if err != nil {
		return err
	}
	if len(data) < 1 {
		return config.ErrEmptyFile
	}

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
	message.SetHeader("From", config.ENV_MAIL_USER)
	message.SetHeader("To", emails...)
	message.SetHeader("Subject", "Here is your file!")
	message.SetBody("text/plain", "Please find the attached file.")

	// Create temporary file to send
	tempFile, err := os.CreateTemp(config.DIR, file.Filename+".*.temp")
	if err != nil {
		return err
	}

	// Write the desired content to the file

	if _, err := tempFile.Write(data); err != nil {
		tempFile.Close()
		os.Remove(tempFile.Name())
		return err
	}

	tempFile.Close()
	defer os.Remove(tempFile.Name())
	// Attach the file
	message.Attach(tempFile.Name())
	// Set up the SMTP dialer
	// port, _ := strconv.Atoi(smtpPort) // Ensure smtpPort is converted to an integer
	port := smtpPort
	dialer := gomail.NewDialer(smtpHost, port, config.ENV_MAIL_USER, config.ENV_MAIL_PASS)

	// Send the email
	if err := dialer.DialAndSend(message); err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	return nil
}
