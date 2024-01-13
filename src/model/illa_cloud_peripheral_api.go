package model

import (
	"errors"
	"fmt"
	"github.com/wneessen/go-mail"
	"log"
	"net/http"

	"github.com/go-resty/resty/v2"
)

const (
	EMAIL_DELIVER_BASEURL                = "https://email.illasoft.com/v1/"
	EMAIL_DELIVER_USAGE_SUBSCRIBE        = "subscribe"
	EMAIL_DELIVER_USAGE_VERIFICATIONCODE = "code"
	EMAIL_DELIVER_INVITE_EMAIL           = "invite"
	EMAIL_DELIVER_SHARE_APP_EMAIL        = "shareApp"
)

const ()

func SendSubscriptionEmail(email string) error {
	client := resty.New()
	resp, err := client.R().
		SetBody(map[string]string{"email": email}).
		Post(EMAIL_DELIVER_BASEURL + EMAIL_DELIVER_USAGE_SUBSCRIBE)
	if resp.StatusCode() != http.StatusOK || err != nil {
		return errors.New("failed to send subscription email")
	}
	fmt.Printf("response: %+v, err: %+v", resp, err)
	return nil
}

func SendVerificationEmail(email, code, usage string) error {
	//client := resty.New()
	//resp, err := client.R().
	//	SetBody(map[string]string{"email": email, "code": code, "usage": usage}).
	//	Post(EMAIL_DELIVER_BASEURL + EMAIL_DELIVER_USAGE_VERIFICATIONCODE)
	//if resp.StatusCode() != http.StatusOK || err != nil {
	//	return errors.New("failed to send verification code email")
	//}
	//fmt.Printf("response: %+v, err: %+v", resp, err)
	m := mail.NewMsg()
	if err := m.From("toni.sender@example.com"); err != nil {
		log.Fatalf("failed to set From address: %s", err)
	}
	if err := m.To(email); err != nil {
		log.Fatalf("failed to set To address: %s", err)
	}
	m.Subject("This is" + usage + "code:  " + code + "  !")
	m.SetBodyString(mail.TypeTextPlain, "Do you like this mail? I certainly do!")
	c, err := mail.NewClient("smtp.gmail.com", mail.WithPort(587), mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithUsername("webstarchina@gmail.com"), mail.WithPassword("segftdvgsncjcrap"))
	if err != nil {
		log.Fatalf("failed to create mail client: %s", err)
	}
	if err := c.DialAndSend(m); err != nil {
		log.Fatalf("failed to send mail: %s", err)
	}
	return nil
}
