package controllers

import (
	"cliphub/services"
	"encoding/json"
	"net/http"
)

type ContactInput struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Company   string `json:"company"`
	Message   string `json:"message"`
}

func HandleContact(w http.ResponseWriter, r *http.Request) {
	var contactInput ContactInput
	err := json.NewDecoder(r.Body).Decode(&contactInput)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid input"))
		return
	}

	if contactInput.FirstName == "" || contactInput.LastName == "" || contactInput.Email == "" || contactInput.Message == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Missing required fields"))
		return
	}
	if len(contactInput.Message) > 1000 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Message exceeds 1000 characters"))
		return
	}

	services.SendContactEmail(contactInput.FirstName, contactInput.LastName, contactInput.Email, contactInput.Company, contactInput.Message)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Message received successfully"))
}
