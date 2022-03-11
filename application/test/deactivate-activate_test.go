package test

import (
	"bytes"
	"github.com/Tambarie/wallet-engine/application/handler"
	"github.com/Tambarie/wallet-engine/application/server"
	"github.com/Tambarie/wallet-engine/domain/service"
	"github.com/Tambarie/wallet-engine/domain/wallet"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go/types"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestDeactivateActivate(t *testing.T) {
	controller := gomock.NewController(t)
	mockService := service.NewMockWalletService(controller)
	h := &handler.Handler{
		WalletService: mockService,
	}
	router := gin.Default()
	server.DefineRouter(router, h)

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

	body := []byte("status changed")

	t.Run("testing for deactivate/activate wallet", func(t *testing.T) {
		mockService.EXPECT().ChangeUserStatus(gomock.Any(), user.Reference).Return(types.Interface{}, nil)

		req, err := http.NewRequest(http.MethodPut, "/api/v1/activate-deactivate/2f509a97-98ac-4ee7-baf0-fd01a5d653a0", bytes.NewBuffer(body))
		if err != nil {
			log.Fatal(err)
			return
		}
		response := httptest.NewRecorder()
		router.ServeHTTP(response, req)
		assert.Contains(t, response.Body.String(), "deactivate successfully")
		assert.Equal(t, http.StatusCreated, response.Code)
	})

}
