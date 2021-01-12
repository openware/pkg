package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Internal database pointer
var db *gorm.DB

// Connect to database MySQL/SQLite using gorm
// gorm (GO ORM for SQL): http://gorm.io/docs/connecting_to_the_database.html
// TODO Switch to Config struct
func Connect(cnf *Config) (*gorm.DB, error) {

	var err error
	var dial gorm.Dialector

	switch cnf.Driver {
	case "memory":
		dial = sqlite.Open(":memory:")

	case "mysql":
		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			cnf.User, cnf.Pass, cnf.Host, cnf.Port, cnf.Name,
		)
		dial = mysql.Open(dsn)

	default:
		return nil, fmt.Errorf("Unsupported DATABASE_DRIVER: %s", cnf.Driver)
	}

	db, err = gorm.Open(dial, &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// FIXME delete
	sql, err := db.DB()
	if err != nil {
		return nil, err
	}

	// FIXME: move into switch case use DSN; won't work with sqlite
	sql.SetMaxOpenConns(cnf.Pool)
	return db, nil
}
