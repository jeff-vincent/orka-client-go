package main

import (
	"fmt"
	"orka-client-go/orka"
	"os"
)

var host string = os.Getenv("ORKA_IP")
var email string = os.Getenv("ORKA_EMAIL")
var password string = os.Getenv("ORKA_PASS")
var license_key string = os.Getenv("ORKA_LICENSE_KEY")

func main() {
	client, err := orka.NewClient(&host, &email, &password, &license_key)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("-------- CREATING VM ---------")
	c, err := client.CreateVM()

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(c)

	fmt.Println("-------- DEPLOYING VM ---------")
	d, err := client.DeployVM()

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(d)

	fmt.Println("-------- GETTING VMs ---------")
	vms, err := client.GetVMs()

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(vms)

	fmt.Println("-------- GETTING VM Status ---------")
	vm_s, err := client.GetVMStatus("myorkavm")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(vm_s)

	fmt.Println("-------- GETTING VM Configs ---------")
	vm_cfgs, err := client.GetVMConfigs()

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(vm_cfgs)

}
