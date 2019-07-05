package food

const (
	QueryInsertFood = `
	INSERT INTO food
		(
			province_id,
			title,
			created_at
		)
	VALUES 
		(
			:province_id,
			:title,
			:created_at
		);
	`
)
