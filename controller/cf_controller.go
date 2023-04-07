package controller

import (
	"biz-d-service/service"
	"github.com/isyscore/isc-gobase/server"
)

func CfController() {
	// 测试入口和出口的拦截情况
	server.Get("ad/:haveMysql", service.Ad)
	server.Get("bd/:haveMysql", service.Bd)
}
