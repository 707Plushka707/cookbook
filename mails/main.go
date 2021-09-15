package main

import (
	"fmt"

	mail "gopkg.in/mail.v2"
)

func main() {
	msg := mail.NewMessage()
	msg.SetHeader("From", "Ahmed@gmail.com")
	msg.SetHeader("To", "Ahmedin@gmail.com")
	msg.SetHeader("Subject", "test golang email")
	msg.SetBody("text/html", "<h2>this is an fucking spam email</h2>")

	d := mail.NewDialer("smtp.gmail.com", 587, "ahmed@email.com", "mySerperSecretPass") // .env
	d.StartTLSPolicy = mail.MandatoryStartTLS

	err := d.DialAndSend(msg)
	if err != nil {
		fmt.Println(err)
	}
}
