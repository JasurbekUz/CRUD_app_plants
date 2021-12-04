package models

type Category struct {
	CategoryId uint16	`json:"categoryId"`
	Name string			`json:"name"`
}

// FOR POST PLANT
type PostPlant struct {
	CategoryId uint16	`json:"categoryId"`
	PlantName string 	`json:"plantName"`	
	Quantity  uint 		`json:"quantity"`
}

// FOR GET ALL PLANTS
type Plant struct {
	PlantId uint16 		`json:"plantId"`
	PlantName string	`json:"plantName"`
}

// FOR GET PLANT BY ID AND GET NEW PLANT
type GetPlant struct {
	PlantName string 	`json:"plantName"`	
	Quantity  uint 		`json:"quantity"`
	Category Category 	`json:"category"`
}



/*plant {

	name,
	Quantity,

	Category {
		id,
		name,
	},
}*/