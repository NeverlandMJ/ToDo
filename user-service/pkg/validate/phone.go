package validate

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/verify/v2"
)

var ERR_CODE_INCORRECT = fmt.Errorf("entered code is incorrect")

type Otp struct {
	TWILIO_ACCOUNT_SID string
	TWILIO_AUTH_TOKEN  string
	VERIFY_SERVICE_SID string
	client             *twilio.RestClient
}

func NewOtp() Otp {
	return Otp{
		TWILIO_ACCOUNT_SID: os.Getenv("TWILIO_ACCOUNT_SID"),
		TWILIO_AUTH_TOKEN:  os.Getenv("TWILIO_AUTH_TOKEN"),
		VERIFY_SERVICE_SID: os.Getenv("VERIFY_SERVICE_SID"),
		client: twilio.NewRestClientWithParams(twilio.ClientParams{
			Username: os.Getenv("TWILIO_ACCOUNT_SID"),
			Password: os.Getenv("TWILIO_AUTH_TOKEN"),
		}),
	}
}

func (o *Otp) SendOtp(to string) (string, error) {
	params := &openapi.CreateVerificationParams{}
	params.SetTo(to)
	params.SetChannel("sms")

	resp, err := o.client.VerifyV2.CreateVerification(o.VERIFY_SERVICE_SID, params)

	if err != nil {
		return "", err
	} 

	return *resp.Sid, nil
}

func (o *Otp) CheckOtp(to, code string) error {
   params := &openapi.CreateVerificationCheckParams{}
   params.SetTo(to)
   params.SetCode(code)

   resp, err := o.client.VerifyV2.CreateVerificationCheck(o.VERIFY_SERVICE_SID, params)

   if err != nil {
       fmt.Println(err.Error())
   } else if *resp.Status == "approved" {
       return nil
   } 
   
   return ERR_CODE_INCORRECT
}


