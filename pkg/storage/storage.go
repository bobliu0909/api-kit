package storage

import (
	"fmt"

	"github.com/rl5c/api-gin/pkg/storage/driver"
	_ "github.com/rl5c/api-gin/pkg/storage/driver/mongo"
	_ "github.com/rl5c/api-gin/pkg/storage/driver/ngcloud"
	"github.com/rl5c/api-gin/pkg/storage/factory"
)

func StorageFactory(location string, config map[string]interface{}) (driver.IStorageDriver, error) {
	driverConfig := factory.ParseStorageDriverConfig(location, config)
	if driverConfig != nil {
		return factory.Create(driverConfig)
	}
	return nil, fmt.Errorf("parse storage factory driver config invalid.")
}