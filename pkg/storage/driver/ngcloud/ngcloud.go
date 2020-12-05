package ngcloud

import (
	"fmt"
	"net/url"
	"path"
	"time"

	"github.com/rl5c/api-gin/pkg/storage/driver"
	"github.com/rl5c/api-gin/pkg/storage/factory"
)

const (
	driverName = driver.NGCloudStorageDriver
	defaultPageSize                = 128
	defaultKeepAlive               = "60s"
	defaultTimeout                 = "30s"
	defaultHTTPMaxIdleConns        = 50
	defaultHTTPMaxIdleConnsPerHost = 50
	defaultHTTPIdleConnTimeout     = "90s"
	sysconfigsCollection           = "sysconfigs"
	clustersCollection             = "clusters"
	servicesCollection             = "services"
)

func Driver() string {
	return driverName
}

func init() {
	factory.Register(driverName, &ngCloudDriverFactory{})
}

type ngCloudDriverFactory struct{}

func (factory *ngCloudDriverFactory) Create(config *factory.StorageDriverConfig) (driver.IStorageDriver, error) {
	driverConfig, err := parseDriverConfig(config.Location, config.Options)
	if err != nil {
		return nil, err
	}
	return NewStorageDriver(driverConfig)
}

func parseDriverConfig(location string, options map[string]interface{}) (*driverConfig, error) {
	keepAlive, _ := time.ParseDuration(defaultKeepAlive)
	timeout, _ := time.ParseDuration(defaultTimeout)
	config := &driverConfig{
		Location:  location,
		PageSize:  defaultPageSize,
		KeepAlive: keepAlive,
		Timeout:   timeout,
	}

	u, err := url.Parse(fmt.Sprint(options["url"]))
	if err != nil {
		return nil, fmt.Errorf("%s 'url' invalid", driverName)
	}

	scheme := u.Scheme
	if scheme == "" {
		scheme = "http"
	}

	rawURL := scheme + "://" + u.Host + path.Clean(u.Path)
	if u.RawQuery != "" {
		rawURL = rawURL + "?" + u.RawQuery
	}
	config.URL = rawURL

	if value, ret := options["pageSize"]; ret {
		pageSize := value.(int)
		if pageSize < 0 {
			return nil, fmt.Errorf("%s 'pageSize' invalid", driverName)
		}
		config.PageSize = pageSize
	}

	if value, ret := options["keepAlive"]; ret {
		keepAlive, err := time.ParseDuration(fmt.Sprint(value))
		if err != nil {
			return nil, fmt.Errorf("%s 'keepAlive' invalid", driverName)
		}
		config.KeepAlive = keepAlive
	}

	if value, ret := options["timeout"]; ret {
		timeout, err := time.ParseDuration(fmt.Sprint(value))
		if err != nil {
			return nil, fmt.Errorf("%s 'timeout' invalid", driverName)
		}
		config.Timeout = timeout
	}
	return config, nil
}

type driverConfig struct {
	Location  string
	URL       string
	PageSize  int
	KeepAlive time.Duration
	Timeout   time.Duration
}

type ngCloudStorageDriver struct {
	config *driverConfig
}

func NewStorageDriver(config *driverConfig) (driver.IStorageDriver, error) {
	return &ngCloudStorageDriver{
		config: config,
	}, nil
}

func (driver *ngCloudStorageDriver) Name() string {
	return Driver()
}

func (driver *ngCloudStorageDriver) Location() string {
	return driver.config.Location
}

func (driver *ngCloudStorageDriver) Open() error {
	return nil
}

func (driver *ngCloudStorageDriver) Close() {
}
