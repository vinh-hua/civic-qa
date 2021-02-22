package client

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/google/uuid"
	"github.com/vivian-hua/civic-qa/services/logAggregator/internal/model"
)

const (
	// VersionBase - URL path version base
	VersionBase = "/v0"
	// LogPath - URL path for logging
	LogPath = "/log"
	// QueryPath - URL path for queries
	QueryPath = "/query"
)

// LogClient is a client to log to and query the logAggregator
type LogClient struct {
	address string
}

// NewLogClient returns a new LogClient after validating the address
func NewLogClient(address string) (*LogClient, error) {
	_, err := url.ParseRequestURI(address)
	if err != nil {
		return nil, err
	}
	return &LogClient{address: address}, nil
}

// Log logs an event using the current UnixTime
func (client *LogClient) Log(correlationID string, service string, status int, notes string) {
	go client.logThread(correlationID, service, status, notes)
}

func (client *LogClient) logThread(correlationID string, service string, status int, notes string) {
	parsedUUID, err := uuid.Parse(correlationID)
	if err != nil {
		log.Printf("Invalid UUID: %v", correlationID)
		return
	}

	entry := model.LogEntry{
		CorrelationID: parsedUUID,
		TimeUnix:      time.Now().Unix(),
		Service:       service,
		StatusCode:    status,
		Notes:         notes,
	}

	data, err := json.Marshal(entry)
	if err != nil {
		log.Printf("Error logging: %v", err)
		return
	}

	httpClient := http.Client{}
	resp, err := httpClient.Post(
		client.address+VersionBase+LogPath,
		"application/json",
		bytes.NewBuffer(data),
	)
	if err != nil {
		log.Printf("Error logging: %v", err)
		return
	}

	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Error logging: %v", err)
			return
		}
		log.Printf("Error logging: %v", string(bodyBytes))
		return
	}
}
