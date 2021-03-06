package database

import (
    "database/sql"
    "fmt"
    "github.com/Pepito-Manaloto/vocabulary-api/pkg/file"
    _ "github.com/go-sql-driver/mysql"
    "github.com/pkg/errors"
)

const (
    DbDriver = "mysql"
)

func Connect(config *file.Config) (db *sql.DB, err error) {

    dataSource := dataSourceFromConfig(config)

    db, err = sql.Open(DbDriver, dataSource)

    err = errors.Wrapf(err, "Connect. Error connecting to database. dataSource=%s", dataSource)

    return
}

func dataSourceFromConfig(config *file.Config) string {
    dbConfig := config.Database

    const dataSource = "%s:%s@tcp(%s)/%s?charset=%s&multiStatements=true"

    return fmt.Sprintf(dataSource, dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Schema, dbConfig.Charset)
}
