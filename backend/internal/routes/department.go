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

    rg.POST("", deptHandler.CreateDepartment)
    rg.GET("", deptHandler.GetAllDepartments)
    rg.GET("/:id", deptHandler.GetDepartmentByID)
    rg.PUT("/:id", deptHandler.UpdateDepartment)
    rg.DELETE("/:id", deptHandler.DeleteDepartment)
}