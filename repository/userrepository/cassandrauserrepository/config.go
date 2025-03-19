package cassandrauserrepository

import (
	"time"

	"github.com/gocql/gocql"
)

type Config struct {
	ConnectTimeout time.Duration
	Hosts          []string
	IdleTimeout    time.Duration
	Keyspace       string
	PoolSize       int
	Table          string
	Timeout        time.Duration
}

func (config Config) toSession() (*gocql.Session, error) {
	cluster := gocql.NewCluster(config.Hosts...)

	cluster.Keyspace = config.Keyspace
	cluster.Timeout = config.Timeout
	cluster.ConnectTimeout = config.ConnectTimeout
	cluster.NumConns = config.PoolSize

	return cluster.CreateSession()
}
