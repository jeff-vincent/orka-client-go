package orka

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) GetVMs([]VM, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/resources/vm/list", c.HostURL), nil)

	if err != nil {
		return
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return
	}
	vms := []VM{}
	data := json.Unmarshal(body, &vms)
	fmt.Println(data)
}
