package handler

import (
	"rental/app/customer"
	"rental/app/tools"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type CustomerHandler struct {
	uc  customer.CustomerUsecase
	log *logrus.Entry
}

func CustomerRoute(uc customer.CustomerUsecase, r *gin.RouterGroup, log *logrus.Entry) {
	h := CustomerHandler{
		uc:  uc,
		log: log,
	}

	v2 := r.Group("customer")

	v2.GET("", h.GetAllCustomers)
	v2.GET("/:id", h.GetDetailCustomer)
	v2.POST("", h.CreateCustomer)
	v2.PUT("/:id", h.UpdateCustomer)
	v2.DELETE("/:id", h.DeleteCustomer)
}

func (h *CustomerHandler) GetAllCustomers(c *gin.Context) {
	result, pagination, code, err := h.uc.GetAllCustomers(c)
	if err != nil {
		h.log.Errorf("get all customers handlers: %v", err)
		c.AbortWithStatusJSON(code, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(code, tools.Response{
		Data:    result,
		Status:  "success",
		Message: "Get All Customer",
		Meta:    pagination,
	})
}

func (h *CustomerHandler) GetDetailCustomer(c *gin.Context) {
	result, code, err := h.uc.GetDetailCustomer(c)
	if err != nil {
		h.log.Errorf("get detail customer handlers: %v", err)
		c.AbortWithStatusJSON(code, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(code, tools.Response{
		Data:    result,
		Status:  "success",
		Message: "Get Detail Customer",
	})
}

func (h *CustomerHandler) CreateCustomer(c *gin.Context) {
	code, err := h.uc.CreateCustomer(c)
	if err != nil {
		h.log.Errorf("create customer handlers: %v", err)
		c.AbortWithStatusJSON(code, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(code, tools.Response{
		Status:  "success",
		Message: "Create Customer",
	})
}

func (h *CustomerHandler) UpdateCustomer(c *gin.Context) {
	code, err := h.uc.UpdateCustomer(c)
	if err != nil {
		h.log.Errorf("update customer handlers: %v", err)
		c.AbortWithStatusJSON(code, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(code, tools.Response{
		Status:  "success",
		Message: "Update Customer",
	})
}

func (h *CustomerHandler) DeleteCustomer(c *gin.Context) {
	code, err := h.uc.DeleteCustomer(c)
	if err != nil {
		h.log.Errorf("delete customer handlers: %v", err)
		c.AbortWithStatusJSON(code, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(code, tools.Response{
		Status:  "success",
		Message: "Delete Customer",
	})
}
