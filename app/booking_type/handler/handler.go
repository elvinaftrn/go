package handler

import (
	bookingtype "rental/app/booking_type"
	"rental/app/tools"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type BookingTypeHandler struct {
	uc  bookingtype.BookingTypeUsecase
	log *logrus.Entry
}

func BookingTypeRoute(uc bookingtype.BookingTypeUsecase, r *gin.RouterGroup, log *logrus.Entry) {
	h := BookingTypeHandler{
		uc:  uc,
		log: log,
	}

	v2 := r.Group("bookingType")

	v2.GET("", h.GetAllBookingTypes)
	v2.GET("/:id", h.GetDetailBookingType)
	v2.POST("", h.CreateBookingType)
	v2.PUT("/:id", h.UpdateBookingType)
	v2.DELETE("/:id", h.DeleteBookingType)
}

func (h *BookingTypeHandler) GetAllBookingTypes(c *gin.Context) {
	result, pagination, code, err := h.uc.GetAllBookingTypes(c)
	if err != nil {
		h.log.Errorf("get all bookingTypes handlers: %v", err)
		c.AbortWithStatusJSON(code, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(code, tools.Response{
		Data:    result,
		Status:  "success",
		Message: "Get All BookingType",
		Meta:    pagination,
	})
}

func (h *BookingTypeHandler) GetDetailBookingType(c *gin.Context) {
	result, code, err := h.uc.GetDetailBookingType(c)
	if err != nil {
		h.log.Errorf("get detail bookingType handlers: %v", err)
		c.AbortWithStatusJSON(code, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(code, tools.Response{
		Data:    result,
		Status:  "success",
		Message: "Get Detail BookingType",
	})
}

func (h *BookingTypeHandler) CreateBookingType(c *gin.Context) {
	code, err := h.uc.CreateBookingType(c)
	if err != nil {
		h.log.Errorf("create bookingType handlers: %v", err)
		c.AbortWithStatusJSON(code, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(code, tools.Response{
		Status:  "success",
		Message: "Create BookingType",
	})
}

func (h *BookingTypeHandler) UpdateBookingType(c *gin.Context) {
	code, err := h.uc.UpdateBookingType(c)
	if err != nil {
		h.log.Errorf("update bookingType handlers: %v", err)
		c.AbortWithStatusJSON(code, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(code, tools.Response{
		Status:  "success",
		Message: "Update BookingType",
	})
}

func (h *BookingTypeHandler) DeleteBookingType(c *gin.Context) {
	code, err := h.uc.DeleteBookingType(c)
	if err != nil {
		h.log.Errorf("delete bookingType handlers: %v", err)
		c.AbortWithStatusJSON(code, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(code, tools.Response{
		Status:  "success",
		Message: "Delete BookingType",
	})
}
