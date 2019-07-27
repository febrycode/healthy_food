package user

const (
	QueryGetListUser = `
		SELECT
			id,
			email,
			name,
			address,
			bio,
			is_admin,
			created_at,
			COALESCE(updated_at, timestamp '0001-01-01 00:00:00') as updated_at
		FROM
			user
		WHERE
			is_admin = 0
		ORDER BY
			id DESC;
	`
	QueryGetUserByEmail = `
		SELECT
			id,
			email,
			password,
			name,
			address,
			bio,
			is_admin,
			created_at,
			COALESCE(updated_at, timestamp '0001-01-01 00:00:00') as updated_at
		FROM 
			user
		WHERE
			email = ?;
	`

	QueryGetUserByUserID = `
		SELECT
			id,
			email,
			name,
			address,
			bio,
			created_at,
			COALESCE(updated_at, timestamp '0001-01-01 00:00:00') as updated_at
		FROM
			user
		WHERE
			id = ?;
	`

	QueryInsertUser = `
		INSERT INTO user
			(
				email,
				password,
				name,
				address,
				bio,
				created_at
			)
		VALUES 
			(
				:email,
				:password,
				:name,
				:address,
				:bio,
				:created_at
			);
	`

	QueryUpdateUser = `
		UPDATE
			user
		SET
			email = :email,
			name = :name,
			address = :address,
			bio = :bio,
			updated_at = :updated_at
		WHERE
			id = :id;
	`
)
