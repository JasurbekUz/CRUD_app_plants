package modules

import (
	. "backend/modules/category/models"
	d "backend/database"
)



func DeleteCtg (id int) ( category Category, _ bool) {

	result := d.DB.Raw("delete from categories where category_id = ? returning *", id).Scan(&category)
 
	return category, result.RowsAffected > 0
}