package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type User struct {
	Id           string    `json:"id,omitempty"`
	Username     string    `json:"username,omitempty"`
	Email        string    `json:"email,omitempty"`
	Password     string    `json:"password,omitempty"`
	Name         string    `json:"name,omitempty"`
	OTP          string    `json:"otp,omitempty"`
	OTPTimestamp time.Time `json:"otpTimestamp,omitempty"`
	EnteredOTP   string    `json:"enteredOtp,omitempty"`
	created_at   time.Time `json:"created_at,omitempty"`
}

type CodeName struct {
	Username string `json:"username,omitempty"`
	Codename string `json:"codename,omitempty"`
}

type signedDetails struct {
	Email string
	uid   string
	jwt.StandardClaims
}
