package models

type Category struct {
	CategoryId uint16	`json:"categoryId"`
	Name string 		`json:"name"`
}

type Plant struct {
	PlantId uint16 		`json:"plantId"`
	CategoryId uint16   `json:"categoryId"`
	PlantName string    `json:"plantName"`
	Quantity uint 		`json:"quaintity"`	
}

type GetCategoryFieldsById struct {
	Name string			`json:"name"`
	Plants []Plant		`json:"plants"` //gorm:"foreignKey:LessonRefer"`
}

type PostCategory struct {
	Name string 	`json:"name"`
}