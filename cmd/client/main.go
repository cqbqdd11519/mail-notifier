package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	"mail-notifier/pkg/mail"
)

const (
	ServiceName = "mail-sender"
)

func main() {
	// Parse env. var.s
	from := os.Getenv("MAIL_FROM")
	if from == "" {
		log.Fatal("MAIL_FROM should be set")
	}

	toRaw := os.Getenv("MAIL_TO")
	if toRaw == "" {
		log.Fatal("MAIL_TO should be set")
	}
	to := strings.Split(toRaw, ",")

	subject := os.Getenv("MAIL_SUBJECT")
	if subject == "" {
		log.Fatal("MAIL_SUBJECT should be set")
	}

	content := os.Getenv("MAIL_CONTENT")
	if content == "" {
		log.Fatal("MAIL_CONTENT should be set")
	}

	// Send mail
	reqMsg := &mail.SendRequest{
		From: from,
		To: to,
		Subject: subject,
		Content: content,
	}

	msgString, err := json.Marshal(reqMsg)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	req, err := http.NewRequest(mail.ServerMethod, ServiceName, bytes.NewBuffer(msgString))
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	respMsg := &mail.SendResponse{}
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(respMsg); err != nil {
		log.Fatal(err)
	}

	if !respMsg.Sent {
		log.Fatal(respMsg.Message)
	}

	if err := resp.Body.Close(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Mail sent (%+v)\n", req)
}
