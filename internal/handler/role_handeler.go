package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/toufiqulislamtanmoy/go_flat_FBI/internal/models"
	"github.com/toufiqulislamtanmoy/go_flat_FBI/internal/service"
)

type RoleHandler struct {
	Service *service.RoleService
}

func NewRoleHandler(service *service.RoleService) *RoleHandler {
	return &RoleHandler{Service: service}
}

func (h *RoleHandler) CreateRole(c *gin.Context) {
	var role models.Role

	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":      "error",
			"status_code": 400,
			"message":     err.Error(),
			"data":        nil,
		})
		return
	}

	if err := h.Service.CreateRole(&role); err != nil {
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
		"data":        role,
	})
}

func (h *RoleHandler) GetAllRole(c *gin.Context) {
	roles, err := h.Service.GetAllRoles()
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
		"massage": "User data fetched successfully",
		"data":    roles,
	})
}

func (h *RoleHandler) GetRoleByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	role, err := h.Service.GetRoleByID(uint(id))

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
		"data":   role,
	})
}

func (h *RoleHandler) UpdateRole(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var role models.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	role.ID = uint(id)
	if err := h.Service.UpdateRole(&role); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":      "success",
		"status_code": 201,
		"message":     "Role Updated Successfully",
		"data":        role,
	})
}

func (h *RoleHandler) DeleteRole(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.Service.DeleteRole(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Role deleted"})
}
