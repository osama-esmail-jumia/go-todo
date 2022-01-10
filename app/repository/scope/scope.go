package scope

import "gorm.io/gorm"

func WhereIfNotNil(attr string, completed *bool) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if completed != nil {
			db.Where(attr, completed)
		}
		return db
	}
}
