package modules

import (
	"fmt"
	. "backend/modules/plant/models"
	d "backend/database"
)

func UpdatePlant (id int, updatedPlant PostPlant) (plant GetPlant, _ bool) {

	
	result := d.DB.Raw(
		`update 
			plants 
		set 
			category_id = coalesce(?, category_id), 
			plant_name = coalesce(? , plant_name), 
			quantity = coalesce(?, quantity) 
		where 
			plant_id = ? returning plant_id`, 
		updatedPlant.CategoryId,
		updatedPlant.PlantName,
		updatedPlant.Quantity,
		id,
	).Scan(&id)

	result = d.DB.Raw("select plant_name, quantity from plants where plant_id = ?", id).Scan(&plant)
	d.DB.Raw("select * from categories where category_id = (select category_id from plants where plant_id = ?)", id).Scan(&plant.Category)
	
	fmt.Println(id)
	
	return plant, result.RowsAffected > 0
}