package handler

import (
	"github.com/Tambarie/wallet-engine/domain/wallet"
	"github.com/Tambarie/wallet-engine/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) ActivateWallet() gin.HandlerFunc {
	return func(context *gin.Context) {
		userReference := context.Param("user-reference")
		activate := context.Query("activate")

		user := &wallet.User{}
		var message string
		var status bool
		if activate == "true" {
			message = "activated successfully"
			status = true
		} else {
			message = "deactivate successfully"
			status = false

		}
		user.ActivateDeactivateWallet(status)
		_, err := h.WalletService.ChangeUserStatus(user.IsActive, userReference)
		if err != nil {
			response.JSON(context, http.StatusNotFound, nil, []string{"could not post to user account "}, "")
			return
		}

		response.JSON(context, http.StatusCreated, gin.H{
			"message": message,
		}, nil, "")
	}
}
