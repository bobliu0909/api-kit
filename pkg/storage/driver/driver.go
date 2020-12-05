package driver

const (
	NGCloudStorageDriver = "ngCloud"
	MongoDBStorageDriver = "mongoDB"
)

type IStorageDriver interface {
	Name() string
	Location() string
	Open() error
	Close()
}

