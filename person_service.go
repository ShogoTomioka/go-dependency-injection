package main

type UserService struct {
	repository *UserRepository
}

func (service *UserService) FindById(uid int) *User {
	return service.repository.FindById(uid)
}

func (service *UserService) FindAll() []*User {
	return service.repository.FindAll()
}

func NewUserService(repository *UserRepository) *UserService {
	return &UserService{repository: repository}
}
