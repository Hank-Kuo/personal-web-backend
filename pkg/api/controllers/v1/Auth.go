package v1

import (
	dto "WebBackend/pkg/api/core/dto"
	middlewares "WebBackend/pkg/api/core/middlewares"
	models "WebBackend/pkg/api/core/models"
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
// @Router /auth/ [post]
func (ctrl AuthController) Login(c *gin.Context) {
	var body dto.LoginReqDto
	var response dto.LoginResDto
	db := c.MustGet("db").(*gorm.DB)

	// format error
	if err := c.ShouldBindJSON(&body); err != nil {
		middlewares.ResponseError(c, 1004, "", err)
	}

	// incorrect account or password
	if err := db.Scopes(models.UserTable()).
		Where("username = ? AND password >= ?", body.Account, body.Password).
		First(&response).Error; err != nil {
		middlewares.ResponseError(c, 1006, "Incorrect account or password", err)
	}

	token, err := middlewares.Jwt(body.Account, "Member")
	if err != nil {
		middlewares.ResponseError(c, 1009, "", err)
	}
	response.Role = "Member"
	response.Token = token
	middlewares.ResponseSuccess(c, "login successful", response)
}

func (ctrl AuthController) Register(c *gin.Context) {

}
