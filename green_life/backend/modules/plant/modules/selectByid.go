package modules

import (

	. "backend/modules/plant/models"
	d "backend/database"
)

func SelectById (id int) (plant GetPlant, _ bool) {

	result := d.DB.Raw("select plant_name, quantity from plants where plant_id = ?", id).Scan(&plant)
	d.DB.Raw("select * from categories where category_id = (select category_id from plants where plant_id = ?)", id).Scan(&plant.Category)
	
	return plant, result.RowsAffected > 0
}

