package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateOTP() string {
	rand.Seed(time.Now().UnixNano())
	OTP := fmt.Sprintf("%04d", rand.Intn(1000000))
	return OTP
}
