package services

import (
	"gin_api_rest/models"
	"gin_api_rest/repositories"

	"github.com/pkg/errors"
)

type UserService struct {
	userRepository *repositories.UserRepository
}

func NewUserService() *UserService {
	return &UserService{}
}

func (us *UserService) GetUsers() ([]models.User, error) {
	users, err := us.userRepository.GetUsers()
	if err != nil {
		return nil, errors.Wrap(err, "erro ao obter usuários")
	}
	return users, nil
}

func (us *UserService) GetUserByID(id int) (*models.User, error) {
	user, err := us.userRepository.GetUserByID(id)
	if err != nil {
		return nil, errors.Wrap(err, "erro ao obter usuário por ID")
	}
	return user, nil
}

func (us *UserService) CreateUser(user *models.User) error {
	err := us.userRepository.CreateUser(user)
	if err != nil {
		return errors.Wrap(err, "erro ao criar usuário")
	}
	return nil
}

func (us *UserService) UpdateUser(user *models.User) error {
	err := us.userRepository.UpdateUser(user)
	if err != nil {
		return errors.Wrap(err, "erro ao atualizar usuário")
	}
	return nil
}

func (us *UserService) DeleteUser(id int) error {
	err := us.userRepository.DeleteUser(id)
	if err != nil {
		return errors.Wrap(err, "erro ao deletar usuário")
	}
	return nil
}
