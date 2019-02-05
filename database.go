package main

import (
	"database/sql"
	"fmt"
)

func ConnectDatabase(c *Config) (*sql.DB, error) {

	dbConnection := fmt.Sprintf("%s:%s@%s([%s]:%s)/%s%s",
		c.User,
		c.Password,
		c.Protocol,
		c.Host,
		c.Port,
		c.DbName)
	return sql.Open("sqlite3", dbConnection)
}
