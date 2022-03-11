package handler

import (
	"github.com/Tambarie/wallet-engine/domain/helpers"
	"github.com/Tambarie/wallet-engine/domain/wallet"
	"github.com/Tambarie/wallet-engine/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

// DebitWallet Handler to DebitWallet of the user
func (h *Handler) DebitWallet() gin.HandlerFunc {
	return func(context *gin.Context) {

		//Getting the user reference
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
			response.JSON(context, http.StatusNotFound, nil, []string{"sorry you can't debit less than N1000.00"}, "")
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
			response.JSON(context, http.StatusNotFound, nil, []string{"Sorry, your account is not active"}, "")
			return
		}

		// querying database for the balance
		t, err := h.WalletService.GetAccountBalance(userID)

		var currentBalance float64
		for _, user := range t {
			currentBalance = user.Balance

		}

		account.Balance = currentBalance

		//checking if the account balance is less than N0:00
		if account.Balance <= 0 {
			response.JSON(context, http.StatusNotFound, nil, []string{"Sorry, your account is insufficient for this transaction"}, "")
			return
		}

		// check if the debit amount is greater than the balance
		if account.Balance < transaction.Amount {
			response.JSON(context, http.StatusNotFound, nil, []string{"Sorry, your account is insufficient for this transaction"}, "")
			return
		}

		// method handles the debit of the wallet
		account.DebitUserWallet(transaction.Amount, transaction.UserID)

		// Handles saving of the transaction of the wallet
		userTransaction, err := h.WalletService.SaveTransaction(transaction)
		if err != nil {
			response.JSON(context, http.StatusNotFound, nil, []string{"could not fetch user reference "}, "")
			return
		}

		// Handles posting of the money to the user's account
		currentAccount, err := h.WalletService.PostToAccount(account)
		if err != nil {
			response.JSON(context, http.StatusNotFound, nil, []string{"could not debit to user account "}, "")
			return
		}

		// json response
		response.JSON(context, http.StatusCreated, gin.H{
			"transaction reference": userTransaction.TransactionReference,
			"amount credited":       userTransaction.Amount,
			"account balance":       currentAccount.Balance,
		},
			nil,
			"account debited successfully")
	}
}
