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
	CreateWallet(userDto *dto.User) (dto.User, error)
}

type DefaultWalletService struct {
	repo wallet.Repository
}

func NewWalletService(repo wallet.Repository) *DefaultWalletService {
	return &DefaultWalletService{
		repo: repo,
	}
}

func (u *DefaultWalletService) CreateWallet(userDto *dto.User) (dto.User, error) {
	hashP, err := helpers.GenerateHashPassword(string(userDto.SecretKey))
	if err != nil {
		log.Printf("An error occurred: %v", err.Error())
		return *userDto, err
	}
	userObject := wallet.User{
		ID:              uuid.NewString(),
		FirstName:       userDto.FirstName,
		LastName:        userDto.LastName,
		Email:           userDto.Email,
		HashedSecretKey: string(hashP),
		CreatedAt:       time.Now(),
	}
	userDB, err := u.repo.CreateWallet(&userObject)
	return userDB.ToDtoResponse(), err
}
