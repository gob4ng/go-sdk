package sql

import (
	"database/sql"
	zapLog "github.com/gob4ng/go-sdk/log"
	"github.com/jinzhu/gorm"
)

type DBOperation interface {
	InitDB() (*gorm.DB, *error)
	DBOpen() (*gorm.DB, *error)
	Select(query string, args ...interface{}) (*gorm.DB, *sql.Rows, *error)
	Insert(query string, args ...interface{}) (*gorm.DB, *error)
	Update(query string, args ...interface{}) (*gorm.DB, *error)
	Delete(query string, args ...interface{}) (*gorm.DB, *error)
}

func (c *DBModel) Select(query string, ctx *zapLog.ZapTrackingContext, args ...interface{}) (*gorm.DB, *sql.Rows, *error) {

	db, err := dBOpen(c)
	if err != nil {
		return db, nil, err
	}

	res, rowError := db.Raw(query, args...).Rows()
	if rowError != nil {
		return db, nil, &rowError
	}

	return db, res, nil
}

func (c *DBModel) Insert(query string, args ...interface{}) (*gorm.DB, *error) {

	db, err := dBOpen(c)
	if err != nil {
		return db, err
	}

	tx := db.Begin()
	tx.Exec(query, args...)
	if err != nil {
		tx.Rollback()
		return db, err
	}

	tx.Commit()
	return db, nil
}

func (c *DBModel) Update(query string, args ...interface{}) (*gorm.DB, *error) {

	db, err := dBOpen(c)
	if err != nil {
		return db, err
	}

	tx := db.Begin()
	tx.Exec(query, args...)
	if err != nil {
		tx.Rollback()
		return db, err
	}

	tx.Commit()
	return db, nil
}

func (c *DBModel) Delete(query string, args ...interface{}) (*gorm.DB, *error) {

	db, err := dBOpen(c)
	if err != nil {
		return db, err
	}

	db.Exec(query, args...)
	if err != nil {
		return db, err
	}

	return db, nil
}
