package province

const (
	QueryGetAllProvince = `
		SELECT
			id,
			name,
			COALESCE(created_at, timestamp '0001-01-01 00:00:00') as created_at,
			COALESCE(updated_at, timestamp '0001-01-01 00:00:00') as updated_at
		FROM
			province;
	`

	QueryGetProvinceByID = `
		SELECT
			id,
			name,
			COALESCE(created_at, timestamp '0001-01-01 00:00:00') as created_at,
			COALESCE(updated_at, timestamp '0001-01-01 00:00:00') as updated_at
		FROM
			province
		WHERE
			id = ?;
	`
)
