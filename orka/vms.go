package orka

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (c *Client) GetVMs() {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/resources/vm/list", c.HostURL), nil)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		return
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return
	}
	vms := VMs{}
	data := json.Unmarshal(body, &vms)
	fmt.Println(data)
	return
}

// sample data
var vm_data = strings.NewReader(`{
	"orka_vm_name": "myorkavm",
	"orka_base_image": "bigsur-ssh-git.img",
	"orka_cpu_core": 6,
	"vcpu_count": 6,
}`)

func (c *Client) CreateVM() {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/resources/vm/create", c.HostURL), vm_data)
	if err != nil {
		return
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return
	}
	vm_created := VMCreated{}
	data := json.Unmarshal(body, &vm_created)
	fmt.Println(data)
	return
}
