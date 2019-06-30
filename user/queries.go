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
			updated_at
		FROM user
		WHERE email = ?;`
)
