package modules

import (
	. "backend/modules/category/models"
	d "backend/database"
)

func UpdateCategory (id int, updatedCtg Category) (category Category, _ bool) {

	result := d.DB.Raw(
		"update categories set category_id = coalesce(?, category_id), name = coalesce(?, name) where category_id = ? returning *", 
		updatedCtg.CategoryId,
		updatedCtg.Name,
		id,
	).Scan(
		&category,
	)
 
	return category, result.RowsAffected > 0
}