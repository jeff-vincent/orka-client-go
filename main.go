package main

import (
	"fmt"
	"orka-client-go/orka"
)

var host string = "-------"
var email string = "-------"
var password string = "------"

func main() {
	client, err := orka.NewClient(&host, &email, &password)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(client)
}
