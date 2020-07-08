package main

import (
	"encoding/json"
	"fmt"
	"log"
	"mail-notifier/pkg/mail"
	"net/http"
)

func mailHandler(w http.ResponseWriter, r *http.Request) {
	req := &mail.SendRequest{}
	resp := &mail.SendResponse{}

	decoder := json.NewDecoder(r.Body)
	encoder := json.NewEncoder(w)

	// Set header
	w.Header().Set("Content-Type", "application/json")

	// Decode request body
	if err := decoder.Decode(req); err != nil {
		log.Println(err)

		resp.Sent = false
		resp.Message = fmt.Sprintf("cannot decode body, error: %s", err.Error())

		w.WriteHeader(http.StatusBadRequest)
		if err := encoder.Encode(resp); err != nil {
			log.Println(err)
		}
		return
	}

	// Send mail
	if err := mail.Send(req.From, req.To, req.Subject, req.Content); err != nil {
		log.Println(err)

		resp.Sent = false
		resp.Message = fmt.Sprintf("cannot send mail, error: %s", err.Error())

		w.WriteHeader(http.StatusBadRequest)
		if err := encoder.Encode(resp); err != nil {
			log.Println(err)
		}
		return
	}

	log.Println("Mail sent")

	// Reply
	resp.Sent = true
	resp.Message = ""
	if err := encoder.Encode(resp); err != nil {
		log.Println(err)
	}
}
