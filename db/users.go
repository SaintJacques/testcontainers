package db

import (
	"context"

	"testcntrns/models"

	sq "github.com/Masterminds/squirrel"
)

func (db *DB) InsertUser(ctx context.Context, name string) (*models.User, error) {
	user := models.User{}
	q, args, err := sq.Insert(user.TableName()).Columns("name").Values(name).Suffix("RETURNING id, name, created_at").PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	err = db.QueryRow(ctx, q, args...).Scan(&user.ID, &user.Name, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (db *DB) SelectUser(ctx context.Context, id int32) (*models.User, error) {
	user := models.User{}
	q, args, err := sq.Select("*").From(user.TableName()).Where(sq.Eq{"id": id}).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	err = db.QueryRow(ctx, q, args...).Scan(&user.ID, &user.Name, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
