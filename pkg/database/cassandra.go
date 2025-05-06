package database

import (
	"github.com/gocql/gocql"
	"url-shortener/config"
)

func NewCassandraCluster(config *config.Config) *gocql.ClusterConfig {
	cluster := gocql.NewCluster(config.Db.Host)
	cluster.Consistency = gocql.Quorum
	cluster.Keyspace = config.Db.Keyspace
	return cluster
}
