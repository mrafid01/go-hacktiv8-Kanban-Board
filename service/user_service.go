package service

import (
	"errors"
	"project3/middleware"
	"project3/model/entity"
	"project3/model/input"
	"project3/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(userInput input.UserRegisterInput) (entity.User, error)
	CreateAdmin(input input.UserRegisterInput) (entity.User, error)
	LoginUser(userInput input.UserLoginInput) (string, error)
	UpdateUser(id_user int, input input.UserUpdateInput) (entity.User, error)
	DeleteUser(id_user int) (entity.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *userService {
	return &userService{userRepository}
}

func (s *userService) CreateUser(input input.UserRegisterInput) (entity.User, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return entity.User{}, err
	}

	user := entity.User{
		FullName: input.FullName,
		Email:    input.Email,
		Password: string(passwordHash),
		Role:     "member",
	}

	userData, err := s.userRepository.CreateUser(user)
	if err != nil {
		return entity.User{}, err
	}

	return userData, nil
}

func (s *userService) CreateAdmin(input input.UserRegisterInput) (entity.User, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return entity.User{}, err
	}

	user := entity.User{
		FullName: input.FullName,
		Email:    input.Email,
		Password: string(passwordHash),
		Role:     "admin",
	}

	userData, err := s.userRepository.CreateUser(user)
	if err != nil {
		return entity.User{}, err
	}

	return userData, nil
}

func (s *userService) LoginUser(userInput input.UserLoginInput) (string, error) {
	userData, err := s.userRepository.FindByEmail(userInput.Email)
	if err != nil {
		return "", err
	}
	if userData.ID == 0 {
		return "", errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(userInput.Password))
	if err != nil {
		return "", errors.New("wrong password")
	}

	token, err := middleware.GenerateToken(userData.ID, userData.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *userService) UpdateUser(id_user int, input input.UserUpdateInput) (entity.User, error) {
	user := entity.User{
		FullName: input.FullName,
		Email:    input.Email,
	}

	_, err := s.userRepository.Update(id_user, user)
	if err != nil {
		return entity.User{}, err
	}

	userData, err := s.userRepository.FindByID(id_user)
	if err != nil {
		return entity.User{}, err
	}

	return userData, nil
}

func (s *userService) DeleteUser(id_user int) (entity.User, error) {
	userData, err := s.userRepository.FindByID(id_user)
	if err != nil {
		return entity.User{}, err
	}
	if userData.ID == 0 {
		return entity.User{}, errors.New("user not found")
	}

	userData, err = s.userRepository.Delete(id_user)
	if err != nil {
		return entity.User{}, err
	}

	return userData, nil
}
