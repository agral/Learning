package main

import (
	"encoding/json"
	"net/http"
)

type stripePayload struct {
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}

type jsonResponse struct {
	IsOk    bool   `json:"isOk"`
	Message string `json:"message"`
	Content string `json:"content"`
	Id      int    `json:"id"`
}

func (app *application) GetPaymentIntent(w http.ResponseWriter, r *http.Request) {
	j := jsonResponse{
		IsOk: true,
	}

	out, err := json.MarshalIndent(j, "", "    ")
	if err != nil {
		app.errorLog.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}
