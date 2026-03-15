package routes

import (
	"auto-pharmacy/internal/controllers"
	"os"

	echojwt "github.com/labstack/echo-jwt/v5"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func RegisterWebRoutes(e *echo.Group, uc *controllers.UserController) {
	e.Use(middleware.CORS("https://example.com", "https://subdomain.example.com"))
	e.Use(echojwt.JWT([]byte(os.Getenv("SECRET_KEY"))))
	medicineRoutes(e.Group("/medicines"))
	supplyRoutes(e.Group("/supplies"))
	tagRoutes(e.Group("/tags"))
	userRoutes(e.Group("/users"), uc)
}

func medicineRoutes(e *echo.Group) {
	e.GET("", controllers.MedicineIndex)
	e.POST("", controllers.MedicineSet)
	e.GET("/:medicine", controllers.MedicineGet)
	e.PUT("/:medicine", controllers.MedicineUpdate)
	e.DELETE("/:medicine", controllers.MedicineDelete)
	e.PUT("/:medicine/restore", controllers.MedicineRestore)
	e.DELETE("/:medicine/force", controllers.MedicineForceDelete)
	e.GET("/:medicine/release", controllers.MedicineRelease)
	e.PUT("/:medicine/associate/:tag", controllers.MedicineAssociateTag)
	e.POST("/delete/mass", controllers.MedicineMassDelete)
}

func supplyRoutes(e *echo.Group) {
	e.GET("", controllers.SupplyIndex)
	e.POST("", controllers.SupplySet)
	e.GET("/:supply", controllers.SupplyGet)
	e.PUT("/:supply", controllers.SupplyUpdate)
	e.DELETE("/:supply", controllers.SupplyDelete)
	e.PUT("/:supply/restore", controllers.SupplyRestore)
	e.DELETE("/:supply/force", controllers.SupplyForceDelete)
	e.POST("/delete/mass", controllers.SupplyMassDelete)
}

func tagRoutes(e *echo.Group) {
	e.GET("", controllers.TagIndex)
	e.POST("", controllers.TagSet)
	e.GET("/:tag", controllers.TagGet)
	e.PUT("/:tag", controllers.TagUpdate)
	e.DELETE("/:tag", controllers.TagDelete)
	e.PUT("/:tag/restore", controllers.TagRestore)
	e.DELETE("/:tag/force", controllers.TagForceDelete)
	// e.PUT("/:tag/associate/:medicine", controllers.MedicineUpdate)
}

func userRoutes(e *echo.Group, uc *controllers.UserController) {
	e.GET("", uc.Index)
	e.POST("", uc.Set)
	e.GET("/:user", uc.Get)
	e.PUT("/:user", uc.Update)
	e.DELETE("/:user", uc.Delete)
	e.PUT("/:user/restore", uc.Restore)
	e.DELETE("/:user/force", uc.ForceDelete)
	e.POST("/delete/mass", uc.MassDelete)
}
