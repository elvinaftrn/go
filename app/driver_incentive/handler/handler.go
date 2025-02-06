package handler

import (
	driverincentive "rental/app/driver_incentive"
	"rental/app/tools"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type DriverIncentiveHandler struct {
	uc  driverincentive.DriverIncentiveUsecase
	log *logrus.Entry
}

func DriverIncentiveRoute(uc driverincentive.DriverIncentiveUsecase, r *gin.RouterGroup, log *logrus.Entry) {
	h := DriverIncentiveHandler{
		uc:  uc,
		log: log,
	}

	v2 := r.Group("driverIncentive")

	v2.GET("", h.GetAllDriverIncentives)
	v2.GET("/:id", h.GetDetailDriverIncentive)
	v2.POST("", h.CreateDriverIncentive)
	v2.DELETE("/:id", h.DeleteDriverIncentive)
}

func (h *DriverIncentiveHandler) GetAllDriverIncentives(c *gin.Context) {
	result, pagination, code, err := h.uc.GetAllDriverIncentives(c)
	if err != nil {
		h.log.Errorf("get all driverIncentives handlers: %v", err)
		c.AbortWithStatusJSON(code, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(code, tools.Response{
		Data:    result,
		Status:  "success",
		Message: "Get All DriverIncentive",
		Meta:    pagination,
	})
}

func (h *DriverIncentiveHandler) GetDetailDriverIncentive(c *gin.Context) {
	result, code, err := h.uc.GetDetailDriverIncentive(c)
	if err != nil {
		h.log.Errorf("get detail driverIncentive handlers: %v", err)
		c.AbortWithStatusJSON(code, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(code, tools.Response{
		Data:    result,
		Status:  "success",
		Message: "Get Detail DriverIncentive",
	})
}

func (h *DriverIncentiveHandler) CreateDriverIncentive(c *gin.Context) {
	code, err := h.uc.CreateDriverIncentive(c)
	if err != nil {
		h.log.Errorf("create driverIncentive handlers: %v", err)
		c.AbortWithStatusJSON(code, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(code, tools.Response{
		Status:  "success",
		Message: "Create DriverIncentive",
	})
}

func (h *DriverIncentiveHandler) DeleteDriverIncentive(c *gin.Context) {
	code, err := h.uc.DeleteDriverIncentive(c)
	if err != nil {
		h.log.Errorf("delete driverIncentive handlers: %v", err)
		c.AbortWithStatusJSON(code, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(code, tools.Response{
		Status:  "success",
		Message: "Delete DriverIncentive",
	})
}
