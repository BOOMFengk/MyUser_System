package v1

import (
	"MyUser_System/config"
	"MyUser_System/internal/service"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// ping健康检查
func Ping(c *gin.Context) {
	appConfig := config.GetGlobalConf().AppConfig
	confInfo, _ := json.MarshalIndent(appConfig, "", " ")
	appInfo := fmt.Sprintf("app_name:%s\nversion:%s\n\n%s", appConfig.AppName, appConfig.Version,
		string(confInfo))
	c.String(http.StatusOK, appInfo)

}

func Register(c *gin.Context) {
	req := &service.RegisterRequest{}
	rsp := &HttpResponse{}

	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Errorf("request json err %v", err)
		rsp.ResponseWithError(c, CodeBodyBindErr, err.Error())
		return
	}
	if err := service.Register(req); err != nil {
		rsp.ResponseWithError(c, CodeBodyBindErr, err.Error())
		return
	}
	rsp.ResponseSuccess(c)

}
