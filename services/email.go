package services

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/smtp"
	"os"
	"text/template"
)

type Constraint interface {
	Reservation | Appointment
	GetEmail() string
}

func sendMail[T Constraint](data T) error {
	type EmailBody struct {
		Data T
		Logo string
		Hero string
	}

	smtpServer := "smtp.gmail.com"
	smtpPort := "587"

	from := " reservations@cozylounge.in"
	password := "hpfa isfh obuz tooh"

	to := []string{
		"info@cozylounge.in",
		data.GetEmail(),
	}

	// mail content
	subject := "Cozy lounge reservation"
	templateData := EmailBody{
		Data: data,
		Logo: "logo",
		Hero: "hero",
	}

	t, err := template.ParseFiles("./assets/template.html")
	if err != nil {
		fmt.Println("error trying to parse email template", err)
		return err
	}

	var body bytes.Buffer
	if err := t.Execute(&body, templateData); err != nil {
		fmt.Println("error trying to execute email template", err)
		return err
	}

	logoPath := "./assets/logo.png"
	heroPath := "./assets/hero.png"
	logoAttachment, err := os.ReadFile(logoPath)
	if err != nil {
		fmt.Println("error reading logo image file", err)
		return err
	}

	heroAttachment, err := os.ReadFile(heroPath)
	if err != nil {
		fmt.Println("error reading hero image file", err)
		return err
	}

	auth := smtp.PlainAuth("", from, password, smtpServer)

	mime := "MIME-version: 1.0;\nContent-Type: multipart/related; boundary=\"MIMEBOUNDARY\"\n\n"
	mime += "--MIMEBOUNDARY\n"
	mime += "Content-Type: text/html; charset=\"UTF-8\"\n\n"
	mime += body.String() + "\n"
	mime += "--MIMEBOUNDARY\n"
	mime += "Content-Type: image/png\n"
	mime += "Content-Transfer-Encoding: base64\n"
	mime += "Content-ID: <logo>\n\n"
	mime += base64.StdEncoding.EncodeToString(logoAttachment) + "\n"

	mime += "--MIMEBOUNDARY\n"
	mime += "Content-Type: image/png\n"
	mime += "Content-Transfer-Encoding: base64\n"
	mime += "Content-ID: <hero>\n\n"
	mime += base64.StdEncoding.EncodeToString(heroAttachment) + "\n"
	mime += "--MIMEBOUNDARY--"

	toHeader := "To: " + to[0] + ", " + to[1] + "\r\n"
	subjectHeader := "Subject: " + subject + "\r\n"
	headers := toHeader + subjectHeader

	msg := []byte(headers + mime)

	err = smtp.SendMail(smtpServer+":"+smtpPort, auth, from, to, msg)
	if err != nil {
		fmt.Println("error sending mail", err)
		return err
	}

	return nil
}
