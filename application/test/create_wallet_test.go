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
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateWallet(t *testing.T) {

	controller := gomock.NewController(t)
	mockService := service.NewMockWalletService(controller)
	router := gin.Default()
	h := &handler.Handler{WalletService: mockService}
	server.DefineRouter(router, h)

	createUser := &wallet.User{
		Reference: "1",
		FirstName: "King",
		LastName:  "Pharoah",
		Password:  "Pharoah",
		Email:     "Pharoah@egypt.com",
		BVN:       "222",
		Currency:  "NGN",
	}
	marshal, err := json.Marshal(createUser)
	if err != nil {
		return
	}
	var userDB []*wallet.User
	//
	mockService.EXPECT().GetUserByEmail(createUser.Email).Return(userDB, nil)
	mockService.EXPECT().CreateWallet(gomock.Any()).Return(createUser, nil)

	t.Run("Test for creating wallet", func(t *testing.T) {

		req, err := http.NewRequest(http.MethodPost, "/api/v1/createWallet", bytes.NewBuffer(marshal))

		if err != nil {
			log.Fatalf("an error occured:%v", err)

		}

		req.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		router.ServeHTTP(response, req)
		if response.Code != http.StatusCreated {
			t.Errorf("expected %v but got %v", http.StatusCreated, response.Code)
		}

		var responseBodyTwo = `"email":"Pharoah@egypt.com","currency":"NGN","date_of_birth":"","created_at":"0001-01-01T00:00:00Z"`
		if !strings.Contains(response.Body.String(), responseBodyTwo) {
			t.Errorf("Expected body to contain %s", responseBodyTwo)
		}
	})

}
