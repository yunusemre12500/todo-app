package mongodbtodorepository

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Config struct {
	Collection     string
	ConnectTimeout time.Duration
	Database       string
	IdleTimeout    time.Duration
	MaxPoolSize    uint64
	MinPoolSize    uint64
	Timeout        time.Duration
	URI            string
}

func (config Config) toClient() (*mongo.Client, error) {
	connectOptions := options.Client().
		ApplyURI(config.URI).
		SetConnectTimeout(config.ConnectTimeout).
		SetMaxConnIdleTime(config.IdleTimeout).
		SetMaxPoolSize(config.MaxPoolSize).
		SetMinPoolSize(config.MinPoolSize).
		SetTimeout(config.Timeout)

	return mongo.Connect(connectOptions)
}
