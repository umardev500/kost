package utils

import (
	"math/rand"
	"strconv"
)

func GenerateOTP() string {
	res := rand.Intn(90000) + 10000
	return strconv.Itoa(res)
}
