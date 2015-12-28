package main

import (
	"fmt"
	"github.com/rishiloyola/zoop-proxy/server"
)

func main() {
	c := proxyserver.New()
	c.Init("/HTTPserver", "/pushpin", "127.0.0.1:2181") //Init(PathOfRemoteServer1,PathOfRemoteServer2,IPofzkclient)
	fmt.Println("started")
	c.Run()
}
