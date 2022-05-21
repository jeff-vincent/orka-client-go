package orka

import "time"

// TODO: define orka objects
type VM struct {
	VirtualMachineName string `json:"virtual_machine_name"`
	VMDeploymentStatus string `json:"vm_deployment_status"`
	Status             []struct {
		Owner                 string `json:"owner"`
		VirtualMachineName    string `json:"virtual_machine_name"`
		VirtualMachineID      string `json:"virtual_machine_id"`
		NodeLocation          string `json:"node_location"`
		NodeStatus            string `json:"node_status"`
		VirtualMachineIP      string `json:"virtual_machine_ip"`
		VncPort               string `json:"vnc_port"`
		ScreenSharingPort     string `json:"screen_sharing_port"`
		SSHPort               string `json:"ssh_port"`
		CPU                   int    `json:"cpu"`
		Vcpu                  int    `json:"vcpu"`
		Gpu                   string `json:"gpu"`
		RAM                   string `json:"RAM"`
		BaseImage             string `json:"base_image"`
		Image                 string `json:"image"`
		ConfigurationTemplate string `json:"configuration_template"`
		VMStatus              string `json:"vm_status"`
		IoBoost               bool   `json:"io_boost"`
		UseSavedState         bool   `json:"use_saved_state"`
		ReservedPorts         []struct {
			HostPort  int    `json:"host_port"`
			GuestPort int    `json:"guest_port"`
			Protocol  string `json:"protocol"`
		} `json:"reserved_ports"`
		CreationTimestamp time.Time `json:"creationTimestamp"`
		Tag               string    `json:"tag"`
		TagRequired       bool      `json:"tag_required"`
		Replicas          int       `json:"replicas"`
	} `json:"status"`
}
