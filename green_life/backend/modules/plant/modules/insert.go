package modules

import (
	. "backend/modules/plant/models"
	d "backend/database"
)

func InsertPlant (newPlant PostPlant) (plant GetPlant) {

	var id uint

	d.DB.Raw(
		"insert into plants (category_id, plant_name, quantity) values (?, ?, ?) returning plant_id",
		newPlant.CategoryId,
		newPlant.PlantName,
		newPlant.Quantity,
	).Scan(&id)

	d.DB.Raw("select plant_name, quantity from plants where plant_id = ?", id).Scan(&plant)
	d.DB.Raw("select * from categories where category_id = (select category_id from plants where plant_id = ?)", id).Scan(&plant.Category)	
 
	return plant
}