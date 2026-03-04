package routes

import (
	"internal-office-backend/internal/handler"
	"internal-office-backend/internal/repository"
	"internal-office-backend/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterDepartmentRoutes(rg *gin.RouterGroup, db *gorm.DB) {

	deptRepo := repository.NewDepartmentRepository(db)
	deptService := service.NewDepartmentService(deptRepo)
	deptHandler := handler.NewDepartmentHandler(deptService)

	departments := rg.Group("/departments")
	{
		departments.POST("", deptHandler.CreateDepartment)
		departments.GET("", deptHandler.GetAllDepartments)
		departments.GET("/:id", deptHandler.GetDepartmentByID)
		departments.PUT("/:id", deptHandler.UpdateDepartment)
		departments.DELETE("/:id", deptHandler.DeleteDepartment)
	}
}