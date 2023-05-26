package routes

import (
	"go_bedu/controllers"
	m "go_bedu/middlewares"
	"go_bedu/repositories"
	"go_bedu/usecase"
	"go_bedu/utils"

	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	mid "github.com/labstack/echo/middleware"
)

func NewRoute(e *echo.Echo, db *gorm.DB) {
	m.Log(e)
	e.Pre(mid.RemoveTrailingSlash())

	adminRepository := repositories.NewAdminRepository(db)
	adminUsecase := usecase.NewAdminUsecase(adminRepository)
	adminController := controllers.NewAdminController(adminUsecase, adminRepository)

	articleRepository := repositories.NewArticleRepository(db)
	articleUsecase := usecase.NewArticleUsecase(articleRepository)
	articleController := controllers.NewArticleController(articleUsecase)

	// Middleware untuk mengatur CORS
	e.Use(mid.CORSWithConfig(mid.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAuthorization, // Menambahkan header Authorization
		},
	}))

	e.Use(m.AllowCORS)

	cv := &utils.CustomValidator{Validators: validator.New()}
	e.Validator = cv

	e.POST("/register", adminController.RegisterAdminController)
	e.POST("/login", adminController.LoginAdminController)

	// Admin Only
	admin := e.Group("/admin")
	admin.GET("", adminController.GetAdminsController, m.VerifyToken)
	admin.GET("/profile", adminController.GetAdminByIdController, m.VerifyToken)
	admin.PUT("", adminController.UpdateAdminController, m.VerifyToken)
	admin.DELETE("", adminController.DeleteAdminController, m.VerifyToken)

	// Article Routes

	admin.GET("/article", articleController.GetAllArticles, m.VerifyToken)
	admin.GET("/article/:id", articleController.GetArticleById, m.VerifyToken)
	admin.POST("/article", articleController.CreateArticle, m.VerifyToken)
	admin.PUT("/article/:id", articleController.UpdateArticle, m.VerifyToken)
	admin.DELETE("/article/:id", articleController.DeleteArticle, m.VerifyToken)
}
