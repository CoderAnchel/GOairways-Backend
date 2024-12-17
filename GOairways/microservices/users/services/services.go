package services

import (
	"bytes"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"github.com/resend/resend-go/v2"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"math/rand"
	"strconv"
	"time"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateRandomPIN() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(900000) + 100000
}

func renderHTMLTemplate(PIN int) (string, error) {
	tmpl, err := template.ParseFiles("./template.html")
	if err != nil {
		return "", err
	}

	var buffer bytes.Buffer
	err = tmpl.Execute(&buffer, PIN)
	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}

func SendMail(PIN int) {
	tmpl := `<!DOCTYPE html>
			<html>
<head>
    <title>Welcome to GOairways</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            line-height: 1.6;
            background-color: #f9f9f9;
            margin: 0;
            padding: 0;
        }
        .email-container {
            max-width: 600px;
            margin: 20px auto;
            background: #ffffff;
            padding: 20px;
            border: 1px solid #dddddd;
            border-radius: 5px;
        }
        .header {
            text-align: center;
            padding-bottom: 20px;
        }
        .header h1 {
            color: #333333;
        }
        .content {
            color: #555555;
        }
        .footer {
            text-align: center;
            margin-top: 20px;
            font-size: 12px;
            color: #999999;
        }
        .button {
            display: inline-block;
            padding: 10px 20px;
            margin-top: 20px;
            background-color: #007bff;
            color: white;
            text-decoration: none;
            border-radius: 5px;
            width: 90%;
        }
        .button:hover {
            background-color: #0056b3;
        }
    </style>
</head>
<body>
<div class="email-container">
    <div class="header">
        <h1>Welcome to GOairways</h1>
    </div>
    <div class="content">
        <p>Thank you for choosing GOairways! Your verification PIN is:</p>
        <h2 style="text-align: center; color: #e03131;">{{.PIN}}</h2>
        <p>To complete your log-in, please enter this PIN in the verification field.</p>
    </div>
    <div class="footer">
        <p>If you did not request this email, please ignore it.</p>
        <p>&copy; 2024 GOairways. All rights reserved.</p>
    </div>
</div>
</body>
</html>`

	data := struct {
		PIN string
	}{
		PIN: strconv.Itoa(PIN),
	}

	t, err := template.New("email").Parse(tmpl)

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		log.Fatalf("Error executing template: %v", err)
	}

	apiKey := "re_5z9KyNeq_5xkskNYQhwyXWyDTyZvKJPyn"

	client := resend.NewClient(apiKey)

	params := &resend.SendEmailRequest{
		From:    "Acme <onboarding@resend.dev>",
		To:      []string{"alex21fd@gmail.com"},
		Html:    buf.String(),
		Subject: "Welcome to GOairways",
		Cc:      []string{"cc@example.com"},
		Bcc:     []string{"bcc@example.com"},
		ReplyTo: "replyto@example.com",
	}

	sent, err := client.Emails.Send(params)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(sent.Id)
}
