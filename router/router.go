package router

import (
	"myGram/controllers"
	"myGram/middleware"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
	}
	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middleware.Authentication())
		photoRouter.POST("/", controllers.CreatePhoto)
		photoRouter.GET("/:photoId", controllers.GetPhotoById)
		photoRouter.PUT("/:photoId", middleware.PhotoAuthorization(), controllers.UpdatePhoto)
		photoRouter.DELETE("/:photoId", middleware.PhotoAuthorization(), controllers.DeletePhoto)
		photoRouter.GET("/", controllers.GetAllPhoto)
	}
	socMedRouter := r.Group("/socMed")
	{
		socMedRouter.Use(middleware.Authentication())
		socMedRouter.POST("/", controllers.CreateSocMed)
		socMedRouter.GET("/:socialMediaId", controllers.GetSocialMediaById)
		socMedRouter.PUT("/:socialMediaId", middleware.SocMedAuthorization(), controllers.UpdateSocialMedia)
		socMedRouter.DELETE("/:socialMediaId", middleware.SocMedAuthorization(), controllers.DeleteSocialMedia)
		socMedRouter.GET("/", controllers.GetAllSocialMedia)
	}

	return r
}
