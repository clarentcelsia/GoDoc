package config

import (
	"database/sql"
	"fmt"

	"github.com/spf13/viper"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func SetConfig() {
	viper.SetConfigFile("database/config/config.json")
}

func Connect() (*sql.DB, error) {
	db := dbConfig(server(), user(), "", scheme())
	mssql, err := sql.Open("mssql", db)

	if err != nil {
		return nil, err
	}

	if errp := mssql.Ping(); errp != nil {
		return nil, errp
	}

	return mssql, nil
}

func server() string {
	return viper.GetString("server")
}

func user() string {
	return viper.GetString("user")
}

func scheme() string {
	return viper.GetString("scheme")
}

func dbConfig(server, user, pass, database string) string {
	return fmt.Sprintf("server=%s; user id=%s; password=%s; database=%s;",
		server,
		user,
		pass,
		database,
	)
}
