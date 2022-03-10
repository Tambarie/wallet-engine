package walletRepo

import "github.com/Tambarie/wallet-engine/domain/wallet"

type Repository interface {
	CreateWallet(u *wallet.User) (*wallet.User, error)
	GetUserByEmail(email string) ([]*wallet.User, error)
	CheckIfPasswordExists(userReference string) ([]*wallet.User, error)
	PostToAccount(u *wallet.Account) (*wallet.Account, error)
	SaveTransaction(t *wallet.Transaction) (*wallet.Transaction, error)
	GetAccountBalance(userID string) ([]*wallet.Account, error)
}
