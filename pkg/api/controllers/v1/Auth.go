package v1

import (
	dto "github.com/Hank-Kuo/personal-web-backend/pkg/api/core/dto"
	libs "github.com/Hank-Kuo/personal-web-backend/pkg/api/core/lib"
	middlewares "github.com/Hank-Kuo/personal-web-backend/pkg/api/core/middlewares"
	models "github.com/Hank-Kuo/personal-web-backend/pkg/api/core/models"

	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type AuthController struct{}

// Login godoc
// @Summary Login
// @Tags Auth
// @Description Login
// @version 1.0
// @Param polygon body  dto.LoginReqDto true "body"
// @produce application/json
// @Success 200 {object} middlewares.Success{data=[]dto.LoginResDto} "success"
// @Router /auth/login [post]
func (ctrl AuthController) Login(c *gin.Context) {
	var body dto.LoginReqDto
	var response dto.LoginResDto
	db := c.MustGet("db").(*gorm.DB)

	// format error
	if err := c.ShouldBindJSON(&body); err != nil {
		middlewares.ResponseError(c, 1004, "", err)
		return
	}

	// incorrect account or password
	if err := db.Scopes(models.UserTable()).
		Where("account = ?", body.Account).
		First(&response).Error; err != nil {
		middlewares.ResponseError(c, 1006, "Invalid account", err)
		return
	}

	if libs.CheckPasswordHash(body.Password, response.Password) == false {
		err := errors.New("Incorrect password")
		middlewares.ResponseError(c, 1006, "Incorrect password", err)
		return
	}

	token, err := middlewares.Jwt(response.Account, response.Role)
	if err != nil {
		middlewares.ResponseError(c, 1009, "", err)
		return
	}

	response.Token = token
	r := dto.SerializeLoginDto(&response)
	middlewares.ResponseSuccess(c, "login successful", r)
}

// Register godoc
// @Summary Register
// @Tags Auth
// @Description Register
// @version 1.0
// @Param polygon body dto.RegisterReqDto true "body"
// @produce application/json
// @Success 200 {object} middlewares.Success{data=""} "success"
// @Router /auth/register [post]
func (ctrl AuthController) Register(c *gin.Context) {
	var body dto.RegisterReqDto
	db := c.MustGet("db").(*gorm.DB)

	// format error
	if err := c.ShouldBindJSON(&body); err != nil {
		middlewares.ResponseError(c, 1004, "", err)
		return
	} else {
		// check user exist
		result := db.Scopes(models.UserTable()).Where("account = ?", body.Account).First(&body)
		if result.RowsAffected > 0 {
			middlewares.ResponseError(c, 1006, "Already exist account", err)
			return
		} else {
			// create user
			body.Role = "Admin"
			if hashPass, err := libs.HashPassword(body.Password); err != nil {
				middlewares.ResponseError(c, 1009, "", err)
				return
			} else {
				body.Password = hashPass
				if err := db.Scopes(models.UserTable()).Create(&body).Error; err != nil {
					middlewares.ResponseError(c, 1003, "Create user fail", err)
					return
				} else {
					middlewares.ResponseSuccess(c, "register successful", "")
				}
			}

		}

	}

}
