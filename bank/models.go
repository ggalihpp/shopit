package bank

import (
	"github.com/ggalihpp/shopit/user"
	"github.com/jinzhu/gorm"
)

// Bank -
type Bank struct {
	gorm.Model
	Code  string             `gorm:"type:varchar(10)"`
	Name  string             `gorm:"type:varchar(50)"`
	Users []user.BankAccount `gorm:"foreignkey:BankCode;association_foreignkey:Code"`
}
