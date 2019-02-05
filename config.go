package main

/*
Config haa three fields,
"Enable" tells us if application should return data or not,
"DatabasePath" tells where our database is,
"Port" tells us port for connect with database.
*/
type Config struct {
	Enabled      bool
	DatabasePath string
	Port         string
}

func NewConfig() *Config {
	return &Config{
		Enabled:      true,
		DatabasePath: "./example.db",
		Port:         "8000",
	}
}
