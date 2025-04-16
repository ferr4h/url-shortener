package database

import (
	"github.com/gocql/gocql"
	"url-shortener/config"
)

func NewCassandraCluster(config *config.Config) *gocql.ClusterConfig {
	cluster := gocql.NewCluster(config.Db.Host)
	return cluster
}
