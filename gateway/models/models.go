package models

// Text - struct
type Text struct {
	Body string `json:"body"`
}

// Messages - struct
type Messages struct {
	From      string `json:"from"`
	ID        string `json:"id"`
	Text      Text   `json:"text"`
	Timestamp string `json:"timestamp"`
	Type      string `json:"type"`
}

// Profile - struct
type Profile struct {
	Name string `json:"name"`
}

// Contract - struct
type Contract struct {
	Profile Profile `json:"profile"`
	Phone   string  `json:"wa_id"`
}

// WhatsappPostBack - struct postback whatsapp
type WhatsappPostBack struct {
	Contacts    []Contract `json:"contacts"`
	MediaID     string     `json:"media_id"`
	MessageType string     `json:"message_type"`
	Messages    []Messages `json:"messages"`
}

// WhatsappMessage - struct whatsapp message receive
type WhatsappMessage struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Message string `json:"message"`
}

// ReceiveParsePostBack - handle receive postBack
func (message *WhatsappMessage) ReceiveParsePostBack(postBack *WhatsappPostBack) {
	for _, contact := range postBack.Contacts {

		message.Name = contact.Profile.Name

		if len(contact.Profile.Name) > 0 {
			message.Name = contact.Profile.Name
		}

		if len(contact.Phone) > 0 {
			message.Phone = contact.Phone
		}
	}

	for _, m := range postBack.Messages {
		if len(m.Text.Body) > 0 {
			message.Message = m.Text.Body
		}
	}
}
