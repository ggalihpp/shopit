package item

import (
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

// Item -
type Item struct {
	gorm.Model
	UserID          uint
	UserName        string         `json:"user_name" gorm:"type:varchar(30)"`
	Name            string         `json:"name" gorm:"type:varchar(100)"`
	ImagesURL       pq.StringArray `json:"images_url" gorm:"type:varchar( 200)[]"`
	Keywords        pq.StringArray `json:"keywords" gorm:"type:varchar(64)[]"`
	ShortSummary    string         `json:"short_summary" gorm:"type:varchar(200)"`
	LongDescription string         `json:"long_description" gorm:"type:zdb.fulltext "`
	Price           int            `json:"price" gorm:"type:numeric"`
	DiscountPercent int            `json:"discount_percent"`
	Stock           int            `json:"stock"`
	Availability    bool           `gorm:"default:true"`
}
