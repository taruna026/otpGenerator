package service

import (
	"math/rand"
	"strconv"
)

func GenerateOtp() string {
	otpLen := 5

	var otp string

	for i := 0; i < otpLen; i++ {
		num := rand.Int() % 10
		otp += strconv.Itoa(num)
	}

	return otp
}
