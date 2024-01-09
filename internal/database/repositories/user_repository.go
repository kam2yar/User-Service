package repositories

import (
	"errors"
	"github.com/kam2yar/user-service/internal/database/entities"
	"github.com/kam2yar/user-service/internal/dto"
)

type UserRepositoryInterface interface {
	All(limit int) []dto.UserDto
	FindByID(id uint) (dto.UserDto, error)
	Create(userDto *dto.UserDto) (dto.UserDto, error)
	Update(userDto *dto.UserDto) (dto.UserDto, error)
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

	return convertToUserDto(&user), nil
}

func (u *UserDatabaseRepository) All(limit int) []dto.UserDto {
	var users []entities.User
	dbc.Limit(limit).Find(&users)

	var result []dto.UserDto
	for _, u := range users {
		result = append(result, convertToUserDto(&u))
	}

	return result
}

func (u *UserDatabaseRepository) Create(userDto *dto.UserDto) (dto.UserDto, error) {
	user := entities.User{
		Name:     userDto.GetName(),
		Email:    userDto.GetEmail(),
		Password: userDto.GetPassword(),
	}

	result := dbc.Create(&user)

	if result.Error != nil {
		return dto.UserDto{}, result.Error
	}

	return convertToUserDto(&user), nil
}

func (u *UserDatabaseRepository) Update(userDto *dto.UserDto) (dto.UserDto, error) {
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
		return dto.UserDto{}, result.Error
	}

	return convertToUserDto(&user), nil
}

func (u *UserDatabaseRepository) Delete(id uint) error {
	var user = entities.User{ID: id}
	result := dbc.Delete(&user)

	if result.RowsAffected == 0 {
		return errors.New("user not found")
	}

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func convertToUserDto(u *entities.User) dto.UserDto {
	user := dto.UserDto{}

	user.SetId(u.ID)
	user.SetEmail(u.Email)
	user.SetName(u.Name)
	user.SetPassword(u.Password)
	user.SetCreatedAt(u.CreatedAt)
	user.SetUpdatedAt(u.UpdatedAt)
	user.SetDeletedAt(u.DeletedAt)

	return user
}
