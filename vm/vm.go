package vm

import (
	"github.com/Avengerist/ProxmoxGo/client"
	"github.com/Avengerist/ProxmoxGo/requests"
)
type Qemu struct {
	node   string `json:"node"`
	vmid   int    `json:"vmid"`
	cores  int    `json:"cores"`
	memory int    `json:"memory"`
	onboot bool   `json:"onboot"`
	pool   string `json:"pool"`
}


func CreateVM(c *client.Client,vmdata map[string]interface{}) {
	request.SendPost(c,"/api2/json/nodes/pve-1/qemu",vmdata)
}