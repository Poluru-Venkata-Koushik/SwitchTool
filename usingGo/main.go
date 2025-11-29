/*

Switchtool is a very simple project used to monitor switches and Run some snapshots on the switches
main.go is the entry point

*/

package main 

import(
	RP "SwitchTool/routerPack"
)

const(
	port = ":9999"
)

func main() {
	
	r:= RP.SetupRouter()
	r.Run(
		port,
	)
}