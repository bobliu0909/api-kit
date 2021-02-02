package app

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/rl5c/api-server/conf"
	"github.com/rl5c/api-server/pkg/logger"
)

func Bootstrap() {
	var configFile string
	flag.StringVar(&configFile, "configFile", "", "set a config path.")
	flag.Parse()

	if configFile == "" {
		value := os.Getenv("APP_CONFIG_NAME")
		if value == "" {
			log.Printf("config env 'APP_CONFIG_NAME' invalid.")
			os.Exit(1)
		}
		configFile = fmt.Sprintf("./conf/%s.yaml", value)
	}

	var err error
	if err := conf.New(configFile); err != nil {
		log.Printf("configuration error, %s.", err.Error())
		os.Exit(1)
	}

	loggerCfg := conf.LoggerConfigValue()
	logger.OPEN(&logger.Args{
		FileName: loggerCfg.LogFile,
		Level:    loggerCfg.LogLevel,
		MaxSize:  loggerCfg.LogSize,
	})

	var daemon *AppDaemon
	stopCh := make(chan struct{})
	if daemon, err = New(stopCh); err != nil {
		logger.ERROR("server daemon construct error, %s.", err.Error())
		os.Exit(1)
	}

	logger.INFO("[#app#] server daemon starting...")
	retry := 0
	daemonCfg := conf.DaemonConfigValue()
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
	close(stopCh)
	logger.INFO("[#app#] server daemon stopped.")
	logger.CLOSE()
}
