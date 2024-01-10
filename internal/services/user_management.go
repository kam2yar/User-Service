package services

import (
	"github.com/kam2yar/user-service/internal/database/repositories"
	"github.com/kam2yar/user-service/internal/dto"
	"golang.org/x/crypto/bcrypt"
	"log"
)

var userRepository repositories.UserRepositoryInterface = &repositories.UserDatabaseRepository{}

func CreateUser(userDto *dto.UserDto) error {
	password := userDto.GetPassword()
	hashed, err := HashPassword(password)
	if err != nil {
		log.Println("hash user password failed", err)
		return err
	}

	userDto.SetPassword(hashed)

	err = userRepository.Create(userDto)
	if err != nil {
		log.Println("create user failed: ", err)
		return err
	}

	return nil
}

func FindUser(id uint) (*dto.UserDto, error) {
	userDto, err := userRepository.FindByID(id)
	if err != nil {
		log.Println("find user failed: ", err)
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
		log.Println("hash user password failed", err)
		return err
	}

	userDto.SetPassword(hashed)

	err = userRepository.Update(userDto)
	if err != nil {
		log.Println("update user failed: ", err)
		return err
	}

	return nil
}

func DeleteUser(id uint) error {
	err := userRepository.Delete(id)
	if err != nil {
		log.Println("delete user failed: ", err)
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
