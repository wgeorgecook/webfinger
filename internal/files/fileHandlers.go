package files

import (
	"encoding/json"
	"log"
	"os"
)

type webfingerResponse struct {
	Subject string           `json:"subject,omitempty"`
	Aliases []string         `json:"aliases,omitempty"`
	Links   []webfingerLinks `json:"links,omitempty"`
}
type webfingerLinks struct {
	Rel      string `json:"rel,omitempty"`
	Type     string `json:"type,omitempty"`
	Href     string `json:"href,omitempty"`
	Template string `json:"template,omitempty"`
}

func ReadPayload() (*webfingerResponse, error) {
	log.Println("attempting to open payload")
	payload, err := os.ReadFile("./data/payload.json")
	if err != nil {
		log.Println(err)
		return nil, err

	}
	log.Println("open successful")

	var resp *webfingerResponse
	if err := json.Unmarshal(payload, &resp); err != nil {
		log.Println(err)
		return nil, err
	}

	return resp, nil
}
