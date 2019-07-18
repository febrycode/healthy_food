package food

const (
	QueryInsertFood = `
		INSERT INTO food
			(
				user_id,
				province_id,
				title,
				created_at
			)
		VALUES 
			(
				:user_id,
				:province_id,
				:title,
				:created_at
			);
	`

	QueryGetListFood = `
		SELECT
			id,
			user_id,
			province_id,
			title
		FROM
			food;
	`
)
