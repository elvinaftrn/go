package router

import (
	"context"
	"rental/app/middleware"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	customerHandler "rental/app/customer/handler"
	customerRepo "rental/app/customer/repository"
	customerUC "rental/app/customer/usecase"

	carHandler "rental/app/car/handler"
	carRepo "rental/app/car/repository"
	carUC "rental/app/car/usecase"

	bookingHandler "rental/app/booking/handler"
	bookingRepo "rental/app/booking/repository"
	bookingUC "rental/app/booking/usecase"

	membershipHandler "rental/app/membership/handler"
	membershipRepo "rental/app/membership/repository"
	membershipUC "rental/app/membership/usecase"

	driverHandler "rental/app/driver/handler"
	driverRepo "rental/app/driver/repository"
	driverUC "rental/app/driver/usecase"

	bookingtypeHandler "rental/app/booking_type/handler"
	bookingtypeRepo "rental/app/booking_type/repository"
	bookingtypeUC "rental/app/booking_type/usecase"

	driverincentiveHandler "rental/app/driver_incentive/handler"
	driverincentiveRepo "rental/app/driver_incentive/repository"
	driverincentiveUC "rental/app/driver_incentive/usecase"
)

type Handlers struct {
	Ctx context.Context
	DB  *gorm.DB
	R   *gin.Engine
	Log *logrus.Entry
}

func (handlers *Handlers) Routes() {
	middleware.Add(handlers.R, middleware.CORSMiddleware())
	v1 := handlers.R.Group("api")

	// Repository
	CustomerRepo := customerRepo.NewCustomerRepository(handlers.DB)
	CarRepo := carRepo.NewCarRepository(handlers.DB)
	BookingRepo := bookingRepo.NewBookingRepository(handlers.DB)
	MembershipRepo := membershipRepo.NewMembershipRepository(handlers.DB)
	DriverRepo := driverRepo.NewDriverRepository(handlers.DB)
	BookingTypeRepo := bookingtypeRepo.NewBookingTypeRepository(handlers.DB)
	DriverIncentiveRepo := driverincentiveRepo.NewDriverIncentiveRepository(handlers.DB)

	// Usecase
	CustomerUC := customerUC.NewCustomerUsecase(CustomerRepo, handlers.Ctx)
	CarUC := carUC.NewCarUsecase(CarRepo, handlers.Ctx)
	BookingUC := bookingUC.NewBookingUsecase(BookingRepo, CarRepo, CustomerRepo, DriverRepo, handlers.Ctx)
	MembershipUC := membershipUC.NewMembershipUsecase(MembershipRepo, handlers.Ctx)
	DriverUC := driverUC.NewDriverUsecase(DriverRepo, handlers.Ctx)
	BookingTypeUC := bookingtypeUC.NewBookingTypeUsecase(BookingTypeRepo, handlers.Ctx)
	DriverIncentiveUC := driverincentiveUC.NewDriverIncentiveUsecase(DriverIncentiveRepo, BookingRepo, CarRepo, handlers.Ctx)

	// Handler
	customerHandler.CustomerRoute(CustomerUC, v1, handlers.Log)
	carHandler.CarRoute(CarUC, v1, handlers.Log)
	bookingHandler.BookingRoute(BookingUC, v1, handlers.Log)
	membershipHandler.MembershipRoute(MembershipUC, v1, handlers.Log)
	driverHandler.DriverRoute(DriverUC, v1, handlers.Log)
	bookingtypeHandler.BookingTypeRoute(BookingTypeUC, v1, handlers.Log)
	driverincentiveHandler.DriverIncentiveRoute(DriverIncentiveUC, v1, handlers.Log)

}
