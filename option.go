package tiktok

import (
	"net/http"
	"time"
)

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

type Logger interface {
	Print(...interface{})
	Printf(string, ...interface{})
}

type option struct {
	client HTTPClient
	logger Logger
}

type Option func(o *option)

// WithHTTPClient setup HTTPClient with given client.
func WithHTTPClient(c HTTPClient) Option {
	return func(o *option) {
		o.client = c
	}
}

// WithLogger setup logger for debug perporse.
// Default Logger will be discard.
func WithLogger(l Logger) Option {
	return func(o *option) {
		o.logger = l
	}
}

type nopLogger struct{}

func (n *nopLogger) Print(...interface{}) {
}

func (n *nopLogger) Printf(string, ...interface{}) {
}

func defaultOpt() *option {
	return &option{
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
		logger: &nopLogger{},
	}
}
