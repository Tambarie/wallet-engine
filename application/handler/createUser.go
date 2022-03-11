package handler

import (
	helpers "github.com/Tambarie/wallet-engine/domain/helpers"
	"github.com/Tambarie/wallet-engine/domain/wallet"
	"github.com/Tambarie/wallet-engine/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"time"
)

// CreateWallet Handler to create wallet for user
func (h *Handler) CreateWallet() gin.HandlerFunc {
	return func(context *gin.Context) {

		var user = &wallet.User{}
		hashedPassword, err := helpers.GenerateHashPassword(user.Password)
		if err != nil {
			log.Fatalf("error :%v", err)
		}

		user.Reference = uuid.New().String()
		user.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
		user.HashedSecretKey = string(hashedPassword)

		// Binding the json
		if errs := helpers.Decode(context, &user); errs != nil {
			response.JSON(context, http.StatusInternalServerError, nil, errs, "")
			return
		}

		// Getting user by email
		userDB, err := h.WalletService.GetUserByEmail(user.Email)
		if err != nil {
			log.Fatalf("error :%v", err)
			return
		}

		// Checking to see if a user already exists
		if len(userDB) == 0 {
			userD, err := h.WalletService.CreateWallet(user)

			if err != nil {
				log.Fatalf("error :%v", err)
				return
			}

			response.JSON(context, http.StatusCreated, gin.H{"data": userD}, nil, "User created successfully")
			return
		} else {
			response.JSON(context, http.StatusNotFound, nil, []string{"User email already exists"}, "")
			return
		}
	}
}
