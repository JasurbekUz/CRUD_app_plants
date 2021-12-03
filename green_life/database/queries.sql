select
	c.name,
	array_agg(p.*)
from plants as p
join categories as c using (category_id)
where category_id = 1
group by c.name
;
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