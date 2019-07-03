package user

const (
	QueryGetUserByEmail = `
		SELECT
			id,
			email,
			name,
			avatar_url,
			address,
			bio,
			created_at,
			COALESCE(updated_at, timestamp '0001-01-01 00:00:00') as updated_at
		FROM user
		WHERE email = ?;`

	QueryInsertUser = `
		INSERT INTO user
			(
				email,
				name,
				avatar_url,
				address,
				bio,
				created_at
			)
		VALUES 
			(
				:email,
				:name,
				:avatar_url,
				:address,
				:bio,
				:created_at
			);
	`
)
