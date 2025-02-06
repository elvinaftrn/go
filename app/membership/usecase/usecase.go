package usecase

import (
	"context"
	"net/http"
	"rental/app"
	"rental/app/membership"
	"rental/app/tools"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MembershipUsecase struct {
	repo membership.MembershipRepository
	ctx  context.Context
}

func NewMembershipUsecase(repo membership.MembershipRepository, ctx context.Context) *MembershipUsecase {
	return &MembershipUsecase{
		repo: repo,
		ctx:  ctx,
	}
}

func (uc *MembershipUsecase) GetAllMemberships(c *gin.Context) ([]app.Membership, *tools.Pagination, int, error) {
	pagination, err := tools.Paginate(c)
	if err != nil {
		return nil, nil, http.StatusBadRequest, err
	}

	result, pagination, err := uc.repo.GetAllMembership(uc.ctx, pagination)
	if err != nil {
		return nil, nil, http.StatusInternalServerError, err
	}

	return result, pagination, http.StatusOK, nil
}

func (uc *MembershipUsecase) GetDetailMembership(c *gin.Context) (*app.Membership, int, error) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 0)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	result, err := uc.repo.GetDetailMembership(uc.ctx, uint(id))
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return result, http.StatusOK, nil
}

func (uc *MembershipUsecase) CreateMembership(c *gin.Context) (int, error) {
	var createMembership *app.Membership
	err := c.ShouldBindJSON(&createMembership)
	if err != nil {
		return http.StatusBadRequest, err
	}

	err = uc.repo.CreateMembership(uc.ctx, createMembership)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusCreated, nil
}

func (uc *MembershipUsecase) UpdateMembership(c *gin.Context) (int, error) {
	var updateMembership *app.Membership
	err := c.ShouldBindJSON(&updateMembership)
	if err != nil {
		return http.StatusBadRequest, err
	}

	ID, err := strconv.ParseInt(c.Param("id"), 10, 0)
	if err != nil {
		return http.StatusBadRequest, err
	}

	updateMembership.ID = uint(ID)

	err = uc.repo.UpdateMembership(uc.ctx, updateMembership)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusAccepted, nil
}

func (uc *MembershipUsecase) DeleteMembership(c *gin.Context) (int, error) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 0)
	if err != nil {
		return http.StatusBadRequest, err
	}

	err = uc.repo.DeleteMembership(uc.ctx, uint(id))
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusNoContent, nil
}
