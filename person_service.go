package main

type UserService struct {
	config     *Config
	repository *UserRepository
}

func (service *UserService) FindAll() []*User {
	return service.repository.FindAll()
}

func NewUserService(config *Config, repository *UserRepository) *UserService {
	return &UserService{config: config, repository: repository}
}
