package v1

import (
	dto "WebBackend/pkg/api/core/dto"
	middlewares "WebBackend/pkg/api/core/middlewares"
	models "WebBackend/pkg/api/core/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"math/rand"
	"time"
)

type CommentsController struct{}

// @Summary get comments
// @Tags Comments
// @Description post comments
// @version 1.0
// @produce application/json
// @Param id path int  true "id"
// @Success 200 {object} middlewares.Success{data=dto.GetCommentsResDto} "success"
// @Failure 400 {object} middlewares.Error "Need ID"
// @Failure 404 {object} middlewares.Error "Not find ID"
// @Router /comments/{id} [get]
func (ctrl CommentsController) Get(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Params.ByName("id")
	var response []dto.GetCommentsResDto
	if err := db.Scopes(models.CommentsTable()).Where("blog_id = ?", id).Find(&response).Error; err != nil {
		middlewares.ResponseError(c, 1002, "", err)
	} else {
		middlewares.ResponseSuccess(c, "GET comments successful", response)
	}
}

// @Summary post comments
// @Tags Comments
// @Description post comments
// @version 1.0
// @produce application/json
// @Accept  json
// @Produce  json
// @Param polygon body dto.PostCommentsReqDto true "body"
// @Success 200 {object} middlewares.Success{data= dto.PostCommentsResDto} "success"
// @Failure 400 {object} middlewares.Error "Need ID"
// @Failure 404 {object} middlewares.Error "Not find ID"
// @Router /comments [post]
func (ctrl CommentsController) Post(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var body dto.PostCommentsReqDto
	if err := c.ShouldBindJSON(&body); err != nil {
		middlewares.ResponseError(c, 1004, "", err)
	} else {
		rand.Seed(time.Now().UnixNano())
		number := rand.Intn(22) + 1 // 1~23
		loc := time.FixedZone("UTC-8", 8*60*60)
		now := time.Now().In(loc)
		body.CreateTime = now
		body.Character = number
		if err := db.Scopes(models.CommentsTable()).Create(&body).Error; err != nil {
			middlewares.ResponseError(c, 1004, "", err)
		}
		middlewares.ResponseSuccess(c, "POST comments successful", "")
	}
}
