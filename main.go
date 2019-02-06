package main

func main() {
	config := NewConfig()

	db, err := ConnectDatabase(config)

	if err != nil {
		panic(err)
	}

	userRepository := NewUserRepository(db)

	personService := NewUserService(userRepository)

	server := NewServer(personService)

	server.Run()
}
