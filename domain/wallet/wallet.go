package wallet

type Transaction struct {
	UserID               string
	TransactionReference string
	Amount               float64
	PhoneNumber          string
	Password             string
}

type Account struct {
	UserID  string
	Balance float64
}

func (a *Account) CreditUserWallet(money float64, userID string) {
	a.UserID = userID
	a.Balance += money
}

func (a *Account) DebitUserWallet(money float64, userID string) {
	a.UserID = userID
	a.Balance -= money
}
