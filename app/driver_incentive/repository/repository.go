package repository

import (
	"context"
	"rental/app"
	"rental/app/tools"

	"gorm.io/gorm"
)

type DriverIncentiveRepository struct {
	db *gorm.DB
}

func NewDriverIncentiveRepository(db *gorm.DB) *DriverIncentiveRepository {
	return &DriverIncentiveRepository{
		db: db,
	}
}

func (r *DriverIncentiveRepository) GetAllDriverIncentive(ctx context.Context, pagination *tools.Pagination) ([]app.DriverIncentive, *tools.Pagination, error) {
	var driverIncentives []app.DriverIncentive

	if err := r.db.Offset(pagination.Offset).Limit(pagination.Limit).Preload("Customer").Preload("Car").Find(&driverIncentives).Error; err != nil {
		return nil, nil, err
	}

	var count int64
	if err := r.db.Model(&app.DriverIncentive{}).Count(&count).Error; err != nil {
		return nil, nil, err
	}

	pagination.Count = int(count)

	pagination = tools.Paging(pagination)
	return driverIncentives, pagination, nil
}

func (r *DriverIncentiveRepository) GetDetailDriverIncentive(ctx context.Context, ID uint) (*app.DriverIncentive, error) {
	var driverIncentive app.DriverIncentive
	if err := r.db.Preload("Customer").Preload("Car").First(&driverIncentive, ID).Error; err != nil {
		return nil, err
	}
	return &driverIncentive, nil
}

func (r *DriverIncentiveRepository) CreateDriverIncentive(ctx context.Context, form *app.DriverIncentive) error {
	tx := r.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Create(form).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func (r *DriverIncentiveRepository) DeleteDriverIncentive(ctx context.Context, ID uint) error {
	var cus *app.DriverIncentive

	tx := r.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Model(&cus).Where("id = ?", ID).First(&cus).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(&cus).Where("id = ?", ID).Delete(&app.DriverIncentive{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}
