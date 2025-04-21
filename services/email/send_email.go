package email

import (
	"bytes"
	"encoding/base64"
	"html/template"
	"net/smtp"
	"os"
	"strings"
)

type Attachment struct {
	Path        string
	ContentType string
	ContentID   string
	Inline      bool
}

// ISendEmail is implemented by services that want to send email to user
type ISendEmail interface {
	GetToSend() []string
	GetSubject() string
	GetTemplateDir() string
	GetTemplateData() any
	GetAttachments() []Attachment
}

func (es *emailService) SendEmail(data ISendEmail) error {
	// get template
	tmpl, err := template.ParseFiles(data.GetTemplateDir())
	if err != nil {
		return err
	}

	var body bytes.Buffer
	if err := tmpl.Execute(&body, data.GetTemplateData()); err != nil {
		return err
	}

	// Prepare attachments
	var mimeBody strings.Builder
	boundary := "MIMEBOUNDARY"

	mimeBody.WriteString("MIME-version: 1.0;\nContent-Type: multipart/related; boundary=\"" + boundary + "\"\n\n")
	mimeBody.WriteString("--" + boundary + "\n")
	mimeBody.WriteString("Content-Type: text/html; charset=\"UTF-8\"\n\n")
	mimeBody.WriteString(body.String() + "\n")

	for _, att := range data.GetAttachments() {
		fileBytes, err := os.ReadFile(att.Path)
		if err != nil {
			return err
		}

		mimeBody.WriteString("--" + boundary + "\n")
		mimeBody.WriteString("Content-Type: " + att.ContentType + "\n")
		mimeBody.WriteString("Content-Transfer-Encoding: base64\n")
		if att.Inline && att.ContentID != "" {
			mimeBody.WriteString("Content-ID: <" + att.ContentID + ">\n")
		}
		mimeBody.WriteString("\n")
		mimeBody.WriteString(base64.StdEncoding.EncodeToString(fileBytes) + "\n")
	}
	mimeBody.WriteString("--" + boundary + "--")

	// Compose headers
	toHeader := "To: " + strings.Join(data.GetToSend(), ", ") + "\r\n"
	subjectHeader := "Subject: " + data.GetSubject() + "\r\n"
	headers := toHeader + subjectHeader

	msg := []byte(headers + mimeBody.String())
	auth := smtp.PlainAuth("", es.from, es.password, es.smtpServer)

	err = smtp.SendMail(es.smtpServer+":"+es.smtpPort, auth, es.from, data.GetToSend(), msg)
	if err != nil {
		return err
	}

	return nil
}
