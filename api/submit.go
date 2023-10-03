package handler

import (
	"net/http"
	"timesheet-bot/scraper"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	url := "https://example.com/login"
	inputs := map[string]string{
		"#username": "YourUsername",
		"#password": "YourPassword",
	}
	buttonSelector := "#login"

	result, err := scraper.FillAndClick(url, inputs, buttonSelector)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(result))
}