package handler

import (
	"rental/app/car"
	"rental/app/tools"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type CarHandler struct {
	uc  car.CarUsecase
	log *logrus.Entry
}

func CarRoute(uc car.CarUsecase, r *gin.RouterGroup, log *logrus.Entry) {
	h := CarHandler{
		uc:  uc,
		log: log,
	}

	v2 := r.Group("car")

	v2.GET("", h.GetAllCars)
	v2.GET("/:id", h.GetDetailCar)
	v2.POST("", h.CreateCar)
	v2.PUT("/:id", h.UpdateCar)
	v2.DELETE("/:id", h.DeleteCar)
}

func (h *CarHandler) GetAllCars(c *gin.Context) {
	result, pagination, code, err := h.uc.GetAllCars(c)
	if err != nil {
		h.log.Errorf("get all cars handlers: %v", err)
		c.AbortWithStatusJSON(code, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(code, tools.Response{
		Data:    result,
		Status:  "success",
		Message: "Get All Car",
		Meta:    pagination,
	})
}

func (h *CarHandler) GetDetailCar(c *gin.Context) {
	result, code, err := h.uc.GetDetailCar(c)
	if err != nil {
		h.log.Errorf("get detail car handlers: %v", err)
		c.AbortWithStatusJSON(code, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(code, tools.Response{
		Data:    result,
		Status:  "success",
		Message: "Get Detail Car",
	})
}

func (h *CarHandler) CreateCar(c *gin.Context) {
	code, err := h.uc.CreateCar(c)
	if err != nil {
		h.log.Errorf("create car handlers: %v", err)
		c.AbortWithStatusJSON(code, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(code, tools.Response{
		Status:  "success",
		Message: "Create Car",
	})
}

func (h *CarHandler) UpdateCar(c *gin.Context) {
	code, err := h.uc.UpdateCar(c)
	if err != nil {
		h.log.Errorf("update car handlers: %v", err)
		c.AbortWithStatusJSON(code, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(code, tools.Response{
		Status:  "success",
		Message: "Update Car",
	})
}

func (h *CarHandler) DeleteCar(c *gin.Context) {
	code, err := h.uc.DeleteCar(c)
	if err != nil {
		h.log.Errorf("delete car handlers: %v", err)
		c.AbortWithStatusJSON(code, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(code, tools.Response{
		Status:  "success",
		Message: "Delete Car",
	})
}
