package orka

import "time"

// TODO: define orka responses
type VMs struct {
	Message                 string                   `json:"message"`
	Help                    Help                     `json:"help"`
	Errors                  []interface{}            `json:"errors"`
	VirtualMachineResources []VirtualMachineResource `json:"virtual_machine_resources"`
}

type VMCreated struct {
	Message string        `json:"message"`
	Help    Help          `json:"help"`
	Errors  []interface{} `json:"errors"`
}

type VMDeployed struct {
	Message         string        `json:"message"`
	Help            Help          `json:"help"`
	Errors          []interface{} `json:"errors"`
	RAM             string        `json:"ram"`
	Vcpu            string        `json:"vcpu"`
	HostCPU         string        `json:"host_cpu"`
	IP              string        `json:"ip"`
	SSHPort         string        `json:"ssh_port"`
	ScreenSharePort string        `json:"screen_share_port"`
	VMID            string        `json:"vm_id"`
	PortWarnings    []interface{} `json:"port_warnings"`
	IoBoost         bool          `json:"io_boost"`
	UseSavedState   bool          `json:"use_saved_state"`
	GpuPassthrough  bool          `json:"gpu_passthrough"`
	VncPort         string        `json:"vnc_port"`
}

type Help struct {
	StartVirtualMachine            string                         `json:"start_virtual_machine,omitempty"`
	StopVirtualMachine             string                         `json:"stop_virtual_machine,omitempty"`
	ResumeVirtualMachine           string                         `json:"resume_virtual_machine,omitempty"`
	SuspendVirtualMachine          string                         `json:"suspend_virtual_machine,omitempty"`
	DeployVirtualMachine           string                         `json:"deploy_virtual_machine,omitempty"`
	VirtualMachineVnc              string                         `json:"virtual_machine_vnc,omitempty"`
	RequiredRequestDataForDeploy   RequiredRequestDataForDeploy   `json:"required_request_data_for_deploy,omitempty"`
	DataForVirtualMachineExecTasks DataForVirtualMachineExecTasks `json:"data_for_virtual_machine_exec_tasks,omitempty"`
}

type RequiredRequestDataForDeploy struct {
	OrkaVMName   string `json:"orka_vm_name"`
	OrkaNodeName string `json:"orka_node_name"`
}

type DataForVirtualMachineExecTasks struct {
	OrkaVMName string `json:"orka_vm_name"`
}

type VirtualMachineResource struct {
	Message               string     `json:"message"`
	VirtualMachineName    string     `json:"virtual_machine_name"`
	VMDeploymentStatus    string     `json:"vm_deployment_status"`
	Status                []VMStatus `json:"status,omitempty"`
	Owner                 string     `json:"owner,omitempty"`
	CPU                   int        `json:"cpu,omitempty"`
	Vcpu                  int        `json:"vcpu,omitempty"`
	BaseImage             string     `json:"base_image,omitempty"`
	Image                 string     `json:"image,omitempty"`
	IoBoost               bool       `json:"io_boost,omitempty"`
	UseSavedState         bool       `json:"use_saved_state,omitempty"`
	GpuPassthrough        bool       `json:"gpu_passthrough,omitempty"`
	ConfigurationTemplate string     `json:"configuration_template,omitempty"`
	Tag                   string     `json:"tag,omitempty"`
	TagRequired           bool       `json:"tag_required,omitempty"`
}

type VMStatus struct {
	Owner                 string         `json:"owner"`
	VirtualMachineName    string         `json:"virtual_machine_name"`
	VirtualMachineID      string         `json:"virtual_machine_id"`
	NodeLocation          string         `json:"node_location"`
	NodeStatus            string         `json:"node_status"`
	VirtualMachineIP      string         `json:"virtual_machine_ip"`
	VncPort               string         `json:"vnc_port"`
	ScreenSharingPort     string         `json:"screen_sharing_port"`
	SSHPort               string         `json:"ssh_port"`
	CPU                   int            `json:"cpu"`
	Vcpu                  int            `json:"vcpu"`
	Gpu                   string         `json:"gpu"`
	RAM                   string         `json:"RAM"`
	BaseImage             string         `json:"base_image"`
	Image                 string         `json:"image"`
	ConfigurationTemplate string         `json:"configuration_template"`
	VMStatus              string         `json:"vm_status"`
	IoBoost               bool           `json:"io_boost"`
	UseSavedState         bool           `json:"use_saved_state"`
	ReservedPorts         []ReservedPort `json:"reserved_ports"`
	CreationTimestamp     time.Time      `json:"creationTimestamp"`
	Tag                   string         `json:"tag"`
	TagRequired           bool           `json:"tag_required"`
	Replicas              int            `json:"replicas"`
}

type ReservedPort struct {
	HostPort  int    `json:"host_port"`
	GuestPort int    `json:"guest_port"`
	Protocol  string `json:"protocol"`
}
