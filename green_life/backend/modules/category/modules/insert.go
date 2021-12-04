package modules

import (
	. "backend/modules/category/models"
	d "backend/database"
)

func InsertCategory (newCategory PostCategory) (category Category) {

	d.DB.Raw("insert into categories (name) values (?) returning *", newCategory.Name).Scan(&category)
 
	return category
}