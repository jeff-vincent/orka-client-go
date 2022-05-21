package orka

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// HostURL - Default Hashicups URL
const HostURL string = "http://10.221.188.100"

// Client -
type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Token      string
	Auth       AuthStruct
}

// AuthStruct -
type AuthStruct struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	LicenseKey string `json:"license-key"`
}

// AuthResponse -
type AuthResponse struct {
	Message string `json:"message"`
	Help    struct {
	} `json:"help"`
	Errors []interface{} `json:"errors"`
	Token  string        `json:"token"`
}

// NewClient -
func NewClient(host, email, password, license_key *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		// Default Orka URL
		HostURL: HostURL,
	}

	if host != nil {
		c.HostURL = *host
	}

	// If username or password not provided, return empty client
	if email == nil || password == nil {
		return &c, nil
	}

	c.Auth = AuthStruct{
		Email:      *email,
		Password:   *password,
		LicenseKey: *license_key,
	}

	ar, err := c.SignIn()
	if err != nil {
		return nil, err
	}

	c.Token = ar.Token

	return &c, nil
}

func (c *Client) doRequest(req *http.Request, authToken *string) ([]byte, error) {
	token := fmt.Sprintf("Bearer %s", c.Token)

	if authToken != nil {
		token = *authToken
	}

	req.Header.Set("Authorization", token)
	req.Header.Add("Content-Type", "application/json")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
