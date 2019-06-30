package repository

import (
	"context"
	"database/sql"
	"encoding/base64"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/febrycode/healthy_food/models"
	"github.com/febrycode/healthy_food/user"
)

type mysqlUserRepository struct {
	Conn *sql.DB
}

// NewMysqlUserRepository will create an object that represent the user.Repository interface
func NewMysqlUserRepository(Conn *sql.DB) user.Repository {
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
		authorID := int64(0)
		err = rows.Scan(
			&t.ID,
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
		t.Author = models.Author{
			ID: authorID,
		}
		result = append(result, t)
	}

	return result, nil
}

func (m *mysqlUserRepository) GetByEmail(ctx context.Context, title string) (res *models.Article, err error) {
	query := `SELECT id,title,content, author_id, updated_at, created_at
  						FROM article WHERE title = ?`

	list, err := m.fetch(ctx, query, title)
	if err != nil {
		return
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return nil, models.ErrNotFound
	}
	return
}

func (m *mysqlUserRepository) Store(ctx context.Context, a *models.Article) error {
	query := `INSERT  article SET title=? , content=? , author_id=?, updated_at=? , created_at=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(ctx, a.Title, a.Content, a.Author.ID, a.UpdatedAt, a.CreatedAt)
	if err != nil {
		return err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return err
	}

	a.ID = lastID
	return nil
}

// DecodeCursor will decode cursor from user for mysql
func DecodeCursor(encodedTime string) (time.Time, error) {
	byt, err := base64.StdEncoding.DecodeString(encodedTime)
	if err != nil {
		return time.Time{}, err
	}

	timeString := string(byt)
	t, err := time.Parse(timeFormat, timeString)

	return t, err
}

// EncodeCursor will encode cursor from mysql to user
func EncodeCursor(t time.Time) string {
	timeString := t.Format(timeFormat)

	return base64.StdEncoding.EncodeToString([]byte(timeString))
}
