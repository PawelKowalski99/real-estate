package database

import (
	"database/sql"
	"errors"
	"github.com/labstack/gommon/log"
	"time"

	"github.com/prinick96/elog"
)

type Database interface {
	ConnectDB() *sql.DB
}

// Max seconds for retry a database connection
const DB_CONNECTION_TIMEOUT = 40

// Try db-data connection
func try(err error, db *sql.DB, counts *int) error {
	if err != nil || db == nil {
		// increase counter
		log.Info("Trying to connect with database", err)
		*counts++

		// can't connect with the database
		if *counts > DB_CONNECTION_TIMEOUT {
			elog.New(elog.PANIC, "Can't connect with the database", err)
		}

		// log and try again
		log.Info("Backing off for a second", err)
		time.Sleep(time.Second)

		return errors.New("can not connect to db")
	}

	return nil
}
