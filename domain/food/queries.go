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
			title,
			COALESCE(created_at, timestamp '0001-01-01 00:00:00') as created_at,
			COALESCE(updated_at, timestamp '0001-01-01 00:00:00') as updated_at
		FROM
			food
		ORDER BY id DESC;
	`

	QueryGetListFoodByTitle = `
		SELECT
			id,
			user_id,
			province_id,
			title,
			COALESCE(created_at, timestamp '0001-01-01 00:00:00') as created_at,
			COALESCE(updated_at, timestamp '0001-01-01 00:00:00') as updated_at
		FROM
			food
		WHERE
			LOWER(title) LIKE ?
		ORDER BY id DESC;
	`
)
