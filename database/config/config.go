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

func ConnectDetail() (*sql.DB, error) {
	db := dbConfig(server(), user(), "", detail_scheme())
	mssql, err := sql.Open("mssql", db)

	if err != nil {
		return nil, err
	}

	if errp := mssql.Ping(); errp != nil {
		return nil, errp
	}

	return mssql, nil
}

func ConnectAccount() (*sql.DB, error) {
	db := dbConfig(server(), user(), "", account_scheme())
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

func detail_scheme() string {
	return viper.GetString("scheme.details")
}

func account_scheme() string {
	return viper.GetString("scheme.account")
}

func dbConfig(server, user, pass, database string) string {
	return fmt.Sprintf("server=%s; user id=%s; password=%s; database=%s;",
		server,
		user,
		pass,
		database,
	)
}
