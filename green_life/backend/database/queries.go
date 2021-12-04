package database

var SELECT_PLANTS = `
	select
		c.category_id,
		c.name,
		p.plant_id,
		p.plant_name,
		p.quantity
	from plants as p
	join categories as c using (category_id)
`