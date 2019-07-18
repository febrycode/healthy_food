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

	QueryGetImageByName = `
		SELECT
			id,
			reference_type,
			reference_id,
			name,
			description,
			created_at
		FROM
			image
		WHERE
			name = ?;
	`

	QueryUpdateImage = `
		UPDATE
			image
		SET
			reference_type = :reference_type,
			reference_id = :reference_id,
			description = :description,
			updated_at = :updated_at
		WHERE
			name = :name;
			`
)
