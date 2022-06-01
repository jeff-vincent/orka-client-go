package orka

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (c *Client) GetVMs() (VMs, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/resources/vm/list/all", c.HostURL), nil)
	req.Header.Add("orka-licensekey", c.Auth.LicenseKey)
	if err != nil {
		fmt.Println(err)
	}

	body, err := c.doRequest(req, nil)

	if err != nil {
		return VMs{}, err
	}

	vms := VMs{}
	err = json.Unmarshal(body, &vms)
	if err != nil {
		return VMs{}, err
	}
	// fmt.Println(vms)
	return vms, nil
}

func (c *Client) GetVMStatus(vm_name string) (VMs, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/resources/vm/status/%v", c.HostURL, vm_name), nil)
	req.Header.Add("orka-licensekey", c.Auth.LicenseKey)
	if err != nil {
		fmt.Println(err)
	}

	body, err := c.doRequest(req, nil)

	if err != nil {
		return VMs{}, err
	}

	vms := VMs{}
	err = json.Unmarshal(body, &vms)
	if err != nil {
		return VMs{}, err
	}
	// fmt.Println(vms)
	return vms, nil
}

func (c *Client) GetVMConfigs() (VMConfigs, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/resources/vm/configs", c.HostURL), nil)
	req.Header.Add("orka-licensekey", c.Auth.LicenseKey)
	if err != nil {
		fmt.Println(err)
	}

	body, err := c.doRequest(req, nil)

	if err != nil {
		return VMConfigs{}, err
	}

	vm_configs := VMConfigs{}
	err = json.Unmarshal(body, &vm_configs)
	if err != nil {
		return VMConfigs{}, err
	}
	// fmt.Println(vms)
	return vm_configs, nil
}

// sample data
var vm_data = strings.NewReader(`{
	"orka_vm_name": "myorkavm",
	"orka_base_image": "bigsur-ssh-git.img",
	"orka_cpu_core": 6,
	"vcpu_count": 6
}`)

func (c *Client) CreateVM() (VMCreated, error) {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/resources/vm/create", c.HostURL), vm_data)

	if err != nil {
		return VMCreated{}, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return VMCreated{}, err
	}
	vm_created := VMCreated{}
	err = json.Unmarshal(body, &vm_created)
	if err != nil {
		return VMCreated{}, err
	}
	return vm_created, nil
}

// sample data
var deploy_data = strings.NewReader(`{
	"orka_vm_name":"myorkavm",
	"orka_node_name":"macpro-4"
}`)

func (c *Client) DeployVM() (VMDeployed, error) {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/resources/vm/deploy", c.HostURL), deploy_data)

	if err != nil {
		fmt.Println(err)
		return VMDeployed{}, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return VMDeployed{}, err
	}
	vm_deployed := VMDeployed{}
	err = json.Unmarshal(body, &vm_deployed)

	if err != nil {
		return VMDeployed{}, err
	}

	return vm_deployed, nil
}
