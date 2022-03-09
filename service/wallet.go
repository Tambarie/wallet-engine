package service

import (
	"github.com/Tambarie/wallet-engine/application/helpers"
	"github.com/Tambarie/wallet-engine/domain/wallet"
	"github.com/Tambarie/wallet-engine/dto"
	"github.com/google/uuid"
	"log"
	"time"
)

type WalletService interface {
	Get(id string) (dto.User, error)
	Create(userDto *dto.User) (dto.User, error)
	CheckUserExists(email string) (dto.User, bool)
}

type DefaultWalletService struct {
	repo wallet.Repository
}

func NewWalletService(repo wallet.Repository) *DefaultWalletService {
	return &DefaultWalletService{
		repo: repo,
	}
}

func (u *DefaultWalletService) Get(id string) (dto.User, error) {
	userInfo, err := u.repo.Get("iuyrwe")
	return userInfo.ToDtoResponse(), err
}

func (u *DefaultWalletService) Create(userDto *dto.User) (dto.User, error) {
	hashP, err := helpers.GenerateHashPassword(userDto.Password)
	if err != nil {
		log.Printf("An error occurred: %v", err.Error())
		return *userDto, err
	}
	userObject := wallet.User{
		ID:           uuid.NewString(),
		FirstName:    userDto.FirstName,
		LastName:     userDto.LastName,
		Email:        userDto.Email,
		HashPassword: string(hashP),
		CreatedAt:    time.Now().Unix(),
		UpdatedAt:    time.Now().Unix(),
	}
	userDB, err := u.repo.Create(&userObject)
	return userDB.ToDtoResponse(), err
}

func (u *DefaultWalletService) CheckUserExists(email string) (dto.User, bool) {
	_, err := u.repo.GetUserByEmail(email)
	return dto.User{}, err == nil
}
