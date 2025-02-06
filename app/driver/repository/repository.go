package repository

import (
	"context"
	"rental/app"
	"rental/app/tools"

	"gorm.io/gorm"
)

type DriverRepository struct {
	db *gorm.DB
}

func NewDriverRepository(db *gorm.DB) *DriverRepository {
	return &DriverRepository{
		db: db,
	}
}

func (r *DriverRepository) GetAllDriver(ctx context.Context, pagination *tools.Pagination) ([]app.Driver, *tools.Pagination, error) {
	var drivers []app.Driver

	if err := r.db.Offset(pagination.Offset).Limit(pagination.Limit).Find(&drivers).Error; err != nil {
		return nil, nil, err
	}

	var count int64
	if err := r.db.Model(&app.Driver{}).Count(&count).Error; err != nil {
		return nil, nil, err
	}

	pagination.Count = int(count)

	pagination = tools.Paging(pagination)
	return drivers, pagination, nil
}

func (r *DriverRepository) GetDetailDriver(ctx context.Context, ID uint) (*app.Driver, error) {
	var driver app.Driver
	if err := r.db.First(&driver, ID).Error; err != nil {
		return nil, err
	}
	return &driver, nil
}

func (r *DriverRepository) CreateDriver(ctx context.Context, form *app.Driver) error {
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

func (r *DriverRepository) UpdateDriver(ctx context.Context, form *app.Driver) error {
	var cus *app.Driver

	tx := r.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	if err := tx.Model(&cus).Where("id = ?", form.ID).First(&cus).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(&cus).Where("id = ?", form.ID).Updates(form).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func (r *DriverRepository) DeleteDriver(ctx context.Context, ID uint) error {
	var cus *app.Driver

	tx := r.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Model(&cus).Where("id = ?", ID).First(&cus).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(&cus).Where("id = ?", ID).Delete(&app.Driver{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}
