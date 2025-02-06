package usecase

import (
	"context"
	"net/http"
	"rental/app"
	"rental/app/car"
	"rental/app/tools"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CarUsecase struct {
	repo car.CarRepository
	ctx  context.Context
}

func NewCarUsecase(repo car.CarRepository, ctx context.Context) *CarUsecase {
	return &CarUsecase{
		repo: repo,
		ctx:  ctx,
	}
}

func (uc *CarUsecase) GetAllCars(c *gin.Context) ([]app.Car, *tools.Pagination, int, error) {
	pagination, err := tools.Paginate(c)
	if err != nil {
		return nil, nil, http.StatusBadRequest, err
	}

	result, pagination, err := uc.repo.GetAllCar(uc.ctx, pagination)
	if err != nil {
		return nil, nil, http.StatusInternalServerError, err
	}

	return result, pagination, http.StatusOK, nil
}

func (uc *CarUsecase) GetDetailCar(c *gin.Context) (*app.Car, int, error) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 0)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	result, err := uc.repo.GetDetailCar(uc.ctx, uint(id))
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return result, http.StatusOK, nil
}

func (uc *CarUsecase) CreateCar(c *gin.Context) (int, error) {
	var createCar *app.Car
	err := c.ShouldBindJSON(&createCar)
	if err != nil {
		return http.StatusBadRequest, err
	}

	err = uc.repo.CreateCar(uc.ctx, createCar)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusCreated, nil
}

func (uc *CarUsecase) UpdateCar(c *gin.Context) (int, error) {
	var updateCar *app.Car
	err := c.ShouldBindJSON(&updateCar)
	if err != nil {
		return http.StatusBadRequest, err
	}

	ID, err := strconv.ParseInt(c.Param("id"), 10, 0)
	if err != nil {
		return http.StatusBadRequest, err
	}

	updateCar.ID = uint(ID)

	err = uc.repo.UpdateCar(uc.ctx, updateCar)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusAccepted, nil
}

func (uc *CarUsecase) DeleteCar(c *gin.Context) (int, error) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 0)
	if err != nil {
		return http.StatusBadRequest, err
	}

	err = uc.repo.DeleteCar(uc.ctx, uint(id))
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusNoContent, nil
}
