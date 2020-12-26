package v1

import (
	dto "WebBackend/pkg/api/core/dto"
	middlewares "WebBackend/pkg/api/core/middlewares"
	models "WebBackend/pkg/api/core/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

type PeopleController struct{}

// GetAll godoc
// @Summary GetAll people
// @Tags People
// @Description get all People
// @version 1.0
// @produce application/json
// @Success 200 {object} middlewares.Success{data=[]models.Person} "success"  成功後返回的值
// @Router /people/ [get]
func (ctrl PeopleController) GetAll(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var body []models.Person
	var test models.Peo
	//db.Model(&test).Association("companies")
	//db.Where("id = ?", 1).First(&test)
	// db.Model(test).Related(&test.Company)

	fmt.Println(test)
	if err := db.Find(&body).Error; err != nil {
		middlewares.ResponseError(c, 1002, "", err)
	} else {
		middlewares.ResponseSuccess(c, "GET all people successful", body)
	}
}

func (ctrl PeopleController) Get(c *gin.Context) {
	id := c.Params.ByName("id")
	var person models.Person
	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("id = ?", id).First(&person).Error; err != nil {
		c.AbortWithStatus(404)
	} else {
		c.JSON(http.StatusOK, person)
	}
}

// Post godoc
// @Summary Create people
// @Tags People
// @Description create People
// @version 1.0
// @Accept  json
// @Produce  json
// @Param polygon body  dto.PeopleDto true "body"
// @produce application/json
// @Success 200 {object} middlewares.Success{data=[]models.Person} "success"  成功後返回的值
// @Router /people/ [post]
func (ctrl PeopleController) Post(c *gin.Context) {
	// @Header 200 {string} Token "qwerty"
	// @Failure 400 {object} httputil.HTTPError
	// @Failure 404 {object} httputil.HTTPError
	// @Failure 500 {object} httputil.HTTPError
	db := c.MustGet("db").(*gorm.DB)
	var body dto.PeopleDto

	if err := c.ShouldBind(&body); err != nil {
		middlewares.ResponseError(c, 1004, "", err)
	} else {
		db.Scopes(models.PeopleTable()).Create(&body)
		c.JSON(200, body)
	}
}

func (ctrl PeopleController) Put(c *gin.Context) {
	var person models.Person
	id := c.Params.ByName("id")
	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("id = ?", id).First(&person).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&person)

	db.Save(&person)
	c.JSON(200, person)
}

func (ctrl PeopleController) Delete(c *gin.Context) {
	id := c.Params.ByName("id")
	db := c.MustGet("db").(*gorm.DB)
	var person models.Person
	d := db.Where("id = ?", id).Delete(&person)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
