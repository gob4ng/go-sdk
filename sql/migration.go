package sql

import (
	"fmt"
	zapLog "github.com/gob4ng/go-sdk/log"
	"os"

	"errors"
	"github.com/h4lim/go-sdk/logging"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	INTERNAL       = "INTERNAL"
	POSGRES_CONFIG = "user=%s password=%s dbname=%s host=%s port=%s sslmode=%s"
	MYSQL_CONFIG   = "%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local"
)

var log = logging.MustGetLogger("go-sdk")

type DBModel struct {
	ServerMode    string
	Driver        string
	Host          string
	Port          string
	Name          string
	Username      string
	Password      string
	ZapLogContext *zapLog.ZapLogContext
}

type LogTracking struct {
	gorm.Model

	LogID          string `db:"log_id"`
	LogType        string `db:"log_type"`
	Severity       string `db:"severity"`
	UnixTimestamp  int64  `db:"unix_timestamp"`
	Environment    string `db:"environment"`
	ClientName     string `db:"client_name"`
	Duration       int64  `db:"duration"`
	CustomMessage  string `db:"custom_message"`
	ServerUrl      string `gorm:"size:5000000"`
	ServerMethod   string `db:"server_method"`
	ServerHeader   string `gorm:"size:5000000"`
	ServerRequest  string `gorm:"size:5000000"`
	ServerResponse string `gorm:"size:5000000"`
	ServerHttpCode int    `db:"server_http_code"`
	ClientUrl      string `gorm:"size:5000000"`
	ClientMethod   string `db:"server_method"`
	ClientHeader   string `gorm:"size:5000000"`
	ClientRequest  string `gorm:"size:5000000"`
	ClientResponse string `gorm:"size:5000000"`
	ClientHttpCode int    `db:"client_http_code"`
}

func migration(db *gorm.DB) {
	db.AutoMigrate(&LogTracking{})
}

func (c *DBModel) InitDB() (*gorm.DB, *error) {

	if c.ZapLogContext == nil {
		newErr := errors.New("please define zap log context")
		return nil, &newErr
	}

	db, err := dBOpen(c)
	migration(db)
	if err != nil {
		log.Errorf(INTERNAL, "Error When Open DB %s ", err)
		return nil, err
	}
	return db, nil
}

func (c *DBModel) DBOpen() (*gorm.DB, *error) {

	if c.ZapLogContext == nil {
		newErr := errors.New("please define zap log context")
		return nil, &newErr
	}

	db, err := dBOpen(c)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func dBOpen(c *DBModel) (*gorm.DB, *error) {

	var connectionUrl string
	switch c.Driver {
	case "postgres":
		connectionUrl = fmt.Sprintf(POSGRES_CONFIG, c.Username, c.Password, c.Name, c.Host, c.Port, "disable")
	case "mysql":
		connectionUrl = fmt.Sprintf(MYSQL_CONFIG, c.Username, c.Password, c.Host, c.Port, c.Name)
	default:
		log.Errorf(logging.INTERNAL, "No Database Selected!, Please check config.toml")
		os.Exit(1)
	}

	db, err := gorm.Open(c.Driver, connectionUrl)
	if err != nil {
		log.Errorf(logging.INTERNAL, "Cannot Connect to DB With Message ", err.Error())
		return nil, &err
	}

	return db, nil
}
