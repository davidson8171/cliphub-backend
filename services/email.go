package services

import (
	"fmt"
	"net/smtp"
	"os"
)

var emailChannel chan EmailChannel

type EmailChannel struct {
	firstName string
	lastName  string
	email     string
	company   string
	message   string
}

func EmailContactService() {
	from := os.Getenv("EMAIL")
	to := os.Getenv("EMAIL_CONTACT")
	password := os.Getenv("EMAIL_PASSWORD")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	auth := smtp.PlainAuth("", from, password, smtpHost)

	emailChannel = make(chan EmailChannel)

	for emailData := range emailChannel {
		companyString := " "
		if emailData.company != "" {
			companyString = " | Unternehmen: " + emailData.company + " "
		}

		messageByte := []byte("Subject: [KONTAKT] von " + emailData.firstName + " " + emailData.lastName + companyString + "| Email: " + emailData.email + "\r\n" +
			"Content-Type: text/plain; charset=\"UTF-8\"\r\n" +
			"\r\n" + emailData.message)

		err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, messageByte)
		if err != nil {
			fmt.Println("[EMAILCONTACTSERVICE] Error:", err)
		}
	}

	close(emailChannel)
}

func SendContactEmail(firstName string, lastName string, email string, company string, message string) {
	emailChannel <- EmailChannel{
		firstName: firstName,
		lastName:  lastName,
		email:     email,
		company:   company,
		message:   message,
	}
}
