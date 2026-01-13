package main

import (
	"ewallet-framework/cmd"
	"ewallet-framework/helpers"
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
