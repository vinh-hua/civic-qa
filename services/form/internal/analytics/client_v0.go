package analytics

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/team-ravl/civic-qa/services/form/internal/model"
)

const (
	v0KeyphrasePath = "/v0/key-phrase"
)

type ClientV0 struct {
	Address string
}

// NewClientV0 returns a new ClientV0
func NewClientV0(address string) (*ClientV0, error) {
	_, err := url.Parse(address)
	if err != nil {
		return nil, err
	}

	return &ClientV0{Address: address}, nil

}

// GetKeyPhrases extracts keyphrases from a formResponse
func (c *ClientV0) GetKeyPhrases(formResponse *model.FormResponse) ([]string, error) {
	// make the request body
	req := struct {
		Subject string `json:"subject"`
		Body    string `json:"body"`
	}{
		Subject: formResponse.Subject,
		Body:    formResponse.Body,
	}

	// create request bytes
	requestBytes, err := json.Marshal(&req)
	if err != nil {
		return nil, err
	}

	// perform the request
	httpClient := &http.Client{}
	resp, err := httpClient.Post(c.Address+v0KeyphrasePath, "application/json", bytes.NewReader(requestBytes))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// read the response
	var tags []string
	err = json.NewDecoder(resp.Body).Decode(&tags)
	if err != nil {
		return nil, err
	}

	return tags, nil
}
