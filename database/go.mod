module github.com/openware/pkg/database

go 1.18

replace github.com/openware/pkg/ika => ../ika

require (
	github.com/openware/pkg/ika v0.0.0-20220904191345-3af2ad698e3b
	gorm.io/driver/mysql v1.3.6
	gorm.io/driver/postgres v1.3.9
	gorm.io/driver/sqlite v1.3.6
	gorm.io/gorm v1.23.8
)

require (
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.12.1 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/pgtype v1.11.0 // indirect
	github.com/jackc/pgx/v4 v4.16.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/mattn/go-sqlite3 v1.14.12 // indirect
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519 // indirect
	golang.org/x/text v0.3.7 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
