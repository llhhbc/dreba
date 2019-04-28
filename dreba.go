package main

import (
	"flag"
	"fmt"
	"github.com/spf13/pflag"
	"net"
	"github.com/gin-gonic/gin"
	"dreba/handler"
	"dreba/models"
)

var (
	port = pflag.Int("port", 9080, "The port to listen to for incoming HTTP requests.")
	addr = pflag.IP("bind-address", net.IPv4(0, 0, 0, 0), "The IP address on which to serve the --port (set to 0.0.0.0 for all interfaces).")
)

func main() {

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	flag.CommandLine.Parse(make([]string, 0)) // Init for glog calls in kubernetes packages

	r := gin.Default()

	rg := r.Group(handler.BaseUrl)

	handler.LoadHandler(rg)

	models.InitDb("root:@tcp(localhost:3306)/dreba?charset=utf8&parseTime=True&loc=Local")
	//models.InitDb("root:dreba@tcp(dreba-db:3306)/dreba?charset=utf8&parseTime=True&loc=Local")

	r.Run(fmt.Sprintf("%s:%d", *addr, *port))
}

