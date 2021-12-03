package modules

import (
	. "backend/modules/category/models"
	d "backend/database"
)

var category Category
var postCategory PostCategory

func InsertCategory (newCategory PostCategory) Category {

	d.DB.Raw("insert into categories (name) values (?) returning *", newCategory.Name).Scan(&category)
 
	return category
}