package otp

import (
	"bytes"
	"context"
	"html/template"
	"os"

	"github.com/umardev500/kost/domain/model/otp"
	"github.com/umardev500/kost/utils"
)

func (ou *otpUsecase) SendOtp(ctx context.Context, params otp.OtpParams) (err error) {
	otpID := utils.GenerateOTP()
	params.Value.Otp = otpID
	err = ou.storage.Set(otpID, []byte("true"), params.Dur)
	if err != nil {
		return
	}

	templateBytes, err := os.ReadFile(params.Location)
	if err != nil {
		ou.storage.Delete(otpID) // delete otp
		return err
	}
	tmpl, err := template.New("otp").Parse(string(templateBytes))
	if err != nil {
		ou.storage.Delete(otpID) // delete otp
		return err
	}

	var resultBuffer bytes.Buffer
	if err := tmpl.Execute(&resultBuffer, params.Value); err != nil {
		ou.storage.Delete(otpID) // delete otp
		return err
	}

	msg := resultBuffer.String()
	utils.SendMail([]string{params.Email}, []string{}, params.Subject, msg)

	return nil
}
