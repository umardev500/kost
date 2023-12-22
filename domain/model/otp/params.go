package otp

import "time"

type OtpParams struct {
	Value    OtpValue      `json:"value"`
	Dur      time.Duration `json:"-"`
	Email    string        `json:"email" validate:"required,email"`
	Location string        `json:"-"`
	Subject  string        `json:"subject" validate:"required"`
}

type OtpValue struct {
	Name string `json:"name" validate:"required" example:"Umar"`
	Otp  string `json:"-"`
}
