package graph

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"gopkg.in/gomail.v2"
)

// config
const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 587
const CONFIG_SENDER_NAME = "Zahrah Ayu Afifah <zahrah.contact@gmail.com>"
const CONFIG_AUTH_EMAIL = "zahrah.contact@gmail.com"

func generateCSVData(data [][]string) (string, error) {
	filename := fmt.Sprintf("truck_%d.csv", time.Now().Unix())
	f, err := os.Create(filename)
	defer f.Close()

	if err != nil {

		return "", errors.New(fmt.Sprintf("failed to open file %v", err))
	}

	w := csv.NewWriter(f)
	err = w.WriteAll(data) // calls Flush internally

	if err != nil {
		return "", err
	}

	return filename, nil
}

func sendEmail(filename string, to string) error {
	// Sender data.
	CONFIG_AUTH_PASSWORD := os.Getenv("EMAILPASSWORD")

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", to)
	mailer.SetHeader("Subject", "KA - Truck Email")
	mailer.SetBody("text/html", "Hello, <b>have a nice day</b>")
	mailer.Attach(filename)

	dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD,
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		return err
	}

	log.Println("Mail sent!")
	return nil
}
