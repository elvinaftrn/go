package handler

import (
	"rental/app/driver"
	"rental/app/tools"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type DriverHandler struct {
	uc  driver.DriverUsecase
	log *logrus.Entry
}

func DriverRoute(uc driver.DriverUsecase, r *gin.RouterGroup, log *logrus.Entry) {
	h := DriverHandler{
		uc:  uc,
		log: log,
	}

	v2 := r.Group("driver")

	v2.GET("", h.GetAllDrivers)
	v2.GET("/:id", h.GetDetailDriver)
	v2.POST("", h.CreateDriver)
	v2.PUT("/:id", h.UpdateDriver)
	v2.DELETE("/:id", h.DeleteDriver)
}

func (h *DriverHandler) GetAllDrivers(c *gin.Context) {
	result, pagination, code, err := h.uc.GetAllDrivers(c)
	if err != nil {
		h.log.Errorf("get all drivers handlers: %v", err)
		c.AbortWithStatusJSON(code, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(code, tools.Response{
		Data:    result,
		Status:  "success",
		Message: "Get All Driver",
		Meta:    pagination,
	})
}

func (h *DriverHandler) GetDetailDriver(c *gin.Context) {
	result, code, err := h.uc.GetDetailDriver(c)
	if err != nil {
		h.log.Errorf("get detail driver handlers: %v", err)
		c.AbortWithStatusJSON(code, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(code, tools.Response{
		Data:    result,
		Status:  "success",
		Message: "Get Detail Driver",
	})
}

func (h *DriverHandler) CreateDriver(c *gin.Context) {
	code, err := h.uc.CreateDriver(c)
	if err != nil {
		h.log.Errorf("create driver handlers: %v", err)
		c.AbortWithStatusJSON(code, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(code, tools.Response{
		Status:  "success",
		Message: "Create Driver",
	})
}

func (h *DriverHandler) UpdateDriver(c *gin.Context) {
	code, err := h.uc.UpdateDriver(c)
	if err != nil {
		h.log.Errorf("update driver handlers: %v", err)
		c.AbortWithStatusJSON(code, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(code, tools.Response{
		Status:  "success",
		Message: "Update Driver",
	})
}

func (h *DriverHandler) DeleteDriver(c *gin.Context) {
	code, err := h.uc.DeleteDriver(c)
	if err != nil {
		h.log.Errorf("delete driver handlers: %v", err)
		c.AbortWithStatusJSON(code, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(code, tools.Response{
		Status:  "success",
		Message: "Delete Driver",
	})
}
