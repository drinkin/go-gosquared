package gosquared

import (
	"net/http"

	"github.com/dghubble/sling"
)

const gosquaredRoot = "https://api.gosquared.com/tracking/v1/"

// Client is a GoSquared client
type Client struct {
	sling     *sling.Sling
	APIKey    string
	SiteToken string
}

// NewClient returns a new Client.
func NewClient(api_key string, site_token string, httpClients ...*http.Client) *Client {
	var httpClient *http.Client

	if len(httpClients) > 0 {
		httpClient = httpClients[0]
	} else {
		httpClient = http.DefaultClient
	}

	base := sling.New().Client(httpClient).Base(gosquaredRoot)
	return &Client{
		sling:     base,
		APIKey:    api_key,
		SiteToken: site_token,
	}
}

type IdentifyParams struct {
	PersonID   string  `json:"person_id"`
	Properties *Person `json:"properties"`
}

type IdentifyResponse struct {
	SentProps bool `json:"sentProps"`
	SentAlias bool `json:"sentAlias"`
	Success   bool `json:"success"`
}

func (c *Client) Identify(id string, properties ...*Person) (*IdentifyResponse, error) {
	var prop *Person
	if len(properties) > 0 {
		prop = properties[0]
	}
	body := &IdentifyParams{
		PersonID:   id,
		Properties: prop,
	}

	var resp IdentifyResponse
	_, err := c.sling.Post("identify/").BodyJSON(body).ReceiveSuccess(&resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil

}
