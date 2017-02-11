package mygoutils

import (
	"log"
	"net/smtp"
	"os"
)

/* Set your username and password as environment variables*/
func SendEmail(from string, subject string, body string, to []string) bool {

	username := os.Getenv("GMAIL_USERNAME")
	password := os.Getenv("GMAIL_PASSWORD")

	msg := "From: " + from + "\n" +
		"To: " + from + "\n" +
		"Subject:  " + subject + "\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", username, password, "smtp.gmail.com"),
		from, to, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return false
	}

	return true
}
