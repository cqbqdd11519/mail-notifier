package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"mail-notifier/internal"
	"mail-notifier/pkg/mail"
)

const (
	ConfigMapPath string = "/tmp/config/users"
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
	if internal.FileExists(ConfigMapPath) {
		file, err := os.Open(ConfigMapPath)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		// Parse line-separated
		for scanner.Scan() {
			text := scanner.Text()

			// Parse comma-separated
			userList := strings.Split(text, ",")
			for i := range userList {
				userList[i] = strings.TrimSpace(userList[i])

				pair := strings.Split(userList[i], "=")
				if len(pair) != 2 {
					log.Fatal(fmt.Errorf("%s is not in <userId>=<email> form", userList[i]))
				}
				to = append(to, pair[1])
			}
		}
	}

	receiverStr := os.Getenv(mail.EnvMailTo)
	if receiverStr != "" {
		receivers := strings.Split(receiverStr, ",")
		for i := range receivers {
			receivers[i] = strings.TrimSpace(receivers[i])
		}
		to = append(to, receivers...)
	}

	if len(to) == 0 {
		log.Println("no mail receiver is found")
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
