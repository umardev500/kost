package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
	"gopkg.in/mail.v2"
)

func SendMail(to []string, cc []string, subject, message string) (err error) {
	log.Debug().Msg("Sending email...")
	start := time.Now()
	host := os.Getenv("EMAIL_HOST")
	port, err := strconv.Atoi(os.Getenv("EMAIL_PORT"))
	if err != nil {
		return
	}

	sender := os.Getenv("EMAIL_SENDER")
	pass := strings.ReplaceAll(os.Getenv("EMAIL_PASS"), " ", "")

	m := mail.NewMessage()
	m.SetHeader("From", sender)
	m.SetHeader("To", to...)
	m.SetHeader("Cc", cc...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", message)

	d := mail.NewDialer(host, int(port), sender, pass)

	if err := d.DialAndSend(m); err != nil {
		msg := fmt.Sprintf("Failed to send email: %s", err.Error())
		log.Error().Msg(msg)
		return err
	}

	ended := time.Since(start)
	msg := fmt.Sprintf("Email sent in \033[32m%s\033[0m", ended)
	log.Debug().Msg(msg)

	return nil
}
