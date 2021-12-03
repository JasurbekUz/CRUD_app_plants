package modules

import (

	. "backend/modules/category/models"
	d "backend/database"
)

var categories []Category

func SelectAll () []Category {

	d.DB.Find(&categories)
	
	return categories	
}