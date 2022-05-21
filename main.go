package main

import (
	"fmt"
	"orka-client-go/orka"
	"os"
)

var host string = os.Getenv("ORKA_IP")
var email string = os.Getenv("ORKA_EMAIL")
var password string = os.Getenv("ORKA_PASS")

func main() {
	client, err := orka.NewClient(&host, &email, &password)

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
