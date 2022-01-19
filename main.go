package main

import (
	"flag"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"url-shortener/common/log"
	"url-shortener/config"
	model "url-shortener/model/redis"
	"url-shortener/router"
)

var isDebugMode *bool

func init() {
	isDebugMode = flag.Bool("debugmode", false, "is in debug mode")
	configFilePath := flag.String("f", "./config.yaml", "config file location")

	flag.Parse()
	log.Initialize(*isDebugMode)

	// 初始化設定檔
	config.Setup("yaml", *configFilePath)

	if config.Root.Server.LogFile {
		log.AddFileWriter()
	}

	log.TraceLog.Info("****** Service Start ******")
}

func main() {
	redisConf := config.Root.Redis

	err := model.Dial(redisConf.Address, redisConf.Password, redisConf.DB)
	if err != nil {
		log.TraceLog.Errorf("[Redis] ping error: %s", err.Error())
		panic(err)
	}
	log.TraceLog.Infof("[Redis] connected")

	gin.SetMode(gin.ReleaseMode)

	handler := router.SetupRouter(*isDebugMode)

	s := &http.Server{
		Addr:           config.Root.Server.Port,
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
