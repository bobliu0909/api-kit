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
	cluster string
	configuration *Configuration
)

type RetryStartup struct {
	Period time.Duration `json:"period" yaml:"period"`
	MaxRetry int `json:"maxRetry" yaml:"maxRetry"`
}

type DaemonConfig struct {
	*RetryStartup `json:"retryStartup" yaml:"retryStartup"`
}

type TLSConfig struct {
	CaCert string `json:"caCert" yaml:"caCert"`
	ServerCert string `json:"serverCert" yaml:"serverCert"`
	ServerKey string `json:"serverKey" yaml:"serverKey"`
}

type APIConfig struct {
	Bind string `json:"bind" yaml:"bind"`
	Version []string `json:"version" yaml:"version"`
	Middleware []string `json:"middleware" yaml:"middleware"`
	Debug bool `json:"debug" yaml:"debug"`
	TLSConfig *TLSConfig `json:"tlsConfig" yaml:"tlsConfig"`
}

type ServerConfig struct {
	KubeConfig string `json:"kubeConfig,omitempty" yaml:"kubeConfig,omitempty"`
	CacheRoot string `json:"cacheRoot,omitempty" yaml:"cacheRoot,omitempty"`
	MaxUpdateFailLimit int  `json:"maxUpdateFailLimit,omitempty" yaml:"maxUpdateFailLimit,omitempty"`
	DelayUpdateInterval string  `json:"delayUpdateInterval,omitempty" yaml:"delayUpdateInterval,omitempty"`
}

type LoggerConfig struct {
	LogFile  string `json:"logFile" yaml:"logFile"`
	LogLevel string `json:"logLevel" yaml:"logLevel"`
	LogSize  int64  `json:"logSize" yaml:"logSize"`
}

type Configuration struct {
	DaemonConfig `json:"daemon" yaml:"daemon"`
	APIConfig `json:"api" yaml:"api"`
	ServerConfig `json:"server" yaml:"server"`
	LoggerConfig `json:"logger" yaml:"logger"`
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
		cluster = ret[0]
	} else {
		name = "dev"
	}

	log.Printf("[#conf#] cluster: %s\n", cluster)
	log.Printf("[#conf#] daemon: %+v\n", config.DaemonConfig.RetryStartup)
	log.Printf("[#conf#] api: %+v\n", config.APIConfig)
	log.Printf("[#conf#] server: %+v\n", config.ServerConfig)
	log.Printf("[#conf#] logger: %+v\n", config.LoggerConfig)
	configuration = config
	return nil
}

func Cluster() string {
	return cluster
}

func DaemonConfigValue() *DaemonConfig {
	if configuration != nil {
		return &configuration.DaemonConfig
	}
	return nil
}

func APIConfigValue() *APIConfig {
	if configuration != nil {
		return &configuration.APIConfig
	}
	return nil
}

func ServerConfigValue() *ServerConfig {
	if configuration != nil {
		return &configuration.ServerConfig
	}
	return nil
}

func LoggerConfigValue() *LoggerConfig {
	if configuration != nil {
		return &configuration.LoggerConfig
	}
	return nil
}