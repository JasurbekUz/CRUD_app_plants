package modules

import (

	. "backend/modules/category/models"
	d "backend/database"
)

func SelectAll () (categories []Category) {

	d.DB.Find(&categories)
	
	return categories	
}