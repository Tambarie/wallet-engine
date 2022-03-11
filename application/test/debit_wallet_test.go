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

func TestDebitWallet(t *testing.T) {
	controller := gomock.NewController(t)
	mockService := service.NewMockWalletService(controller)
	h := &handler.Handler{
		WalletService: mockService,
	}
	router := gin.Default()
	server.DefineRouter(router, h)

	transaction := &wallet.Transaction{
		UserID:               "tambarie",
		TransactionReference: "ta3434",
		Amount:               3000,
		PhoneNumber:          "090232323423",
		Password:             "23jk",
	}

	marshalledTransaction, err := json.Marshal(&transaction)
	if err != nil {
		log.Fatal(err)
	}

	user := &wallet.User{
		Reference:       "2f509a97-98ac-4ee7-baf0-fd01a5d653a0",
		FirstName:       "King",
		LastName:        "Pharoah",
		Email:           "Pharh@egypt.com",
		BVN:             "",
		Currency:        "NGN",
		Password:        "emma",
		HashedSecretKey: "",
		DateOfBirth:     "1945-01-12",
		CreatedAt:       time.Now(),
		IsActive:        true,
	}

	user.IsActive = true
	//
	//marshalledUser, err := json.Marshal(&user)
	//if err != nil{
	//	log.Fatal(err)
	//	return
	//}

	var userDB []*wallet.User

	t.Run("test foe debit_user", func(t *testing.T) {

		mockService.EXPECT().CheckIfPasswordExists(gomock.Any()).Return(userDB, nil).Times(1)

		req, err := http.NewRequest(http.MethodPost, "/api/v1/debitWallet/:user-reference", bytes.NewBuffer(marshalledTransaction))
		if err != nil {
			log.Fatal(err)
			return
		}
		req.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		router.ServeHTTP(response, req)
		assert.Contains(t, response.Body.String(), "Sorry, your account is not active")
	})
}
