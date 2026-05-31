package handlers

import (
	"encoding/json"
	"net/http"
)

type Invoice struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func ListInvoices(w http.ResponseWriter, _ *http.Request) {
	items := []Invoice{{ID: 1, Name: "sample"}}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(items)
}
