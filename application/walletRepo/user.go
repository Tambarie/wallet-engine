package walletRepo

import "github.com/Tambarie/wallet-engine/domain/wallet"

// Repository interface
type Repository interface {
	CreateWallet(u *wallet.User) (*wallet.User, error)
	GetUserByEmail(email string) ([]*wallet.User, error)
	CheckIfPasswordExists(userReference string) ([]*wallet.User, error)
	PostToAccount(u *wallet.Wallet) (*wallet.Wallet, error)
	SaveTransaction(t *wallet.Transaction) (*wallet.Transaction, error)
	GetAccountBalance(userID string) ([]*wallet.Wallet, error)
	ChangeUserStatus(isActive bool, userReference string) (interface{}, error)
}
