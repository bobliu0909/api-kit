package conf

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	yaml "gopkg.in/yaml.v2"
)

var (
	location string
	configuration *Configuration
)

type RetryStartup struct {
	Period time.Duration `json:"period" yaml:"period"`
	MaxRetry int `json:"maxRetry" yaml:"maxRetry"`
}

type Daemon struct {
	*RetryStartup `json:"retryStartup" yaml:"retryStartup"`
}

type TLSConfig struct {
	CaCert string `json:"caCert" yaml:"caCert"`
	ServerCert string `json:"serverCert" yaml:"serverCert"`
	ServerKey string `json:"serverKey" yaml:"serverKey"`
}

type API struct {
	Bind string `json:"bind" yaml:"bind"`
	Version []string `json:"version" yaml:"version"`
	Middleware []string `json:"middleware" yaml:"middleware"`
	Debug bool `json:"debug" yaml:"debug"`
	TLSConfig *TLSConfig `json:"tlsConfig" yaml:"tlsConfig"`
}

type Logger struct {
	LogFile  string `json:"logFile" yaml:"logFile"`
	LogLevel string `json:"logLevel" yaml:"logLevel"`
	LogSize  int64  `json:"logSize" yaml:"logSize"`
}

type Configuration struct {
	Daemon Daemon `json:"daemon" yaml:"daemon"`
	API API `json:"api" yaml:"api"`
	Cluster map[string]interface{} `json:"cluster" yaml:"cluster"`
	Storage map[string]interface{} `json:"storage" yaml:"storage"`
	Logger Logger `json:"logger" yaml:"logger"`
}

func New(filePath string) error {
	fd, err := os.OpenFile(filePath, os.O_RDONLY, 0755)
	if err != nil {
		return err
	}

	defer fd.Close()
	data, err := ioutil.ReadAll(fd)
	if err != nil {
		return err
	}

	config := &Configuration{}
	if err = yaml.Unmarshal(data, config); err != nil {
		return err
	}

	if err = config.parseEnv(); err != nil {
		return err
	}

	fi, err := fd.Stat()
	if err != nil {
		return err
	}

	name := strings.ToLower(fi.Name())
	ret := strings.SplitN(name, ".", 2)
	if len(ret) >= 2 {
		location = ret[0]
	} else {
		name = "dev"
	}

	log.Printf("[#conf#] location: %s\n", location)
	log.Printf("[#conf#] daemon: %+v\n", config.Daemon.RetryStartup)
	log.Printf("[#conf#] api: %+v\n", config.API)
	log.Printf("[#conf#] cluster: %+v\n", config.Cluster)
	log.Printf("[#conf#] storage: %+v\n", config.Storage)
	log.Printf("[#conf#] logger: %+v\n", config.Logger)
	configuration = config
	return nil
}

func Location() string {
	return location
}

func DaemonConfig() *Daemon {
	if configuration != nil {
		return &configuration.Daemon
	}
	return nil
}

func APIConfig() *API {
	if configuration != nil {
		return &configuration.API
	}
	return nil
}

func ClusterConfig() map[string]interface{} {
	if configuration != nil {
		return configuration.Cluster
	}
	return nil
}

func StorageConfig() map[string]interface{} {
	if configuration != nil {
		return configuration.Storage
	}
	return nil
}

func LoggerConfig() *Logger {
	if configuration != nil {
		return &configuration.Logger
	}
	return nil
}