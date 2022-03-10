package handler

import (
	"github.com/Tambarie/wallet-engine/domain/helpers"
	"github.com/Tambarie/wallet-engine/domain/wallet"
	"github.com/Tambarie/wallet-engine/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
)

func (h *Handler) DebitWallet() gin.HandlerFunc {
	return func(context *gin.Context) {

		userID := context.Param("user-reference")

		transaction := &wallet.Transaction{}
		transaction.UserID = userID
		transaction.TransactionReference = uuid.New().String()

		if err := helpers.Decode(context, &transaction); err != nil {
			response.JSON(context, http.StatusNotFound, nil, []string{"transaction could not be decoded"}, "")
			return
		}

		if transaction.Amount < 1000 {
			response.JSON(context, http.StatusNotFound, nil, []string{"sorry you can't debit less than N1000.00"}, "")
			return
		}
		userDB, err := h.WalletService.CheckIfPasswordExists(userID)
		if err != nil {
			response.JSON(context, http.StatusNotFound, nil, []string{"could not fetch user reference "}, "")
			return
		}

		var hashedPassword string
		for _, user := range userDB {
			hashedPassword = user.HashedSecretKey
		}

		if correct := helpers.CheckPasswordHash(transaction.Password, []byte(hashedPassword)); correct {
			response.JSON(context, http.StatusNotFound, nil, []string{"Invalid password"}, "")
			return
		}

		account := &wallet.Account{}

		// query db for balance
		t, err := h.WalletService.GetAccountBalance(userID)

		var currentBalance float64
		for _, user := range t {
			currentBalance = user.Balance

		}

		account.Balance = currentBalance
		log.Println(account.Balance)
		if account.Balance < 0 {
			response.JSON(context, http.StatusNotFound, nil, []string{"Sorry, your account is insufficient for this transaction"}, "")
			return
		}

		account.DebitUserWallet(transaction.Amount, transaction.UserID)

		userTransaction, err := h.WalletService.SaveTransaction(transaction)
		if err != nil {
			response.JSON(context, http.StatusNotFound, nil, []string{"could not fetch user reference "}, "")
			return
		}

		currentAccount, err := h.WalletService.PostToAccount(account)
		if err != nil {
			response.JSON(context, http.StatusNotFound, nil, []string{"could not debit to user account "}, "")
			return
		}

		response.JSON(context, http.StatusCreated, gin.H{
			"transaction reference": userTransaction.TransactionReference,
			"amount credited":       userTransaction.Amount,
			"account balance":       currentAccount.Balance,
		},
			nil,
			"account debited successfully")
	}
}
