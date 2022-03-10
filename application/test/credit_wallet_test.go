package test

import (
	"bytes"
	"encoding/json"
	"github.com/Tambarie/wallet-engine/application/handler"
	"github.com/Tambarie/wallet-engine/application/server"
	"github.com/Tambarie/wallet-engine/domain/service"
	"github.com/Tambarie/wallet-engine/domain/wallet"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCreditWallet(t *testing.T) {
	controller := gomock.NewController(t)
	mockService := service.NewMockWalletService(controller)
	router := gin.Default()
	h := &handler.Handler{
		WalletService: mockService,
	}
	server.DefineRouter(router, h)

	transaction := &wallet.Transaction{
		UserID:               "ahhsh",
		TransactionReference: "jhasjh",
		Amount:               10000,
		PhoneNumber:          "08037700350",
		Password:             "jbkbkxbkx",
	}

	user := &wallet.User{
		Reference: "1",
		FirstName: "King",
		LastName:  "Pharoah",
		Password:  "Pharoah",
		Email:     "Pharoah@egypt.com",
		BVN:       "222",
		Currency:  "NGN",
	}

	marshalledTransaction, err := json.Marshal(transaction)
	if err != nil {
		log.Println(err)
		return
	}
	var userDB []*wallet.User
	mockService.EXPECT().CheckIfPasswordExists(gomock.Any()).Return(userDB, nil)
	t.Run("testing if password exists", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPost, "/api/v1/creditWallet/:user-reference", bytes.NewBuffer(marshalledTransaction))
		if err != nil {
			log.Fatalf("an error occured: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		router.ServeHTTP(response, req)

		assert.Contains(t, response.Body.String(), "sorry, please activate your account")
	})

	t.Run("test for account balance", func(t *testing.T) {
		var walletDB []*wallet.Wallet
		mockService.EXPECT().GetAccountBalance(gomock.Any()).Return(walletDB)

	})

}
