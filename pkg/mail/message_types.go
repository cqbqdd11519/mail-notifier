package mail

type SendRequest struct {
	From    string   `json:"from"`
	To      []string `json:"to"`
	Subject string   `json:"subject"`
	Content string   `json:"content"`
}

type SendResponse struct {
	Sent    bool   `json:"sent"`
	Message string `json:"message,omitempty"`
}
