package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"

	"github.com/febrycode/healthy_food/models"
	"github.com/febrycode/healthy_food/user"
)

type mysqlUserRepository struct {
	Conn *sqlx.DB
}

// NewMysqlUserRepository will create an object that represent the user.Repository interface
func NewMysqlUserRepository(Conn *sqlx.DB) user.Repository {
	return &mysqlUserRepository{Conn}
}

func (m *mysqlUserRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.User, error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			logrus.Error(err)
		}
	}()

	result := make([]*models.User, 0)
	for rows.Next() {
		t := new(models.User)
		err = rows.Scan(
			&t.ID,
			&t.Email,
			&t.Name,
			&t.AvatarURL,
			&t.Address,
			&t.Bio,
			&t.CreatedAt,
			&t.UpdatedAt,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		result = append(result, t)
	}

	return result, nil
}

func (m *mysqlUserRepository) GetByEmail(ctx context.Context, title string) (res *models.User, err error) {
	list, err := m.fetch(ctx, user.QueryGetUserByEmail, title)
	if err != nil {
		logrus.Error(err)
		return res, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return nil, models.ErrNotFound
	}

	return res, nil
}
