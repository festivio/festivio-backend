package validator

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/festivio/festivio-backend/config"
)

type numverify struct {
	Valid    bool   `json:"valid"`
	LineType string `json:"line_type"`
}

func IsValidPhoneNumber(number string) (bool, error) {
	cfg := config.MustLoad()
	safePhone := url.QueryEscape(number)
	url := fmt.Sprintf("http://apilayer.net/api/validate?access_key=%s&number=%s", cfg.ExternalApi.NumverifyAccessKey, safePhone)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return false, err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return false, err
	}
	defer resp.Body.Close()
	var record numverify
	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
		return false, err
	}
	if !record.Valid && record.LineType != "mobile" {
		return false, nil
	}
	return true, nil
}
