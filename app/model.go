package app

import (
	"time"
)

type Customer struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	Name         string `gorm:"not null" json:"name"`
	NIK          string `gorm:"unique;not null" json:"nik"`
	PhoneNumber  string `gorm:"not null" json:"phone_number"`
	MembershipID *uint  `gorm:"foreignKey:MembershipID" json:"membership_id"`

	Membership *Membership `gorm:"references:ID"`

	Booking []Booking `gorm:"foreignKey:CustomerID"`
}

type Car struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	Name      string  `gorm:"unique;not null" json:"name"`
	Stock     int     `gorm:"not null" json:"stock"`
	DailyRent float64 `gorm:"not null" json:"daily_rent"`

	Booking []Booking `gorm:"foreignKey:CarID"`
}

type Booking struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	CustomerID      uint      `gorm:"foreignKey:CustomerID" json:"customer_id"`
	CarID           uint      `gorm:"foreignKey:CarID"  json:"car_id"`
	StartRent       time.Time `gorm:"type:date" json:"start_rent"`
	EndRent         time.Time `gorm:"type:date" json:"end_rent"`
	TotalCost       *float64  `gorm:"type:int" json:"total_cost"`
	Finished        bool      `gorm:"type:bool" json:"finished"`
	Discount        *float64  `json:"discount"`
	BookingTypeID   uint      `gorm:"foreignKey:BookingTypeID" json:"booking_type_id"`
	DriverID        *uint     `gorm:"foreignKey:DriverID" json:"driver_id"`
	TotalDriverCost *float64  `json:"total_driver_cost"`

	Customer    Customer    `gorm:"references:ID"`
	Car         Car         `gorm:"references:ID"`
	BookingType BookingType `gorm:"references:ID"`
	Driver      *Driver     `gorm:"references:ID"`

	DriverIncentive []DriverIncentive `gorm:"foreignKey:BookingID"`
}

type Membership struct {
	ID                 uint    `gorm:"primaryKey"`
	Name               string  `json:"name"`
	DiscountPercentage float64 `json:"discount_percentage"`

	Customer []Customer `gorm:"foreignKey:MembershipID"`
}

type Driver struct {
	ID          uint    `gorm:"primaryKey"`
	Name        string  `json:"name"`
	NIK         string  `json:"nik"`
	PhoneNumber string  `json:"phone_number"`
	DailyCost   float64 `json:"daily_cost"`

	Booking []Booking `gorm:"foreignKey:DriverID"`
}

type DriverIncentive struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	BookingID uint    `gorm:"foreignKey:BookingID" json:"booking_id"`
	Incentive float64 `json:"incentive"`

	Booking Booking `gorm:"references:ID"`
}

type BookingType struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `json:"booking_type"`
	Description string `json:"description"`

	Booking []Booking `gorm:"foreignKey:BookingTypeID"`
}
