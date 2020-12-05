package factory

import (
	"fmt"
)

type StorageDriverConfig struct {
	Driver     string                 `json:"driver"`
	Location   string                 `json:"location"`
	Options map[string]interface{} `json:"parameters"`
}

func ParseStorageDriverConfig(location string, config map[string]interface{}) *StorageDriverConfig {
	driverConfig := &StorageDriverConfig{Location: location}
	if config != nil {
		for driver, options := range config {
			driverConfig.Driver = driver
			driverConfig.Options = CleanupMapValue(options).(map[string]interface{})
			break
		}
	}
	return driverConfig
}

func CleanupInterfaceArray(in []interface{}) []interface{} {

	res := make([]interface{}, len(in))
	for i, v := range in {
		res[i] = CleanupMapValue(v)
	}
	return res
}

func CleanupInterfaceMap(in map[interface{}]interface{}) map[string]interface{} {

	res := make(map[string]interface{})
	for k, v := range in {
		res[fmt.Sprintf("%v", k)] = CleanupMapValue(v)
	}
	return res
}

func CleanupMapValue(v interface{}) interface{} {

	switch v := v.(type) {
	case []interface{}:
		return CleanupInterfaceArray(v)
	case map[interface{}]interface{}:
		return CleanupInterfaceMap(v)
	case string:
		return v
	case int:
		return (int)(v)
	case int8:
		return (int8)(v)
	case int16:
		return (int16)(v)
	case int32:
		return (int32)(v)
	case int64:
		return (int64)(v)
	case uint:
		return (uint)(v)
	case uint8:
		return (uint8)(v)
	case uint16:
		return (uint16)(v)
	case uint32:
		return (uint32)(v)
	case uint64:
		return (uint64)(v)
	case float32:
		return (float32)(v)
	case float64:
		return (float64)(v)
	default:
		return fmt.Sprintf("%v", v)
	}
}
