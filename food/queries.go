package food

const (
	QueryInsertFood = `
	INSERT INTO user
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
	SELECT LAST_INSERT_ID();`
)
