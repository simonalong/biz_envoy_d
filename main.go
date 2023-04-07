package main

import (
	"biz-d-service/config"
	"biz-d-service/router"
	"github.com/isyscore/isc-gobase/server"
)

func main() {
	router.Register()
	config.GetDb()
	server.Run()
}
