package service

import (
	bizConfig "biz-d-service/config"
	"biz-d-service/pojo/domain"
	"github.com/gin-gonic/gin"
	"github.com/isyscore/isc-gobase/config"
	"github.com/isyscore/isc-gobase/http"
	"github.com/isyscore/isc-gobase/logger"
	"github.com/isyscore/isc-gobase/server/rsp"
	baseTime "github.com/isyscore/isc-gobase/time"
	netHttp "net/http"
	"time"
)

func Ad(c *gin.Context) {
	//logger.Info("接口 Ad ")
	headers := netHttp.Header{}
	if c.GetHeader("x-request-id") != "" {
		headers["x-request-id"] = []string{c.GetHeader("x-request-id")}
		headers["x-b3-traceid"] = []string{c.GetHeader("x-b3-traceid")}
		headers["x-b3-spanid"] = []string{c.GetHeader("x-b3-spanid")}
		headers["x-b3-parentspanid"] = []string{c.GetHeader("x-b3-parentspanid")}
		headers["x-b3-sampled"] = []string{c.GetHeader("x-b3-sampled")}
		headers["x-b3-flags"] = []string{c.GetHeader("x-b3-flags")}
		headers["x-ot-span-context"] = []string{c.GetHeader("x-ot-span-context")}
	}
	haveMysql := c.Param("haveMysql")
	data, err := http.GetOfStandard(config.GetValueString("biz.url.e") + "/api/e/de/" + haveMysql, headers, map[string]string{})
	if err != nil {
		logger.Error("接口 Ad 报错 %v", err.Error())
		rsp.FailedOfStandard(c, 500, err.Error())
		return
	}

	if haveMysql == "true" {
		bizConfig.Db.Model(domain.BizEnvoyD{
			Id: 1,
		}).Update("times", baseTime.TimeToStringYmdHms(time.Now()))
	}

	rsp.SuccessOfStandard(c, data)
}

func Bd(c *gin.Context) {
	//logger.Info("接口 Bd ")
	headers := netHttp.Header{}
	if c.GetHeader("x-request-id") != "" {
		headers["x-request-id"] = []string{c.GetHeader("x-request-id")}
		headers["x-b3-traceid"] = []string{c.GetHeader("x-b3-traceid")}
		headers["x-b3-spanid"] = []string{c.GetHeader("x-b3-spanid")}
		headers["x-b3-parentspanid"] = []string{c.GetHeader("x-b3-parentspanid")}
		headers["x-b3-sampled"] = []string{c.GetHeader("x-b3-sampled")}
		headers["x-b3-flags"] = []string{c.GetHeader("x-b3-flags")}
		headers["x-ot-span-context"] = []string{c.GetHeader("x-ot-span-context")}
	}
	haveMysql := c.Param("haveMysql")
	data, err := http.GetOfStandard(config.GetValueString("biz.url.e") + "/api/e/de/" + haveMysql, headers, map[string]string{})
	if err != nil {
		logger.Error("接口 Bd 报错", err.Error())
		rsp.FailedOfStandard(c, 500, err.Error())
		return
	}

	if haveMysql == "true" {
		bizConfig.Db.Model(domain.BizEnvoyD{
			Id: 1,
		}).Update("times", baseTime.TimeToStringYmdHms(time.Now()))
	}

	rsp.SuccessOfStandard(c, data)
}
