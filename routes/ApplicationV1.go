package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/plasticube/microservices-inspect/controllers/inspection"
	"github.com/plasticube/microservices-inspect/controllers/medicine"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/plasticube/microservices-inspect/docs"
)

// @title Boilerplate Golang
// @version 1.0
// @description Documentation's Boilerplate Golang
// @termsOfService http://swagger.io/terms/

// @contact.name Alejandro Gabriel Guerrero
// @contact.url http://github.com/gbrayhan
// @contact.email gbrayhan@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /v1
func ApplicationV1Router(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		// Documentation Swagger
		{
			v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		}
		// Medicines
		v1Medicines := v1.Group("/medicine")
		{
			//v1Medicines.Use(middlewares.CheckAuth)
			v1Medicines.POST("/", medicine.NewMedicine)
			v1Medicines.GET("/:id", medicine.GetMedicinesByID)
			v1Medicines.GET("/", medicine.GetAllMedicines)
			v1Medicines.PUT("/:id", medicine.UpdateMedicine)
			//v1Medicines.DELETE("/:id", medicine.DeleteMedicine)
		}
		// Inspection
		v1Inspection := v1.Group("/inspection")
		{
			//v1Inspection.Use(middlewares.CheckAuth)
			v1Inspection.POST("/", inspection.NewInspection)
			v1Inspection.GET("/:id", inspection.GetInspectionById)
			v1Inspection.GET("/", inspection.GetAllInspection)
			v1Inspection.GET("/page", inspection.GetAllInspectionByPage)
			v1Inspection.PUT("/:id", inspection.UpdateInspection)
			//v1Inspection.DELETE("/:id", inspection.DeleteInspection)
		}
	}
}
