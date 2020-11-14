package v1

import (
	dto "WebBackend/pkg/api/core/dto"
	middlewares "WebBackend/pkg/api/core/middlewares"
	models "WebBackend/pkg/api/core/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type BlogController struct{}

// @Summary GetAll blog
// @Tags Blog
// @Description get all Blog
// @version 1.0
// @produce application/json
// @Success 200 {object} middlewares.Success{data=[]dto.GetAllBlogResDto} "success"
// @Router /blog/ [get]
func (ctrl BlogController) GetAll(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var response []dto.GetAllBlogResDto
	if err := db.Scopes(models.BlogTable()).Find(&response).Error; err != nil {
		middlewares.ResponseError(c, 1002, "", err)
	} else {
		middlewares.ResponseSuccess(c, "GET all blogs successful", response)
	}
}
