package modules

import (
	d "backend/database"
)

func DeletePlant (id int) ( plant string, _ bool) {

	result := d.DB.Raw("delete from plants where plant_id = ? returning plant_name", id).Scan(&plant)
 
	return plant, result.RowsAffected > 0
}