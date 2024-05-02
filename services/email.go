package services

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/smtp"
	"os"
	"text/template"
)

type EmailBody struct {
	Contact
	Logo string
}

func sendMail(data Contact) error {
	smtpServer := "smtp.gmail.com"
	smtpPort := "587"

	from := "rohan@mennr.tech"
	password := "tkhv bcse ubwn wmgj"

	to := []string{
		"vermarohan031@gmail.com",
	}

	// mail content
	subject := "New form submission"
	templateData := EmailBody{
		Contact: data,
		Logo:    "logo",
	}
	fmt.Println(templateData.Contact.Name)

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

	imagePath := "./assets/logo.png"
	attachement, err := os.ReadFile(imagePath)
	if err != nil {
		fmt.Println("error reading image file", err)
		return err
	}

	// pdfPath := "trial.pdf"
	// pdfAttachment, err := os.ReadFile(pdfPath)
	// if err != nil {
	// 	log.Fatal("error reading pdf file", err)
	// }

	auth := smtp.PlainAuth("", from, password, smtpServer)

	mime := "MIME-version: 1.0;\nContent-Type: multipart/related; boundary=\"MIMEBOUNDARY\"\n\n"
	mime += "--MIMEBOUNDARY\n"
	mime += "Content-Type: text/html; charset=\"UTF-8\"\n\n"
	mime += body.String() + "\n"
	mime += "--MIMEBOUNDARY\n"
	mime += "Content-Type: image/png\n"
	mime += "Content-Transfer-Encoding: base64\n"
	mime += "Content-ID: <logo>\n\n"
	mime += base64.StdEncoding.EncodeToString(attachement) + "\n"
	mime += "--MIMEBOUNDARY--"

	// with attachemnt
	// mime := "MIME-version: 1.0; \nContent-Type: multipart/mixed; boundary=\"MIMEBOUNDARY\"\n\n"
	// mime += "--MIMEBOUNDARY\n"
	// mime += "Content-Type: multipart/related; boundary=\"RELEATEDBOUNDARY\"\n\n"
	// mime += "--RELEATEDBOUNDARY\n"
	// mime += "Content-Type: text/html; charset=\"UTF-8\"\n\n"
	// mime += body.String() + "\n"
	// mime += "--RELEATEDBOUNDARY\n"
	// mime += "Content-Type: image/png\n"
	// mime += "Content-Transfer-Encoding: base64\n"
	// mime += "Content-ID: <logo>\n\n"
	// mime += base64.StdEncoding.EncodeToString(attachement) + "\n"
	// mime += "--RELEATEDBOUNDARY--\n"
	// mime += "--MIMEBOUNDARY\n"
	// mime += "Content-Type: application/pdf; name=\"rohan's wpr.pdf\"\n"
	// mime += "Content-Transfer-Encoding: base64\n"
	// // mime += "Content-Disposition: attachment; filename=\"rohan'ss wpr.pdf\"\n\n" // this is recommended way
	// mime += "Content-Disposition: attachment;\n\n"
	// mime += base64.StdEncoding.EncodeToString(pdfAttachment) + "\n"
	// mime += "--MIMEBOUNDARY--\n"

	toHeader := "To: " + to[0] + "\r\n"
	subjectHeader := "Subject: " + subject + "\r\n"
	headers := toHeader + subjectHeader

	msg := []byte(headers + mime)

	err = smtp.SendMail(smtpServer+":"+smtpPort, auth, from, to, msg)
	if err != nil {
		fmt.Println("error sending mail", err)
		return err
	}

	fmt.Println("mail sent!")
	return nil
}
