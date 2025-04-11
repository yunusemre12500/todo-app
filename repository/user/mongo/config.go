package mongo

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Config struct {
	Collection       string
	ConnectTimeout   time.Duration
	ConnectionPool   *ConnectionPoolConfig
	Database         string
	Direct           bool
	OperationTimeout time.Duration
	ReplicaSet       string
	URI              string
}

func (c *Config) toClient() (*mongo.Client, error) {
	opts := options.Client().
		ApplyURI(c.URI).
		SetConnectTimeout(c.ConnectTimeout).
		SetDirect(c.Direct).
		SetMaxConnecting(c.ConnectionPool.MaximumConnectiong).
		SetMaxConnIdleTime(c.ConnectionPool.IdleTimeout).
		SetMaxPoolSize(c.ConnectionPool.MaximumConnections).
		SetMinPoolSize(c.ConnectionPool.MinimumConnections).
		SetReplicaSet(c.ReplicaSet).
		SetTimeout(c.OperationTimeout)

	return mongo.Connect(opts)
}

type ConnectionPoolConfig struct {
	IdleTimeout        time.Duration
	MaximumConnectiong uint64
	MaximumConnections uint64
	MinimumConnections uint64
}
