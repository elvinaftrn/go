package usecase

import (
	"context"
	"net/http"
	"rental/app"
	"rental/app/booking"
	"rental/app/car"
	"rental/app/customer"
	"rental/app/driver"
	"rental/app/tools"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookingUsecase struct {
	repo         booking.BookingRepository
	repoCar      car.CarRepository
	repoCustomer customer.CustomerRepository
	repoDriver   driver.DriverRepository
	ctx          context.Context
}

func NewBookingUsecase(repo booking.BookingRepository, repoCar car.CarRepository, repoCustomer customer.CustomerRepository, repoDriver driver.DriverRepository, ctx context.Context) *BookingUsecase {
	return &BookingUsecase{
		repo:         repo,
		repoCar:      repoCar,
		repoCustomer: repoCustomer,
		repoDriver:   repoDriver,
		ctx:          ctx,
	}
}

func (uc *BookingUsecase) GetAllBookings(c *gin.Context) ([]app.Booking, *tools.Pagination, int, error) {
	pagination, err := tools.Paginate(c)
	if err != nil {
		return nil, nil, http.StatusBadRequest, err
	}

	result, pagination, err := uc.repo.GetAllBooking(uc.ctx, pagination)
	if err != nil {
		return nil, nil, http.StatusInternalServerError, err
	}

	return result, pagination, http.StatusOK, nil
}

func (uc *BookingUsecase) GetDetailBooking(c *gin.Context) (*app.Booking, int, error) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 0)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	result, err := uc.repo.GetDetailBooking(uc.ctx, uint(id))
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return result, http.StatusOK, nil
}

func (uc *BookingUsecase) CreateBooking(c *gin.Context) (int, error) {
	var form *app.Booking
	err := c.ShouldBindJSON(&form)
	if err != nil {
		return http.StatusBadRequest, err
	}

	createBooking := app.Booking{}
	createBooking.CustomerID = form.CustomerID
	createBooking.CarID = form.CarID
	createBooking.StartRent = form.StartRent
	createBooking.EndRent = form.EndRent
	createBooking.Finished = form.Finished
	createBooking.BookingType = form.BookingType
	createBooking.DriverID = form.DriverID

	//count total cost
	if createBooking.TotalCost == nil {
		createBooking.TotalCost = new(float64)
	}

	car, err := uc.repoCar.GetDetailCar(c, createBooking.CarID)
	if err != nil {
		return http.StatusBadRequest, err
	}

	duration := createBooking.EndRent.Sub(createBooking.StartRent)
	days := duration.Hours() / 24

	*createBooking.TotalCost = days * car.DailyRent

	//count discount
	if createBooking.Discount == nil {
		createBooking.Discount = new(float64)
	}
	customer, err := uc.repoCustomer.GetDetailCustomer(c, createBooking.CustomerID)
	if err != nil {
		return http.StatusBadRequest, err
	}

	if customer.MembershipID != nil && *customer.MembershipID != 0 {
		discount := *createBooking.TotalCost * customer.Membership.DiscountPercentage / 100
		*createBooking.Discount = discount
	} else {
		*createBooking.Discount = 0
	}

	//count total driver cost

	if createBooking.TotalDriverCost == nil {
		createBooking.TotalDriverCost = new(float64)
	}

	if createBooking.DriverID != nil && *createBooking.DriverID != 0 {

		driver, err := uc.repoDriver.GetDetailDriver(c, *createBooking.DriverID)
		if err != nil {
			return http.StatusBadRequest, err
		}

		if createBooking.BookingTypeID == 2 {
			*createBooking.TotalDriverCost = days * driver.DailyCost
		} else if createBooking.BookingTypeID == 1 {
			*createBooking.TotalDriverCost = 0
		}
	}

	err = uc.repo.CreateBooking(uc.ctx, &createBooking)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusCreated, nil
}

func (uc *BookingUsecase) UpdateBooking(c *gin.Context) (int, error) {
	var updateBooking *app.Booking
	err := c.ShouldBindJSON(&updateBooking)
	if err != nil {
		return http.StatusBadRequest, err
	}

	ID, err := strconv.ParseInt(c.Param("id"), 10, 0)
	if err != nil {
		return http.StatusBadRequest, err
	}

	updateBooking.ID = uint(ID)

	err = uc.repo.UpdateBooking(uc.ctx, updateBooking)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusAccepted, nil
}

func (uc *BookingUsecase) DeleteBooking(c *gin.Context) (int, error) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 0)
	if err != nil {
		return http.StatusBadRequest, err
	}

	err = uc.repo.DeleteBooking(uc.ctx, uint(id))
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusNoContent, nil
}
