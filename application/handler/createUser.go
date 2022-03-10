package handler

import (
	"fmt"
	helpers "github.com/Tambarie/wallet-engine/domain/helpers"
	"github.com/Tambarie/wallet-engine/domain/wallet"
	"github.com/Tambarie/wallet-engine/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"time"
)

func (h *Handler) CreateUser() gin.HandlerFunc {
	return func(context *gin.Context) {
		var user = &wallet.User{}
		hashedPassword, err := helpers.GenerateHashPassword(user.Password)
		if err != nil {
			fmt.Println(err)
		}

		user.Reference = uuid.New().String()
		user.CreatedAt = time.Now().UTC()
		user.HashedSecretKey = string(hashedPassword)
		//log.Println(user.Password)
		if errs := helpers.Decode(context, &user); errs != nil {
			fmt.Println(errs)
			response.JSON(context, http.StatusInternalServerError, nil, errs, "")
			return
		}

		userDB, err := h.WalletService.GetUserByEmail(user.Email)
		if err != nil {
			log.Println(err)
			return
		}
		if len(userDB) == 0 {
			userD, err := h.WalletService.CreateWallet(user)

			if err != nil {
				log.Println(err)
				return
			}

			response.JSON(context, http.StatusCreated, gin.H{"data": userD}, nil, "User created successfully")
			return
		} else {
			log.Println("here")
			response.JSON(context, http.StatusNotFound, nil, []string{"User email already exists"}, "")
			return
		}
	}
}
