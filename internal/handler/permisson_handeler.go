package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/toufiqulislamtanmoy/go_flat_FBI/internal/models"
	"github.com/toufiqulislamtanmoy/go_flat_FBI/internal/service"
)

type PermissionHandler struct {
	Service *service.PermissionService
}

func NewPermissionHandler(service *service.PermissionService) *PermissionHandler {
	return &PermissionHandler{Service: service}
}

func (h *PermissionHandler) CreatePlan(c *gin.Context) {
	var permission models.Permission

	if err := c.ShouldBindJSON(&permission); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":      "error",
			"status_code": 400,
			"message":     err.Error(),
			"data":        nil,
		})
		return
	}

	if err := h.Service.CreatePlan(&permission); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":      "error",
			"status_code": 500,
			"message":     err.Error(),
			"data":        nil,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":      "success",
		"status_code": 201,
		"message":     "Role Created Successfully",
		"data":        permission,
	})
}

func (h *PermissionHandler) GetAllRole(c *gin.Context) {
	plans, err := h.Service.GetAllRoles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"massage": "Plan data fetched successfully",
		"data":    plans,
	})
}

func (h *PermissionHandler) GetPlanByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	permission, err := h.Service.GetRoleByID(uint(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "error",
			"reason": "Role not found",
			"data":   nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"reason": "Role data fetched successfully",
		"data":   permission,
	})
}

func (h *PermissionHandler) UpdateRole(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var permission models.Permission
	if err := c.ShouldBindJSON(&permission); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	permission.ID = uint(id)
	if err := h.Service.UpdateRole(&permission); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":      "success",
		"status_code": 201,
		"message":     "Role Updated Successfully",
		"data":        permission,
	})
}

func (h *PermissionHandler) DeleteRole(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.Service.DeleteRole(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Role deleted"})
}
