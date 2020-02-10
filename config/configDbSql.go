package config

import (
	"database/sql"
	"fmt"
	_ "github.com/newrelic/go-agent/_integrations/nrpq"
	"log"
)

// GlobalDbSQL is a global variable to handle sql database connection
var GlobalDbSQL *sql.DB

// InitDb is a func to initialize sql database connection
func InitDb(dbConfig SQLDatabase) (con *sql.DB, err error) {
	connInfo := fmt.Sprintf(`%s:%s@tcp(%s:%s)/%s`, dbConfig.User, dbConfig.Password,
		dbConfig.Host, dbConfig.Port, dbConfig.Name)
	switch dbConfig.Connection {
	case "nrpostgres":
		connInfo = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Name, dbConfig.Password)
		break
	}
	con, err = sql.Open(dbConfig.Connection, connInfo)
	if err != nil {
		log.Panic(err)
	}

	err = con.Ping()
	if err != nil {
		log.Panic(err)
	}


	return con, err
}

// built-in function which start before the main function
func init() {
	db := SQLDatabase{
		Name:       "kygo",
		User:       "postgres",
		Password:   "postgres",
		Port:       "5432",
		Connection: "nrpostgres",
		Host:       "127.0.0.1",
	}

	GlobalDbSQL, _ = InitDb(db)
}

