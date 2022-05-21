package orka

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (c *Client) GetVMs() {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/resources/vm/list", c.HostURL), nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	body, err := c.doRequest(req, nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	vms := VMs{}
	errs := json.Unmarshal(body, &vms)
	if errs != nil {
		fmt.Println(errs)
	}
	fmt.Println(vms)
	return
}

// sample data
var vm_data = strings.NewReader(`{
	"orka_vm_name": "myorkavm",
	"orka_base_image": "bigsur-ssh-git.img",
	"orka_cpu_core": 6,
	"vcpu_count": 6
}`)

func (c *Client) CreateVM() {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/resources/vm/create", c.HostURL), vm_data)

	if err != nil {
		fmt.Println(err)
		return
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	vm_created := VMCreated{}
	errs := json.Unmarshal(body, &vm_created)
	if errs != nil {
		fmt.Println(errs)
	}
	fmt.Println(vm_created)
	return
}

// sample data
var deploy_data = strings.NewReader(`{
	"orka_vm_name":"myorkavm",
	"orka_node_name":"macpro-4"
}`)

func (c *Client) DeployVM() {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/resources/vm/deploy", c.HostURL), deploy_data)

	if err != nil {
		fmt.Println(err)
		return
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	vm_deployed := VMDeployed{}
	errs := json.Unmarshal(body, &vm_deployed)

	if errs != nil {
		fmt.Println(errs)
	}

	fmt.Println(vm_deployed)
	return
}
