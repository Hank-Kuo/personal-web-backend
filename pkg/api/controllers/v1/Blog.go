package v1

import (
	dto "github.com/Hank-Kuo/personal-web-backend/pkg/api/core/dto"
	libs "github.com/Hank-Kuo/personal-web-backend/pkg/api/core/lib"
	middlewares "github.com/Hank-Kuo/personal-web-backend/pkg/api/core/middlewares"
	models "github.com/Hank-Kuo/personal-web-backend/pkg/api/core/models"

	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"time"
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

// @Summary Get blog
// @Tags Blog
// @Description get Blog
// @version 1.0
// @produce application/json
// @Param id path int  true "id"
// @Success 200 {object} middlewares.Success{data=dto.GetBlogHTMLResDto} "success"
// @Failure 400 {object} middlewares.Error "Need ID"
// @Failure 404 {object} middlewares.Error "Not find ID"
// @Router /blog/{id} [get]
func (ctrl BlogController) Get(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Params.ByName("id")
	var r dto.GetBlogResDto
	if err := db.Scopes(models.BlogTable()).Where("id = ?", id).Find(&r).Error; err != nil {
		middlewares.ResponseError(c, 1002, "", err)
	} else {
		if html, err := libs.Crawler(r.Link); err != nil {
			middlewares.ResponseError(c, 1002, "", err)
		} else {
			response := dto.SerializeGetBlog(&r, html)
			middlewares.ResponseSuccess(c, "GET blogs successful", response)
		}

	}
}

// @Summary Post blog
// @Tags Blog
// @Description post Blog
// @version 1.0
// @produce application/json
// @Success 200 {object} middlewares.Success{data=""} "success"
// @Failure 400 {object} middlewares.Error "Need ID"
// @Failure 404 {object} middlewares.Error "Not find ID"
// @Router /blog [post]
func (ctrl BlogController) Post(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var body dto.PostBlogReqDto
	if err := c.ShouldBindJSON(&body); err != nil {
		middlewares.ResponseError(c, 1004, "", err)
	} else {
		isValidUrl := libs.IsValidUrl(body.ImgLink) == true && libs.IsValidUrl(body.Link) == true

		if isValidUrl == false {
			err := errors.New("Invalid URL")
			middlewares.ResponseError(c, 1004, "Invalid URL", err)
		}

		loc := time.FixedZone("UTC-8", 8*60*60)
		now := time.Now().In(loc)
		r := dto.SerializePostBlog(&body, now)
		if err := db.Scopes(models.BlogTable()).Create(&r).Error; err != nil {
			middlewares.ResponseError(c, 1003, "Create Blog fail", err)
		} else {
			emoji := &dto.PostBlogEmojiDto{BlogID: r.ID}
			if err := db.Scopes(models.EmojiTable()).Create(&emoji).Error; err != nil {
				middlewares.ResponseError(c, 1003, "Create Emoji fail", err)
			} else {
				middlewares.ResponseSuccess(c, "GET blogs successful", "")
			}
		}
	}
}

func (ctrl BlogController) Put(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Params.ByName("id")
	var origin dto.GetAllBlogResDto
	var body dto.PutBlogVisitorDto

	if err := c.ShouldBindJSON(&body); err != nil {
		middlewares.ResponseError(c, 1004, "", err)
		return
	}
	if err := db.Scopes(models.BlogTable()).Where("id = ?", id).First(&origin).Error; err != nil {
		middlewares.ResponseError(c, 1002, "", err)
	} else {
		origin.Visitor = origin.Visitor + 1
		if err := db.Model(models.Blog{}).Where("id = ?", id).Updates(&origin).Error; err != nil {
			middlewares.ResponseError(c, 1004, "", err)
			return
		}
		middlewares.ResponseSuccess(c, "PUT Visitor successful", "")
	}

}
