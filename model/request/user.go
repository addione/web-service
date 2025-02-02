package request

import (
	"github.com/addione/web-service/model"
	"github.com/gin-gonic/gin"
)

type CreateUserParams struct {
	Name        string `binding:"required"`
	Email       string `json:"email" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Address     string
	Details     model.Details
	Password    string `binding:"required"`
}

func ValidateCreateUser(ctx *gin.Context) (*CreateUserParams, error) {
	var cup CreateUserParams

	err := ctx.ShouldBindJSON(&cup)
	if err != nil {
		return nil, err
	}
	return &cup, nil
}
