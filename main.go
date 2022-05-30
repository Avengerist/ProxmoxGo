package main

import (
	"fmt"

	"github.com/Avengerist/ProxmoxGo/client"
	"github.com/Avengerist/ProxmoxGo/vm"
)

func main() {
	var (
		url      string = "https://..."
		username string = "username@pam"
		password string = "password"
	)

	client := client.GetTicket(url, username, password)
	fmt.Println(client)

	var vmdata = map[string]interface{}{
		"node": "pve-1",
		"vmid": "41234",
	}

	vm.CreateVM(client, vmdata)
}
