package mongo

import (
	"fmt"
	"strings"

	"github.com/rl5c/api-gin/pkg/storage/driver"
	"github.com/rl5c/api-gin/pkg/storage/factory"
)

const (
	driverName = driver.MongoDBStorageDriver
	defaultDataBase      = "nes"
	sysconfigsCollection = "sysconfigs"
	clustersCollection   = "clusters"
	servicesCollection   = "services"
)

func Driver() string {
	return driverName
}

func init() {
	factory.Register(driverName, &mongoDriverFactory{})
}

type mongoDriverFactory struct{}

func (factory *mongoDriverFactory) Create(config *factory.StorageDriverConfig) (driver.IStorageDriver, error) {
	driverConfig, err := parseDriverConfig(config.Location, config.Options)
	if err != nil {
		return nil, err
	}
	return NewStorageDriver(driverConfig)
}

func parseDriverConfig(location string, options map[string]interface{}) (*driverConfig, error) {
	config := &driverConfig{
		Location: location,
		DataBase: defaultDataBase,
	}

	hosts := strings.TrimSpace(fmt.Sprint(options["hosts"]))
	if hosts == "" {
		return nil, fmt.Errorf("%s 'hosts' invalid", driverName)
	}
	config.Hosts = hosts

	value, ret := options["database"]
	if ret {
		database := strings.TrimSpace(fmt.Sprint(value))
		if database == "" {
			return nil, fmt.Errorf("%s 'database' invalid", driverName)
		}
		config.DataBase = database
	}

	if value, ret := options["auth"]; ret {
		if auth, ok := value.(map[string]interface{}); ok {
			if user, ret := auth["user"]; ret {
				config.Auth["user"] = user.(string)
			}
			if password, ret := auth["password"]; ret {
				config.Auth["password"] = password.(string)
			}
		}
	}

	if value, ret = options["options"]; ret {
		if options, ok := value.([]interface{}); ok {
			for _, option := range options {
				config.Options = append(config.Options, option.(string))
			}
		}
	}
	return config, nil
}

type driverConfig struct {
	Location string
	Hosts    string
	DataBase string
	Auth     map[string]string
	Options  []string
}

type mongoDBStorageDriver struct {
	location string
}

func NewStorageDriver(config *driverConfig) (driver.IStorageDriver, error) {
	return &mongoDBStorageDriver{
		location: config.Location,
	}, nil
}

func (driver *mongoDBStorageDriver) Name() string {
	return Driver()
}

func (driver *mongoDBStorageDriver) Location() string {
	return driver.location
}

func (driver *mongoDBStorageDriver) Open() error {
	return nil
}

func (driver *mongoDBStorageDriver) Close() {
}