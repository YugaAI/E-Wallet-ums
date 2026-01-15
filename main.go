package main

import (
	"ewallet-ums/cmd"
	"ewallet-ums/helpers"
)

func main() {
	//load config
	helpers.SetUpConfig()

	//loadlog
	helpers.SetupLog()

	//load database
	helpers.SetupMySQL()

	//run GRPC
	go cmd.ServerGRPC()

	//run HTTP
	cmd.ServerHTTP()
}
