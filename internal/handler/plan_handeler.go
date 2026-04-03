package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/toufiqulislamtanmoy/go_flat_FBI/internal/models"
	"github.com/toufiqulislamtanmoy/go_flat_FBI/internal/service"
)

type PlanHandler struct {
	Service *service.PlanService
}

func NewPlanHandler(service *service.PlanService) *PlanHandler {
	return &PlanHandler{Service: service}
}

func (h *PlanHandler) CreatePlan(c *gin.Context) {
	var plan models.Plan

	if err := c.ShouldBindJSON(&plan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":      "error",
			"status_code": 400,
			"message":     err.Error(),
			"data":        nil,
		})
		return
	}

	if err := h.Service.CreatePlan(&plan); err != nil {
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
		"data":        plan,
	})
}

func (h *PlanHandler) GetAllRole(c *gin.Context) {
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

func (h *PlanHandler) GetPlanByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	plan, err := h.Service.GetRoleByID(uint(id))

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
		"data":   plan,
	})
}

func (h *PlanHandler) UpdateRole(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var plan models.Plan
	if err := c.ShouldBindJSON(&plan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	plan.ID = uint(id)
	if err := h.Service.UpdateRole(&plan); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":      "success",
		"status_code": 201,
		"message":     "Role Updated Successfully",
		"data":        plan,
	})
}

func (h *PlanHandler) DeleteRole(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.Service.DeleteRole(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Role deleted"})
}
