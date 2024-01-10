package repositories

import (
	"errors"
	"github.com/kam2yar/user-service/internal/database/entities"
	"github.com/kam2yar/user-service/internal/dto"
)

type UserRepositoryInterface interface {
	List(limit int) *[]dto.UserDto
	FindByID(id uint) (dto.UserDto, error)
	Create(userDto *dto.UserDto) error
	Update(userDto *dto.UserDto) error
	Delete(id uint) error
}

type UserDatabaseRepository struct{}

func (u *UserDatabaseRepository) FindByID(id uint) (dto.UserDto, error) {
	var user = entities.User{ID: id}
	result := dbc.Limit(1).First(&user, id)

	if result.RowsAffected == 0 {
		return dto.UserDto{}, errors.New("user not found")
	}

	if result.Error != nil {
		return dto.UserDto{}, result.Error
	}

	userDto := dto.UserDto{}
	fillUserDto(&user, &userDto)

	return userDto, nil
}

func (u *UserDatabaseRepository) List(limit int) *[]dto.UserDto {
	var users []entities.User
	dbc.Limit(limit).Order("id asc").Find(&users)

	var result []dto.UserDto
	for _, user := range users {
		userDto := dto.UserDto{}
		fillUserDto(&user, &userDto)

		result = append(result, userDto)
	}

	return &result
}

func (u *UserDatabaseRepository) Create(userDto *dto.UserDto) error {
	user := entities.User{
		Name:     userDto.GetName(),
		Email:    userDto.GetEmail(),
		Password: userDto.GetPassword(),
	}

	result := dbc.Create(&user)
	if result.Error != nil {
		return result.Error
	}

	fillUserDto(&user, userDto)
	return nil
}

func (u *UserDatabaseRepository) Update(userDto *dto.UserDto) error {
	user := entities.User{
		ID:        userDto.GetId(),
		Name:      userDto.GetName(),
		Email:     userDto.GetEmail(),
		Password:  userDto.GetPassword(),
		CreatedAt: userDto.GetCreatedAt(),
		UpdatedAt: userDto.GetUpdatedAt(),
		DeletedAt: userDto.GetDeletedAt(),
	}

	result := dbc.Save(&user)

	if result.Error != nil {
		return result.Error
	}

	fillUserDto(&user, userDto)
	return nil
}

func (u *UserDatabaseRepository) Delete(id uint) error {
	var user = entities.User{ID: id}
	result := dbc.Delete(&user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func fillUserDto(u *entities.User, userDto *dto.UserDto) {
	userDto.SetId(u.ID)
	userDto.SetEmail(u.Email)
	userDto.SetName(u.Name)
	userDto.SetPassword(u.Password)
	userDto.SetCreatedAt(u.CreatedAt)
	userDto.SetUpdatedAt(u.UpdatedAt)
	userDto.SetDeletedAt(u.DeletedAt)
}
