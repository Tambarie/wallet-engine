package helpers

import (
	"github.com/Tambarie/wallet-engine/infrastructure/servererrors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Decode(c *gin.Context, v interface{}) []string {
	if err := c.ShouldBindJSON(v); err != nil {
		var errs []string
		verr, ok := err.(validator.ValidationErrors)
		if ok {
			for _, fieldErr := range verr {
				errs = append(errs, servererrors.NewFieldError(fieldErr).String())
			}
		} else {
			errs = append(errs, "internal server error")
		}
		return errs
	}
	return nil
}
