package courier

import (
	"github.com/ggalihpp/shopit/user"
	"github.com/jinzhu/gorm"
)

// Courier -
type Courier struct {
	gorm.Model
	Code       string       `gorm:"type:varchar(10)"`
	Name       string       `gorm:"type:varchar(100)"`
	PricePerKM int          `gorm:"type:numeric"`
	Order      []user.Order `gorm:"foreignkey:CourierCode;association_foreignkey:Code"`
}
