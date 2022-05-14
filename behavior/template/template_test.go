package template

import (
	"fmt"
	"testing"
)

func TestTemplate(t *testing.T) {
	t.Run("template", func(t *testing.T) {
		smsOTP := &sms{}
		o := otp{
			iOtp: smsOTP,
		}
		o.genAndSendOTP(4)

		fmt.Println("")
		emailOTP := &email{}
		o = otp{
			iOtp: emailOTP,
		}
		o.genAndSendOTP(4)
	})
}
