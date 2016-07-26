package main

import (
	"github.com/sebdehne/ordercapture/server"
	"github.com/sebdehne/ordercapture/apiv1"
)

func main() {
	server.RunServer("ordercapture", apiv1.Config())
}
