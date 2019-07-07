package image

const (
	QueryInsertImage = `
		INSERT INTO image
			(
				reference_type,
				reference_id,
				name,
				description,
				created_at
			)
		VALUES 
			(
				:reference_type,
				:reference_id,
				:name,
				:description,
				:created_at
			);
	`
)
