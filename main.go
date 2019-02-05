package main

func main() {
	config := NewConfig()

	db, err := ConnectDatabase(config)

	if err != nil {
		panic(err)
	}

	userRepository := NewUserRepository(db)

	personService := NewUserService(config, userRepository)

	server := NewServer(config, personService)

	server.Run()
}
