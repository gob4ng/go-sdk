package nosql

import (
	"errors"
	"github.com/gob4ng/go-sdk/log"
	"github.com/gocql/gocql"
	"time"
)

type BaseModel struct {
	ID        gocql.UUID
	createdAt time.Time
	createdBy string
	updatedAt time.Time
	updatedBy string
	deletedAt time.Time
	deletedBy string
}

type CassandraSetup struct {
	keySpace   string
	clusterIps []string
	debug      bool
	logContext *log.ZapLogContext
	tracking   *log.ZapTrackingContext
}

type cassandraConfig struct {
	cassandraSetup CassandraSetup
	cluster        gocql.ClusterConfig
}

func NewConfig(setup CassandraSetup) (*cassandraConfig, *error) {
	cluster := gocql.NewCluster(setup.clusterIps...)
	cluster.Keyspace = setup.keySpace
	cluster.Consistency = gocql.Quorum

	if setup.logContext == nil {
		newError := errors.New("please define zap log context")
		return nil, &newError
	}

	if setup.tracking == nil {
		newError := errors.New("please define zap log tracking context")
		return nil, &newError
	}

	config := cassandraConfig{
		cassandraSetup: CassandraSetup{
			clusterIps: setup.clusterIps,
			keySpace:   setup.keySpace,
			debug:      setup.debug,
			logContext: setup.logContext,
			tracking:   setup.tracking,
		},
		cluster: *cluster,
	}

	return &config, nil
}

func (c cassandraConfig) Open() (*gocql.Session, *error) {
	session, err := createSession(c.cluster)

	if err != nil {
		return nil, err
	}

	return session, nil
}

func (c cassandraConfig) Create(query string, value ...string) *error {

	session, err := createSession(c.cluster)
	if err != nil {
		return err
	}

	if c.cassandraSetup.debug {
		go c.cassandraSetup.logContext.Debug(*c.cassandraSetup.tracking, query)
	}

	if err := session.Query(query, value).Exec(); err != nil {
		return &err
	}

	defer session.Close()

	return nil
}

func createSession(cluster gocql.ClusterConfig) (*gocql.Session, *error) {
	session, err := cluster.CreateSession()

	return session, &err
}
