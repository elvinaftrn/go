package usecase

import (
	"context"
	"net/http"
	"rental/app"
	"rental/app/driver"
	"rental/app/tools"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DriverUsecase struct {
	repo driver.DriverRepository
	ctx  context.Context
}

func NewDriverUsecase(repo driver.DriverRepository, ctx context.Context) *DriverUsecase {
	return &DriverUsecase{
		repo: repo,
		ctx:  ctx,
	}
}

func (uc *DriverUsecase) GetAllDrivers(c *gin.Context) ([]app.Driver, *tools.Pagination, int, error) {
	pagination, err := tools.Paginate(c)
	if err != nil {
		return nil, nil, http.StatusBadRequest, err
	}

	result, pagination, err := uc.repo.GetAllDriver(uc.ctx, pagination)
	if err != nil {
		return nil, nil, http.StatusInternalServerError, err
	}

	return result, pagination, http.StatusOK, nil
}

func (uc *DriverUsecase) GetDetailDriver(c *gin.Context) (*app.Driver, int, error) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 0)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	result, err := uc.repo.GetDetailDriver(uc.ctx, uint(id))
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return result, http.StatusOK, nil
}

func (uc *DriverUsecase) CreateDriver(c *gin.Context) (int, error) {
	var createDriver *app.Driver
	err := c.ShouldBindJSON(&createDriver)
	if err != nil {
		return http.StatusBadRequest, err
	}

	err = uc.repo.CreateDriver(uc.ctx, createDriver)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusCreated, nil
}

func (uc *DriverUsecase) UpdateDriver(c *gin.Context) (int, error) {
	var updateDriver *app.Driver
	err := c.ShouldBindJSON(&updateDriver)
	if err != nil {
		return http.StatusBadRequest, err
	}

	ID, err := strconv.ParseInt(c.Param("id"), 10, 0)
	if err != nil {
		return http.StatusBadRequest, err
	}

	updateDriver.ID = uint(ID)

	err = uc.repo.UpdateDriver(uc.ctx, updateDriver)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusAccepted, nil
}

func (uc *DriverUsecase) DeleteDriver(c *gin.Context) (int, error) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 0)
	if err != nil {
		return http.StatusBadRequest, err
	}

	err = uc.repo.DeleteDriver(uc.ctx, uint(id))
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusNoContent, nil
}
