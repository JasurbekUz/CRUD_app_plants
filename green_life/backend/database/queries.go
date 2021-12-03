package database

var CT_GET_BY_ID = `
	select
		c.name,
		array_agg(p.*)
	from plants as p
	join categories as c using (category_id)
	where category_id = ?
	group by c.name
`