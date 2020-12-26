package v1

import (
	dto "WebBackend/pkg/api/core/dto"
	middlewares "WebBackend/pkg/api/core/middlewares"
	models "WebBackend/pkg/api/core/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type EmojiController struct{}

// @Summary get emoji
// @Tags Emoji
// @Description get emoji
// @version 1.0
// @produce application/json
// @Param id path int  true "id"
// @Success 200 {object} middlewares.Success{data=dto.GetEmojResDto} "success"
// @Failure 400 {object} middlewares.Error "Need ID"
// @Failure 404 {object} middlewares.Error "Not find ID"
// @Router /emoji/{id} [get]
func (ctrl EmojiController) Get(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Params.ByName("id")
	var response dto.GetEmojResDto
	if err := db.Scopes(models.EmojiTable()).Where("blog_id = ?", id).Find(&response).Error; err != nil {
		middlewares.ResponseError(c, 1002, "", err)
	} else {
		middlewares.ResponseSuccess(c, "GET emoji successful", response)
	}
}

// @Summary put emoji
// @Tags Emoji
// @Description put emoji
// @version 1.0
// @produce application/json
// @Accept  json
// @Produce  json
// @Param polygon body  dto.PutEmojiReqDto true "body"
// @Param id path int  true "id"
// @Success 200 {object} middlewares.Success{data=dto.PutEmojiResDto} "success"
// @Failure 400 {object} middlewares.Error "Need ID"
// @Failure 404 {object} middlewares.Error "Not find ID"
// @Router /emoji/{id} [put]
func (ctrl EmojiController) Put(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Params.ByName("id")

	var origin dto.PutEmojiReqDto
	var body dto.PutEmojiReqDto

	if err := c.ShouldBindJSON(&body); err != nil {
		middlewares.ResponseError(c, 1004, "", err)
		return
	}
	if err := db.Scopes(models.EmojiTable()).Where("blog_id = ?", id).First(&origin).Error; err != nil {
		middlewares.ResponseError(c, 1002, "", err)
	} else {
		if err := db.Model(models.Emoji{}).Where("blog_id = ?", id).Updates(&body).Error; err != nil {
			middlewares.ResponseError(c, 1004, "", err)
			return
		}
		middlewares.ResponseSuccess(c, "PUT emoji successful", "")
	}

}
