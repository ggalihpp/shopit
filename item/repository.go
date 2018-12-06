package item

import "github.com/jinzhu/gorm"

// Repository is
type Repository struct {
	db *gorm.DB
}

// Connection - Maintain connection
func Connection(database *gorm.DB) Repository {
	return Repository{database}
}

// InsertItem - Will insert new item to inventory
func (r *Repository) InsertItem(data *Item) (err error) {

	err = r.db.Create(data).Error

	return
}

// SearchItem - Will to querying a search
func (r *Repository) SearchItem(query string) (res []Item, err error) {
	err = r.db.Where("items ==> ?", query).Find(&res).Error

	return
}
