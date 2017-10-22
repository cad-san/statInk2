package statink2

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

// Client is a http client for stat.ink (Splatoon2 version)
type Client struct {
	hc     *http.Client
	URL    *url.URL
	APIKey string
	Logger *log.Logger
}

const (
	statInkDomainURL = "https://stat.ink"
	battleURL        = "api/v2/battle"
)

// NewClient generates Client, http client object for stat.ink
// It provides a http client with empty cookiejar.
func NewClient(key string, logger *log.Logger) (*Client, error) {
	return newClient(statInkDomainURL, key, logger)
}

func newClient(domain, key string, logger *log.Logger) (*Client, error) {
	client := &Client{
		hc: &http.Client{},
	}
	var err error
	client.URL, err = url.ParseRequestURI(domain)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse url : %s", domain)
	}

	if len(key) == 0 {
		return nil, errors.New("missing api key")
	}
	client.APIKey = key

	if logger == nil {
		client.Logger = log.New(ioutil.Discard, "", log.LstdFlags)
	} else {
		client.Logger = logger
	}

	return client, nil
}
