package handler

import (
	"rental/app/membership"
	"rental/app/tools"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type MembershipHandler struct {
	uc  membership.MembershipUsecase
	log *logrus.Entry
}

func MembershipRoute(uc membership.MembershipUsecase, r *gin.RouterGroup, log *logrus.Entry) {
	h := MembershipHandler{
		uc:  uc,
		log: log,
	}

	v2 := r.Group("membership")

	v2.GET("", h.GetAllMemberships)
	v2.GET("/:id", h.GetDetailMembership)
	v2.POST("", h.CreateMembership)
	v2.PUT("/:id", h.UpdateMembership)
	v2.DELETE("/:id", h.DeleteMembership)
}

func (h *MembershipHandler) GetAllMemberships(c *gin.Context) {
	result, pagination, code, err := h.uc.GetAllMemberships(c)
	if err != nil {
		h.log.Errorf("get all memberships handlers: %v", err)
		c.AbortWithStatusJSON(code, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(code, tools.Response{
		Data:    result,
		Status:  "success",
		Message: "Get All Membership",
		Meta:    pagination,
	})
}

func (h *MembershipHandler) GetDetailMembership(c *gin.Context) {
	result, code, err := h.uc.GetDetailMembership(c)
	if err != nil {
		h.log.Errorf("get detail membership handlers: %v", err)
		c.AbortWithStatusJSON(code, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(code, tools.Response{
		Data:    result,
		Status:  "success",
		Message: "Get Detail Membership",
	})
}

func (h *MembershipHandler) CreateMembership(c *gin.Context) {
	code, err := h.uc.CreateMembership(c)
	if err != nil {
		h.log.Errorf("create membership handlers: %v", err)
		c.AbortWithStatusJSON(code, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(code, tools.Response{
		Status:  "success",
		Message: "Create Membership",
	})
}

func (h *MembershipHandler) UpdateMembership(c *gin.Context) {
	code, err := h.uc.UpdateMembership(c)
	if err != nil {
		h.log.Errorf("update membership handlers: %v", err)
		c.AbortWithStatusJSON(code, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(code, tools.Response{
		Status:  "success",
		Message: "Update Membership",
	})
}

func (h *MembershipHandler) DeleteMembership(c *gin.Context) {
	code, err := h.uc.DeleteMembership(c)
	if err != nil {
		h.log.Errorf("delete membership handlers: %v", err)
		c.AbortWithStatusJSON(code, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(code, tools.Response{
		Status:  "success",
		Message: "Delete Membership",
	})
}
