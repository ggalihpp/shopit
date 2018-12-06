package user

import (
	"time"

	"github.com/ggalihpp/shopit/item"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

// User - Main table of User
type User struct {
	gorm.Model
	Username     string `gorm:"type:varchar(30)" json:"user_name"`
	Password     string `gorm:"-" json:"password"`
	Email        string `gorm:"type:varchar(60)" json:"email"`
	HashPassword string `gorm:"type:varchar(200)"`
	PhoneNumber  string `gorm:"type:varchar(20)" json:"phone_number"`
	Addresses    []Address
	BankAccounts []BankAccount
	Items        []item.Item
	Orders       []Order
}

// Address - Will contains list of address of users
type Address struct {
	gorm.Model
	UserID  uint
	Address string
}

// TableName - Will change the table name
func (Address) TableName() string {
	return "user_addresses"
}

// BankAccount - Will contains list of bank account of users
type BankAccount struct {
	gorm.Model
	UserID        uint
	BankCode      string `gorm:"type:varchar(10)"`
	AccountNumber string `gorm:"type:varchar(20)"`
}

// TableName - Will change the table name
func (BankAccount) TableName() string {
	return "user_banks"
}

// Order - Will contain all order of users
type Order struct {
	gorm.Model
	UserID      uint
	TotalPrice  int    `gorm:"type:numeric"`
	Discount    int    `gorm:"type:numeric"`
	CourierCode string `gorm:"type:varchar(10)"`
	ExpiredDate time.Time
	IsSucceed   bool
	OrderItems  []OrderItem
}

// OrderItem - Will contain all item of an order
type OrderItem struct {
	gorm.Model
	OrderID uint
	// ITEM INFO
	UserID          uint
	name            string         `gorm:"type:varchar(100)"`
	ImagesURL       pq.StringArray `gorm:"type:varchar(200)[]"`
	ShortSummary    string         `gorm:"type:varchar(200)"`
	LongDescription string
	//////////
	Price    int `gorm:"type:numeric"`
	Quantity int
}
