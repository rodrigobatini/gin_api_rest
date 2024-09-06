package repositories

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"gin_api_rest/config"
	"gin_api_rest/models"

	_ "github.com/lib/pq"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() (*UserRepository, error) {
	db, err := config.GetDb()

	if err != nil {
		return nil, errors.Wrap(err, "erro ao conectar ao banco de dados")
	}

	return &UserRepository{db: db}, nil
}

func (ur *UserRepository) GetUsers() ([]models.User, error) {
	ur, err := NewUserRepository()
	if err != nil {
		return nil, errors.Wrap(err, "erro ao conectar ao banco de dados")
	}
	var users []models.User
	err_find := ur.db.Find(&users).Error
	if err_find != nil {
		return nil, errors.Wrap(err, "erro ao consultar usuários")
	}
	return users, nil
}

func (ur *UserRepository) GetUserByID(id int) (*models.User, error) {
	var user models.User
	err := ur.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("usuário não encontrado")
		}
		return nil, errors.Wrap(err, "erro ao consultar usuário")
	}

	return &user, nil
}

func (ur *UserRepository) CreateUser(user *models.User) error {
	err := ur.db.Create(user).Error
	if err != nil {
		return errors.Wrap(err, "erro ao criar usuário")
	}

	return nil
}

func (ur *UserRepository) UpdateUser(user *models.User) error {
	err := ur.db.Model(&models.User{}).Where("id = ?", user.ID).Updates(map[string]interface{}{"name": user.Name, "email": user.Email}).Error
	if err != nil {
		return errors.Wrap(err, "erro ao atualizar usuário")
	}

	return nil
}

func (ur *UserRepository) DeleteUser(id int) error {
	err := ur.db.Delete(&models.User{}, "id = ?", id).Error
	if err != nil {
		return errors.Wrap(err, "erro ao deletar usuário")
	}

	return nil
}
