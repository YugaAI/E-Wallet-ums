package cmd

import (
	"ewallet-framework/helpers"
	"log"
	"net"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func ServerGRPC() {
	list, err := net.Listen("tcp", ":"+helpers.GetEnv("GRCP_PORT", "7000"))
	if err != nil {
		log.Fatal("failed serve to GRCP", err)
	}

	s := grpc.NewServer()

	logrus.Info("start listening GRPC on port: " + helpers.GetEnv("GRCP_PORT", "7000"))
	if err := s.Serve(list); err != nil {
		log.Fatal("failed serve to GRCP", err)
	}
}
