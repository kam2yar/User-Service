package services

import (
	"github.com/kam2yar/user-service/internal/database/repositories"
	"github.com/kam2yar/user-service/internal/dto"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

var userRepository repositories.UserRepositoryInterface = &repositories.UserDatabaseRepository{}

func CreateUser(userDto *dto.UserDto) error {
	password := userDto.GetPassword()
	hashed, err := HashPassword(password)
	if err != nil {
		zap.L().Warn("hash user password failed", zap.Error(err))
		return err
	}

	userDto.SetPassword(hashed)

	err = userRepository.Create(userDto)
	if err != nil {
		zap.L().Warn("create user failed", zap.Error(err))
		return err
	}

	return nil
}

func FindUser(id uint) (*dto.UserDto, error) {
	userDto, err := userRepository.FindByID(id)
	if err != nil {
		zap.L().Warn("find user failed", zap.Error(err))
		return &userDto, err
	}

	return &userDto, nil
}

func List(limit int) *[]dto.UserDto {
	return userRepository.List(limit)
}

func UpdateUser(userDto *dto.UserDto) error {
	password := userDto.GetPassword()
	hashed, err := HashPassword(password)
	if err != nil {
		zap.L().Warn("hash user password failed", zap.Error(err))
		return err
	}

	userDto.SetPassword(hashed)

	err = userRepository.Update(userDto)
	if err != nil {
		zap.L().Warn("update user failed", zap.Error(err))
		return err
	}

	return nil
}

func DeleteUser(id uint) error {
	err := userRepository.Delete(id)
	if err != nil {
		zap.L().Warn("delete user failed", zap.Error(err))
		return err
	}

	return nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
