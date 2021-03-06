package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"mail-notifier/internal"
	"mail-notifier/pkg/mail"
)

func main() {
	// Env. var. check
	envs := []string{mail.EnvSmtpServer, mail.EnvSmtpUser, mail.EnvSmtpPassword}
	if err := internal.CheckEnv(envs); err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()

	log.Printf("Handler set to %s (%s)\n", mail.ServerPath, mail.ServerMethod)
	router.HandleFunc(mail.ServerPath, mailHandler).Methods(mail.ServerMethod)

	addr := fmt.Sprintf(":%d", mail.ServerPort)
	log.Printf("Server is running on %s\n", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatal(err)
	}
}
