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
)
