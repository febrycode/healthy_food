package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"

	"github.com/febrycode/healthy_food/models"
	"github.com/febrycode/healthy_food/user"
)

type mysqlUserRepository struct {
	DB *sqlx.DB
}

// NewMysqlUserRepository will create an object that represent the user.Repository interface
func NewMysqlUserRepository(DB *sqlx.DB) user.Repository {
	return &mysqlUserRepository{DB}
}

func (m *mysqlUserRepository) GetByEmail(ctx context.Context, email string) (res models.User, err error) {
	err = m.DB.GetContext(ctx, &res, user.QueryGetUserByEmail, email)
	if err != nil {
		logrus.Error(err)
		return res, err
	}

	return res, nil
}

func (m *mysqlUserRepository) CreateUser(ctx context.Context, userData *models.User) error {
	_, err := m.DB.NamedQuery(user.QueryInsertUser, &userData)
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}
