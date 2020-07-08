package mail

import (
	"fmt"
	"net/smtp"
	"os"
	"strconv"
	"strings"
)

func Send(from string, to []string, subject string, content string) error {
	server, err := serverInfo()
	if err != nil {
		return err
	}

	auth, err := auth(server)
	if err != nil {
		return err
	}

	fromBody := "From: <" + from + ">\r\n"

	toBody := "To: "
	for i, t := range to {
		if i != 0 {
			toBody += ", "
		}
		toBody += "<" + t + ">"
	}
	toBody += "\r\n"

	subjectBody := "Subject: " + subject + "\r\n"

	msg := []byte(fromBody + toBody + subjectBody + "\r\n" + content)

	return smtp.SendMail(server.host+":"+strconv.Itoa(server.port), auth, from, to, msg)
}

type SmtpInfo struct {
	host     string
	port     int
	user     string
	password string
}

func serverInfo() (*SmtpInfo, error) {
	info := &SmtpInfo{}

	// Server URL
	serverUrl := os.Getenv(EnvSmtpServer)

	// Split into hostname, port
	hosts := strings.Split(serverUrl, ":")
	if len(hosts) != 2 {
		return nil, fmt.Errorf("%s(%s) invalid, it should be [IP]:[PORT] form", EnvSmtpServer, serverUrl)
	}

	info.host = hosts[0]
	port, err := strconv.Atoi(hosts[1])
	if err != nil {
		return nil, err
	}
	info.port = port

	// Username
	info.user = os.Getenv(EnvSmtpUser)

	// Password
	info.password = os.Getenv(EnvSmtpPassword)

	return info, nil
}

func auth(server *SmtpInfo) (auth smtp.Auth, err error) {
	return smtp.PlainAuth("", server.user, server.password, server.host), nil
}
