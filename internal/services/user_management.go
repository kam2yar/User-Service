package services

import (
	"github.com/kam2yar/user-service/internal/database/repositories"
	"github.com/kam2yar/user-service/internal/dto"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func CreateUser(userDto *dto.UserDto) error {
	var userRepository repositories.UserRepositoryInterface = &repositories.UserDatabaseRepository{}

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
	var userRepository repositories.UserRepositoryInterface = &repositories.UserDatabaseRepository{}

	userDto, err := userRepository.FindByID(id)
	if err != nil {
		log.Println("create user failed: ", err)
		return &userDto, err
	}

	return &userDto, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
