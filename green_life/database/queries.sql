select
	c.name,
	array_agg(p.*)
from plants as p
join categories as c using (category_id)
where category_id = 1
group by c.name
;

select
	c.name,
	array_agg(p.*)
from plants as p
join categories as c using (category_id)
group by c.name
;

update 
	plants 
set 
	category_id = coalesce(2, category_id), 
	plant_name = coalesce(null , plant_name), 
	quantity = coalesce(100, quantity) 
where 
	plant_id = 10;

/*
{	
	1,
	manzarali uy o'simliklari,
	
	{
		{
			1,
			1,
			limon,
			30,
		},

	 	{
	 		4,
	 		1,
	 		Mandarin,
	 		30
	 	},
	 },
}
*/