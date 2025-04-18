package database

import (
	"github.com/gocql/gocql"
	"url-shortener/config"
)

func NewCassandraCluster(config *config.Config) *gocql.ClusterConfig {
	cluster := gocql.NewCluster(config.Db.Host)
	cluster.Consistency = gocql.Quorum
	cluster.Keyspace = "url_shortener" //TODO: add to config
	return cluster
}
