package factory

import (
	"fmt"

	"github.com/rl5c/api-gin/pkg/storage/driver"
)

// driverFactories stores an internal mapping between storage driver names and their respective
// factories
var driverFactories = make(map[string]StorageDriverFactory)

// StorageDriverFactory is a factory interface for creating storage.IStorageDriver interfaces
// Storage drivers should call Register() with a factory to make the driver available by name.
type StorageDriverFactory interface {
	// Create returns a new storage.IStorageDriver with the given parameters
	// Parameters will vary by driver and may be ignored
	// Each parameter key must only consist of lowercase letters and numbers
	Create(config *StorageDriverConfig) (driver.IStorageDriver, error)
}

// Register makes a storage driver available by the provided name.
// If Register is called twice with the same name or if driver factory is nil, it panics.
// Additionally, it is not concurrency safe. Most Storage Drivers call this function
// in their init() functions. See the documentation for StorageDriverFactory for more.
func Register(name string, factory StorageDriverFactory) {
	if factory == nil {
		panic("Must not provide nil StorageDriverFactory")
	}
	_, registered := driverFactories[name]
	if registered {
		panic(fmt.Sprintf("StorageDriverFactory named %s already registered", name))
	}
	driverFactories[name] = factory
}

// Create a new storagedriver.StorageDriver with the given name and
// parameters. To use a driver, the StorageDriverFactory must first be
// registered with the given name. If no drivers are found, an
// InvalidStorageDriverError is returned
func Create(config *StorageDriverConfig) (driver.IStorageDriver, error) {
	driverFactory, ret := driverFactories[config.Driver]
	if !ret {
		return nil, InvalidStorageDriverError{config.Driver}
	}
	return driverFactory.Create(config)
}

// InvalidStorageDriverError records an attempt to construct an unregistered storage driver
type InvalidStorageDriverError struct {
	Name string
}

func (err InvalidStorageDriverError) Error() string {
	return fmt.Sprintf("storage driver not registered: %s", err.Name)
}
