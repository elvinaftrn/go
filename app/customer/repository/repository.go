package repository

import (
	"context"
	"rental/app"
	"rental/app/tools"

	"gorm.io/gorm"
)

type CustomerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *CustomerRepository {
	return &CustomerRepository{
		db: db,
	}
}

func (r *CustomerRepository) GetAllCustomer(ctx context.Context, pagination *tools.Pagination) ([]app.Customer, *tools.Pagination, error) {
	var customers []app.Customer

	if err := r.db.Offset(pagination.Offset).Limit(pagination.Limit).Preload("Membership").Find(&customers).Error; err != nil {
		return nil, nil, err
	}

	var count int64
	if err := r.db.Model(&app.Customer{}).Count(&count).Error; err != nil {
		return nil, nil, err
	}

	pagination.Count = int(count)

	pagination = tools.Paging(pagination)
	return customers, pagination, nil
}

func (r *CustomerRepository) GetDetailCustomer(ctx context.Context, ID uint) (*app.Customer, error) {
	var customer app.Customer
	if err := r.db.Preload("Membership").First(&customer, ID).Error; err != nil {
		return nil, err
	}
	return &customer, nil
}

func (r *CustomerRepository) CreateCustomer(ctx context.Context, form *app.Customer) error {
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

func (r *CustomerRepository) UpdateCustomer(ctx context.Context, form *app.Customer) error {
	var cus *app.Customer

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

func (r *CustomerRepository) DeleteCustomer(ctx context.Context, ID uint) error {
	var cus *app.Customer

	tx := r.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Model(&cus).Where("id = ?", ID).First(&cus).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(&cus).Where("id = ?", ID).Delete(&app.Customer{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}
