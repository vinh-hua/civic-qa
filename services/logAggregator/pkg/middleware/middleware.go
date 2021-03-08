package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/urfave/negroni"
)

const (
	defaultTimeout  = 5 * time.Second
	defaultHostname = "unavailable"
	logPath         = "/v0/log"
)

// Config contains settings to perform logging
type Config struct {
	// Exposed fields
	AggregatorAddress string
	ServiceName       string
	SkipSuccesses     bool
	StdoutErrors      bool
	Timeout           time.Duration

	// Unexposed fields
	hostname string
}

// newLogEntry represents a LogEntry before it is added to the database
type newLogEntry struct {
	CorrelationID string `json:"correlationID"`
	TimeUnix      int64  `json:"timeUnix"`
	HTTPMethod    string `json:"httpMethod"`
	RequestPath   string `json:"requestPath"`
	Service       string `json:"service"`
	StatusCode    int    `json:"statusCode"`
	Hostname      string `json:"hostname"`
	Notes         string `json:"notes"`
}

// NewAggregatorMiddleware returns a mux.MiddlewareFunc
// To log all requests to an instance of the logAggregator
func NewAggregatorMiddleware(config *Config) func(http.Handler) http.Handler {
	// Establish hostname
	hostname, err := os.Hostname()
	if err != nil {
		log.Printf("Error getting hostname: %v", err)
		hostname = defaultHostname
	}
	config.hostname = hostname

	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Wrap the response writer so we can get the status after
			wrapped := negroni.NewResponseWriter(w)
			// serve the request as usual
			h.ServeHTTP(wrapped, r)

			// create the log entry
			entry := &newLogEntry{
				CorrelationID: r.Header.Get("X-Correlation-ID"),
				TimeUnix:      time.Now().Unix(),
				HTTPMethod:    r.Method,
				RequestPath:   r.URL.String(),
				Service:       config.ServiceName,
				StatusCode:    wrapped.Status(),
				Hostname:      config.hostname,
			}

			// if SkipSuccesses, don't log unless this request failed
			if config.SkipSuccesses && entry.StatusCode < 400 {
				return
			}
			// perform the logging on a seperate thread
			go logToAggregator(config, entry)
		})
	}
}

// logToAggregator makes a POST request to the log aggregator service
// to log the given newLogEntry
func logToAggregator(config *Config, entry *newLogEntry) {

	// encode the log
	payload := new(bytes.Buffer)
	err := json.NewEncoder(payload).Encode(entry)
	if err != nil {
		handleErr(config, fmt.Sprintf("Failed to marshal log entry: %v", err))
		return
	}

	// ensure non-zero timeout
	if config.Timeout == 0*time.Second {
		config.Timeout = defaultTimeout
	}
	// send the log
	client := http.Client{Timeout: config.Timeout}
	resp, err := client.Post(config.AggregatorAddress+logPath, "application/json", payload)
	if err != nil {
		handleErr(config, fmt.Sprintf("Failed POST log: %v", err))
		return
	}

	// check the aggregator response
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		if config.StdoutErrors {
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Printf("Aggregator response [%d], could not read body", resp.StatusCode)
				return
			}
			log.Printf("Aggregator response [%d]: %v", resp.StatusCode, string(body))
		}
		return
	}
}

func handleErr(config *Config, errMsg string) {
	if config.StdoutErrors {
		log.Println(errMsg)
	}
}
