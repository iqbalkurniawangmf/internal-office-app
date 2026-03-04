package handler

import (
    "net/http"
    "strconv"

    "internal-office-backend/internal/model"
    "internal-office-backend/internal/service"

    "github.com/gin-gonic/gin"
)

type DepartmentHandler struct {
    service service.DepartmentService
}

func NewDepartmentHandler(service service.DepartmentService) *DepartmentHandler {
    return &DepartmentHandler{service}
}

// IMPLEMENTATION 
// POST
func (h *DepartmentHandler) CreateDepartment(c *gin.Context) {
    var dept model.Department

    if err := c.ShouldBindJSON(&dept); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.service.CreateDepartment(&dept); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, dept)
}

// GET
func (h *DepartmentHandler) GetAllDepartments(c *gin.Context) {
    departments, err := h.service.GetAllDepartments()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, departments)
}

// GET :id
func (h *DepartmentHandler) GetDepartmentByID(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
        return
    }

    department, err := h.service.GetDepartmentByID(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "department not found"})
        return
    }

    c.JSON(http.StatusOK, department)
}

// PUT :id
func (h *DepartmentHandler) UpdateDepartment(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
        return
    }

    var dept model.Department
    if err := c.ShouldBindJSON(&dept); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.service.UpdateDepartment(uint(id), &dept); err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "department updated"})
}

// DELETE :id
func (h *DepartmentHandler) DeleteDepartment(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
        return
    }

    if err := h.service.DeleteDepartment(uint(id)); err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "department deleted"})
}