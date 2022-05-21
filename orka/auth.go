package orka

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// SignIn - Get a new token for user
func (c *Client) SignIn() (*AuthResponse, error) {
	if c.Auth.Email == "" || c.Auth.Password == "" {
		return nil, fmt.Errorf("define email and password")
	}
	rb, err := json.Marshal(c.Auth)
	fmt.Println(string(rb))
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/token", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	body, err := c.doRequest(req, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	ar := AuthResponse{}
	err = json.Unmarshal(body, &ar)
	if err != nil {
		return nil, err
	}

	return &ar, nil
}
