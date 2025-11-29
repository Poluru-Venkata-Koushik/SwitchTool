/*

Switchtool is a very simple project used to monitor switches and Run some snapshots on the switches
main.go is the entry point

*/

package main

import (
	DB "SwitchTool/DB"
	RP "SwitchTool/routerPack"
	"log"
	"os"
	"sync"
)

const(
	port = ":9999"
)

var (
	mongoClient sync.Once
)

func main() {

	r:= RP.SetupRouter()
	mongoClient.Do(
		func() {
			MClient := os.Getenv("MD_CLIENT_STOOL")
			if MClient==""{
				Client, err := DB.ConnectToDB()
				if err!=nil{
					log.Fatalln(err)
				}else{
					log.Println(Client)
				}
			}			
		})
	r.Run(
		port,
	)
}