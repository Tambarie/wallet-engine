package wallet

type Transaction struct {
	UserID               string
	TransactionReference string
	Amount               float64
	PhoneNumber          string
	Password             string
}

type Wallet struct {
	UserID    string
	Balance   float64
	Reference string
}

func (w *Wallet) CreditUserWallet(money float64, userID string) {
	w.UserID = userID
	w.Balance += money
}

func (w *Wallet) DebitUserWallet(money float64, userID string) {
	w.UserID = userID
	w.Balance -= money
}
