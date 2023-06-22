package models

import "time"

type User struct {
	Id           string    `json:"id,omitempty"`
	Username     string    `json:"username,omitempty"`
	Email        string    `json:"email,omitempty"`
	Password     string    `json:"password,omitempty"`
	Name         string    `json:"name,omitempty"`
	OTP          string    `json:"otp,omitempty"`
	OTPTimestamp time.Time `json:"otpTimestamp,omitempty"`
	EnteredOTP   string    `json:"enteredOtp,omitempty"`
}
