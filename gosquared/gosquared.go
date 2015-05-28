package gosquared

import (
	"net/http"

	"github.com/dghubble/sling"
)

const gosquaredRoot = "https://api.gosquared.com/tracking/v1"

// Client is a GoSquared client
type Client struct {
	sling *sling.Sling
}

// NewClient returns a new Client.
func NewClient(httpClient *http.Client) *Client {
	base := sling.New().Client(httpClient).Base(gosquaredRoot)
	return &Client{
		sling: base,
	}
}
