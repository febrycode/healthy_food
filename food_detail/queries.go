package food_detail

const (
	QueryInsertFoodDetail = `
		INSERT INTO food_detail
			(
				reference_type,
				reference_id,
				description,
				created_at
			)
		VALUES 
			(
				:reference_type,
				:reference_id,
				:description,
				:created_at
			);
	`
)
