package utils

import (

	"gopkg.in/gomail.v2"
)


func SendEmail(email string, otp string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "httplocalhost3030@gmail.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "OTP Verification")
	m.SetBody("text/plain", "Your OTP: "+otp)

	d := gomail.NewDialer("smtp.gmail.com", 587, "httplocalhost3030@gmail.com", "mbifnxkqblhzrnpo")

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
