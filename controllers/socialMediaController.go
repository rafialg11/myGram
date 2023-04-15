package controllers

import (
	"myGram/database"
	"myGram/helpers"
	"myGram/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// CreateSocialMedia godoc
// @Summary Post detail for a given id
// @Description Post SocialMedia Detail corresponding to the input id
// @Tags socialMedia
// @Accept json
// @Produce json
// @Param models.SocialMedia body models.SocialMedia true "create socmed"
// @Success 200 {object} models.SocialMedia
// @Router /socMed/ [post]
func CreateSocMed(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	User := models.User{}
	SocialMedia := models.SocialMedia{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserID = userID
	SocialMedia.User = &User

	if err := db.First(&User, "id = ?", userID).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Gagal Mendapatkan Data"})
		return
	}

	err := db.Debug().Create(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, SocialMedia)
}

// GetSocialMediaById godoc
// @Summary Get detail for the given id
// @Description Get SocialMedia Detail corresponding to the input id
// @Tags socialMedia
// @Accept json
// @Produce json
// @Param id path int true "Id of the commment"
// @Success 200 {object} models.SocialMedia
// @Router /socMed/{id} [put]
func GetSocialMediaById(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	SocialMedia := models.SocialMedia{}
	User := models.User{}

	socialMediaId, _ := strconv.Atoi(c.Param("socialMediaId"))

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.User = &User

	if err := db.First(&SocialMedia, "id = ?", socialMediaId).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	userID := SocialMedia.UserID

	err := db.First(&User, "id = ?", userID).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, SocialMedia)
}

// UpdateSocialMedia godoc
// @Summary Update SocialMedia identified by the given id
// @Description Update the comment corresponding to the input id
// @Tags socialMedia
// @Accept json
// @Produce json
// @Param id path int true "Id of the commment to be updated"
// @Success 200 {object} models.SocialMedia
// @Router /socMed/{id} [put]
func UpdateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	SocialMedia := models.SocialMedia{}

	socialMediaId, _ := strconv.Atoi(c.Param("socialMediaId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserID = userID
	SocialMedia.ID = uint(socialMediaId)

	err := db.Model(&SocialMedia).Where("id = ?", socialMediaId).Updates(models.SocialMedia{Name: SocialMedia.Name, SocialUrl: SocialMedia.SocialUrl}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, SocialMedia)
}

// DeleteSocialMedia godoc
// @Summary Delete SocialMedia identified by the given id
// @Description Delete the comment corresponding to the input id
// @Tags socialMedia
// @Accept json
// @Produce json
// @Param id path int true "Id of the commment to be deleted"
// @Success 200 {object} models.SocialMedia
// @Router /socMed/{id} [delete]
func DeleteSocialMedia(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	SocialMedia := models.SocialMedia{}

	socialMediaId, _ := strconv.Atoi(c.Param("socialMediaId"))
	SocialMedia.ID = uint(socialMediaId)
	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	err := db.Where("id = ?", socialMediaId).Delete(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, "Deleted")
}

// GetAllSocialMedia godoc
// @Summary Get Details
// @Description Get details of All SocialMedia
// @Tags socialMedia
// @Accept json
// @Produce json
// @Success 200 {object} models.SocialMedia
// @Router /socMed/ [get]
func GetAllSocialMedia(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	var SocialMedia []models.SocialMedia

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	err := db.Find(&SocialMedia).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, SocialMedia)
}
