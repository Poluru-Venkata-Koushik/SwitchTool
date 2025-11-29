package db

import (
	"SwitchTool/schema"
	"net"
)

var (

	Inventory = []schema.Device{
		{Hostname: "Arista_1", IP: net.ParseIP("192.168.1.1"),  Model: "Arista DCS 7130"},
	}
)