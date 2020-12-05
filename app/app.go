package app

import (
	"log"
	"os"
	"time"

	"github.com/rl5c/api-gin/conf"
	"github.com/rl5c/api-gin/logger"
)

func Bootstrap() {
	var err error
	if err := conf.New("./conf/config.yaml"); err != nil {
		log.Printf("load configuration error, %s.", err.Error())
		os.Exit(1)
	}

	loggerCfg := conf.LoggerConfig()
	logger.OPEN(&logger.Args{
		FileName: loggerCfg.LogFile,
		Level:    loggerCfg.LogLevel,
		MaxSize:  loggerCfg.LogSize,
	})

	var daemon *AppDaemon
	if daemon, err = New(); err != nil {
		log.Printf("server daemon construct error, %s.", err.Error())
		os.Exit(1)
	}

	logger.INFO("[#app#] server daemon starting...")
	retry := 0
	daemonCfg := conf.DaemonConfig()
	for {
		if err != nil {
			if daemonCfg.RetryStartup == nil || daemonCfg.MaxRetry <= retry {
				logger.ERROR("[#app#] server daemon starting error, %s.", err.Error())
				logger.CLOSE()
				os.Exit(1)
			}
			retry = retry + 1
			logger.ERROR("[#app#] server daemon %d times start error, retry after %s again, %s.", retry, daemonCfg.Period, err.Error())
			time.Sleep(daemonCfg.Period)
		}
		if err = daemon.Startup(); err != nil {
			continue
		}
		break
	}
	logger.INFO("[#app#] server daemon started.")
	processWaitForSignal(nil)
	daemon.Stop()
	logger.INFO("[#app#] server daemon stopped.")
	logger.CLOSE()
}

