package main

/*
Config haa three fields,
"Enable" tells us if application should return data or not,
"DatabasePath" tells where our database is,
"Port" tells us port for connect with database.
*/

type Config struct {
	Host     string
	Database string
	DbName   string
	User     string
	Password string
	Port     string
	Protocol string
}

func NewConfig() *Config {
	return &Config{
		Host:     "localhost",
		Database: "sqlite3",
		DbName:   "testdb",
		User:     "root",
		Password: "",
		Port:     "3306",
		Protocol: "tcp",
	}
}
