package service

import (
	"errors"
	"tokobelanja-golang/model/entity"
	"tokobelanja-golang/model/input"
	"tokobelanja-golang/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(userInput input.RegisterUserInput) (entity.User, error)
	GetUserByEmail(email string) (entity.User, error)
	TopUp(saldo int, IDUser int) (entity.User, error)
	UpdateUser(ID int, input input.UpdateUserInput) (entity.User, error)
	GetUserByID(ID int) (entity.User, error)
	DeleteUser(ID int) (entity.User, error)
	CheckUserAdmin(ID int) (bool, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *userService {
	return &userService{userRepository}
}

func (s *userService) CreateUser(input input.RegisterUserInput) (entity.User, error) {
	newUser := entity.User{}
	newUser.Email = input.Email
	newUser.FullName = input.FullName
	newUser.Role = "customer"
	newUser.Balance = 0

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	if err != nil {
		return entity.User{}, err
	}

	newUser.Password = string(passwordHash)

	checkSameUser, err := s.userRepository.CheckSameEmail(input.Email)

	if err != nil {
		return entity.User{}, err
	}

	if checkSameUser.ID != 0 {
		return entity.User{}, errors.New("Email already registered!")
	}

	createdUser, err := s.userRepository.Save(newUser)

	if err != nil {
		return entity.User{}, err
	}

	return createdUser, nil
}

func (s *userService) GetUserByEmail(email string) (entity.User, error) {
	userResult, err := s.userRepository.GetByEmail(email)

	if err != nil {
		return entity.User{}, err
	}

	return userResult, nil
}

func (s *userService) TopUp(saldo int, IDUser int) (entity.User, error) {
	userResult, err := s.userRepository.GetByID(IDUser)

	if err != nil {
		return entity.User{}, err
	}

	if userResult.ID == 0 {
		return entity.User{}, errors.New("user not found!")
	}

	saldoTerkini := userResult.Balance + saldo

	updated, err := s.userRepository.Update(IDUser, entity.User{Balance: saldoTerkini})

	if err != nil {
		return entity.User{}, err
	}

	return updated, nil
}

func (s *userService) GetUserByID(ID int) (entity.User, error) {
	user, err := s.userRepository.GetByID(ID)

	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return entity.User{}, nil
	}

	return user, nil
}

func (s *userService) UpdateUser(ID int, input input.UpdateUserInput) (entity.User, error) {
	userResult, err := s.userRepository.GetByID(ID)

	if err != nil {
		return entity.User{}, err
	}

	if userResult.ID == 0 {
		return entity.User{}, errors.New("user not found!")
	}

	updatedUser := entity.User{
		FullName: input.FullName,
		Email:    input.Email,
	}

	userUpdate, err := s.userRepository.Update(ID, updatedUser)

	if err != nil {
		return userUpdate, err
	}

	return userUpdate, nil
}

func (s *userService) DeleteUser(ID int) (entity.User, error) {

	userdata, err := s.GetUserByID(ID)

	if err != nil {
		return entity.User{}, err
	}

	if userdata.Role == "admin" {
		return entity.User{}, errors.New("Admin can not destroy self!")
	}

	_, err = s.userRepository.Delete(ID)

	if err != nil {
		return entity.User{}, err
	}

	return entity.User{}, nil

}

func (s *userService) CheckUserAdmin(ID int) (bool, error) {
	userdata, err := s.GetUserByID(ID)

	if err != nil {
		return false, err
	}

	if userdata.Role != "admin" {
		return false, nil
	}

	return true, nil
}
