package modules

import (

	. "backend/modules/category/models"
	d "backend/database"
)

func SelectById (id int) (categoryFields GetCategoryFieldsById, _ bool) {

	result := d.DB.Raw("select name from categories where category_id = ?", id).Scan(&categoryFields.Name)
	d.DB.Raw("select * from plants where category_id = ?", id).Scan(&categoryFields.Plants)
	
	return categoryFields, result.RowsAffected > 0
}

