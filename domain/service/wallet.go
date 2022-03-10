package service

import (
	"github.com/Tambarie/wallet-engine/application/walletRepo"
	"github.com/Tambarie/wallet-engine/domain/wallet"
)

type WalletService interface {
	CreateWallet(userDto *wallet.User) (*wallet.User, error)
	CheckUserExists(email string) ([]*wallet.User, error)
	CheckIfPasswordExists(userReference string) ([]*wallet.User, error)
	SaveTransaction(t *wallet.Transaction) (*wallet.Transaction, error)
	PostToAccount(a *wallet.Account) (*wallet.Account, error)
	GetAccountBalance(userID string) ([]*wallet.Account, error)
}

type DefaultWalletService struct {
	repo walletRepo.Repository
}

func NewWalletService(repo walletRepo.Repository) *DefaultWalletService {
	return &DefaultWalletService{
		repo: repo,
	}
}

func (u *DefaultWalletService) CreateWallet(userDto *wallet.User) (*wallet.User, error) {
	return u.repo.CreateWallet(userDto)
}

func (u *DefaultWalletService) CheckUserExists(email string) ([]*wallet.User, error) {
	return u.repo.GetUserByEmail(email)
}

func (u *DefaultWalletService) CheckIfPasswordExists(userReference string) ([]*wallet.User, error) {
	return u.repo.CheckIfPasswordExists(userReference)
}

func (u *DefaultWalletService) PostToAccount(a *wallet.Account) (*wallet.Account, error) {
	return u.repo.PostToAccount(a)
}

func (u *DefaultWalletService) SaveTransaction(t *wallet.Transaction) (*wallet.Transaction, error) {
	return u.repo.SaveTransaction(t)
}

func (u *DefaultWalletService) GetAccountBalance(userID string) ([]*wallet.Account, error) {
	return u.repo.GetAccountBalance(userID)
}
