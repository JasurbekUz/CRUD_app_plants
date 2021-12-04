package modules

import (

	. "backend/modules/plant/models"
	d "backend/database"
)

func SelectAll () (plants []Plant) {

	d.DB.Find(&plants)
	
	return plants	
}