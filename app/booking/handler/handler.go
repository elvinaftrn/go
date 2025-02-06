package handler

import (
	"rental/app/booking"
	"rental/app/tools"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type BookingHandler struct {
	uc  booking.BookingUsecase
	log *logrus.Entry
}

func BookingRoute(uc booking.BookingUsecase, r *gin.RouterGroup, log *logrus.Entry) {
	h := BookingHandler{
		uc:  uc,
		log: log,
	}

	v2 := r.Group("booking")

	v2.GET("", h.GetAllBookings)
	v2.GET("/:id", h.GetDetailBooking)
	v2.POST("", h.CreateBooking)
	v2.PUT("/:id", h.UpdateBooking)
	v2.DELETE("/:id", h.DeleteBooking)
}

func (h *BookingHandler) GetAllBookings(c *gin.Context) {
	result, pagination, code, err := h.uc.GetAllBookings(c)
	if err != nil {
		h.log.Errorf("get all bookings handlers: %v", err)
		c.AbortWithStatusJSON(code, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(code, tools.Response{
		Data:    result,
		Status:  "success",
		Message: "Get All Booking",
		Meta:    pagination,
	})
}

func (h *BookingHandler) GetDetailBooking(c *gin.Context) {
	result, code, err := h.uc.GetDetailBooking(c)
	if err != nil {
		h.log.Errorf("get detail booking handlers: %v", err)
		c.AbortWithStatusJSON(code, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(code, tools.Response{
		Data:    result,
		Status:  "success",
		Message: "Get Detail Booking",
	})
}

func (h *BookingHandler) CreateBooking(c *gin.Context) {
	code, err := h.uc.CreateBooking(c)
	if err != nil {
		h.log.Errorf("create booking handlers: %v", err)
		c.AbortWithStatusJSON(code, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(code, tools.Response{
		Status:  "success",
		Message: "Create Booking",
	})
}

func (h *BookingHandler) UpdateBooking(c *gin.Context) {
	code, err := h.uc.UpdateBooking(c)
	if err != nil {
		h.log.Errorf("update booking handlers: %v", err)
		c.AbortWithStatusJSON(code, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(code, tools.Response{
		Status:  "success",
		Message: "Update Booking",
	})
}

func (h *BookingHandler) DeleteBooking(c *gin.Context) {
	code, err := h.uc.DeleteBooking(c)
	if err != nil {
		h.log.Errorf("delete booking handlers: %v", err)
		c.AbortWithStatusJSON(code, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(code, tools.Response{
		Status:  "success",
		Message: "Delete Booking",
	})
}
