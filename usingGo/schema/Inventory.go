package schema

import "net"

type Device struct{
	Hostname string `json:"Hostname"`
	IP 		 net.IP `json:"IP"`
	Model 	 string `json:"Model"`
	
}
