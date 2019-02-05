package main

import "database/sql"

type UserRepository struct {
	database *sql.DB
}

func (repository *UserRepository) FindAll() []*User {
	rows, _ := repository.database.Query(
		"SELECT id, name, age FROM users;")

	defer rows.Close()

	users := []*User{}

	for rows.Next() {
		var (
			id   int
			name string
			age  int
		)

		rows.Scan(&id, &name, &age)

		users = append(users, &User{
			Id:   id,
			Name: name,
			Age:  age,
		})
	}
	return users
}

func NewUserRepository(database *sql.DB) *UserRepository {
	return &UserRepository{database: database}
}
