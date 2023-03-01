package orm

import (
	"github.com/gocql/gocql"
	"time"
)

type BaseModel struct {
	ID        string
	createdAt time.Time
	createdBy string
	updatedAt time.Time
	updatedBy string
	deletedAt time.Time
	deletedBy string
}

type CassandraContext struct {
	ClusterIps []string
	KeySpace   string
}

func NewConfig(keySpace string, clusterIps ...string) gocql.ClusterConfig {
	cluster := gocql.NewCluster(clusterIps...)
	cluster.Keyspace = keySpace
	cluster.Consistency = gocql.Quorum
	return *cluster
}
