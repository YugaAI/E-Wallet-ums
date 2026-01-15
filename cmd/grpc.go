package cmd

import (
	"ewallet-ums/cmd/proto/tokenvalidation"
	"ewallet-ums/helpers"
	"log"
	"net"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func ServerGRPC() {
	//init dependency
	dependency := dependencyInject()

	s := grpc.NewServer()
	// list method
	tokenvalidation.RegisterTokenValidationServer(s, dependency.ValidationTokenAPI)

	list, err := net.Listen("tcp", ":"+helpers.GetEnv("GRCP_PORT", "7000"))
	if err != nil {
		log.Fatal("failed serve to GRCP", err)
	}

	logrus.Info("start listening GRPC on port: " + helpers.GetEnv("GRCP_PORT", "7000"))
	if err := s.Serve(list); err != nil {
		log.Fatal("failed serve to GRCP", err)
	}
}
