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
	client.CreateVM()
	fmt.Println("-------- DEPLOYING VM ---------")
	client.DeployVM()
	fmt.Println("-------- GETTING VMs ---------")
	client.GetVMs()
}
