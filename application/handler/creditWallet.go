package handler

import (
	"github.com/Tambarie/wallet-engine/domain/helpers"
	"github.com/Tambarie/wallet-engine/domain/wallet"
	"github.com/Tambarie/wallet-engine/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

// CreditWallet Handler to credit wallet of the user
func (h *Handler) CreditWallet() gin.HandlerFunc {
	return func(context *gin.Context) {

		// Getting the user reference
		userID := context.Param("user-reference")
		transaction := &wallet.Transaction{}
		transaction.UserID = userID
		transaction.TransactionReference = uuid.New().String()

		// Binding the json
		if err := helpers.Decode(context, &transaction); err != nil {
			response.JSON(context, http.StatusNotFound, nil, []string{"transaction could not be decoded"}, "")
			return
		}

		// Check if the transaction amount is less than 1000
		if transaction.Amount < 1000 {
			response.JSON(context, http.StatusNotFound, nil, []string{"sorry you can't deposit less than N1000.00"}, "")
			return
		}

		// Checking for the authenticity of the password
		userDB, err := h.WalletService.CheckIfPasswordExists(userID)
		if err != nil {
			response.JSON(context, http.StatusNotFound, nil, []string{"could not fetch user reference "}, "")
			return
		}

		var hashedPassword string
		var checkIfUserActive bool
		for _, user := range userDB {
			hashedPassword = user.HashedSecretKey
			checkIfUserActive = user.IsActive
		}

		// Confirming the password provided by the user
		if correct := helpers.CheckPasswordHash(transaction.Password, []byte(hashedPassword)); correct {
			response.JSON(context, http.StatusNotFound, nil, []string{"Invalid password"}, "")
			return
		}

		account := &wallet.Wallet{}
		wUser := &wallet.User{}
		wUser.IsActive = checkIfUserActive

		// Checking if the user is active
		if wUser.IsActive == false {
			response.JSON(context, http.StatusNotFound, nil, []string{"sorry, please activate your account"}, "")
			return
		}

		// querying database for the balance
		t, err := h.WalletService.GetAccountBalance(userID)

		var currentBalance float64
		for _, user := range t {
			currentBalance = user.Balance

		}
		account.Balance = currentBalance

		// Handles the crediting of the wallet
		account.CreditUserWallet(transaction.Amount, transaction.UserID)

		// Handles saving of the transaction of the wallet
		userTransaction, err := h.WalletService.SaveTransaction(transaction)
		if err != nil {
			response.JSON(context, http.StatusNotFound, nil, []string{"could not fetch user reference "}, "")
			return
		}

		// Handles posting of the money to the user's account
		currentAccount, err := h.WalletService.PostToAccount(account)
		if err != nil {
			response.JSON(context, http.StatusNotFound, nil, []string{"could not post to user account "}, "")
			return
		}

		// json response
		response.JSON(context, http.StatusCreated, gin.H{
			"transaction reference": userTransaction.TransactionReference,
			"amount credited":       userTransaction.Amount,
			"account balance":       currentAccount.Balance,
		},
			nil,
			"account credit successfully")
	}
}
