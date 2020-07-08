package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"mail-notifier/internal"
	"mail-notifier/pkg/mail"
)

func main() {
	// Check env. var.s
	envs := []string{mail.EnvMailFrom, mail.EnvMailSubject, mail.EnvMailContent, mail.EnvMailServer}
	if err := internal.CheckEnv(envs); err != nil {
		log.Println(err)
		os.Exit(0)
	}

	// Parse env. var.s
	from := os.Getenv(mail.EnvMailFrom)
	subject := os.Getenv(mail.EnvMailSubject)
	content := os.Getenv(mail.EnvMailContent)

	// Read list of 'to' from Env.Var.s
	var to []string
	toMap := internal.ParseUserEnv()
	for _, v := range toMap {
		to = append(to, v)
	}

	if len(to) == 0 {
		log.Println("no appr-<userid> environmental variables are found")
		os.Exit(0)
	}

	// Send mail
	reqMsg := &mail.SendRequest{
		From:    from,
		To:      to,
		Subject: subject,
		Content: content,
	}

	msgString, err := json.Marshal(reqMsg)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	req, err := http.NewRequest(mail.ServerMethod, os.Getenv(mail.EnvMailServer), bytes.NewBuffer(msgString))
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

	log.Printf("Mail sent (%+v)\n", respMsg)
}
