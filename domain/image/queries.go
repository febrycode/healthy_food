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
			COALESCE(created_at, timestamp '0001-01-01 00:00:00') as created_at,
			COALESCE(updated_at, timestamp '0001-01-01 00:00:00') as updated_at
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

	QueryGetImageByReferenceID = `
		SELECT
			id,
			reference_type,
			reference_id,
			name,
			description,
			COALESCE(created_at, timestamp '0001-01-01 00:00:00') as created_at,
			COALESCE(updated_at, timestamp '0001-01-01 00:00:00') as updated_at
		FROM
			image
		WHERE
			reference_id = ?;
	`
)
