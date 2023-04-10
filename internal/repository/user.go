package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/caviarman/shop_caviar/internal/entity"
	"github.com/doug-martin/goqu/v9"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgtype/zeronull"
	"github.com/jackc/pgx/v4"
)

func (r *Repository) GetUsers(ctx context.Context) ([]*entity.User, error) {
	rows, err := r.pool.Query(ctx, `SELECT 
		"id",
		"name",
		"email",
		"phone",
		"telegram_username",
		"telegram_id",
		"created_at",
		"updated_at" 
	FROM "user";`)
	if err != nil && err != pgx.ErrNoRows {
		return nil, fmt.Errorf("r.pool.Query: %w", err)
	}

	var users []*entity.User

	for rows.Next() {
		var (
			user      entity.User
			name      pgtype.Text
			email     pgtype.Text
			phone     pgtype.Text
			tusername pgtype.Text
			cr        pgtype.Timestamp
			up        pgtype.Timestamp
		)

		err := rows.Scan(
			&user.ID,
			&name,
			&email,
			&phone,
			&tusername,
			&user.TelegramID,
			&cr,
			&up,
		)
		if err != nil {
			return nil, fmt.Errorf("rows.Scan: %w", err)
		}

		user.Name = name.String
		user.Email = email.String
		user.Phone = phone.String
		user.TelegramUsername = tusername.String
		user.CreatedAt = cr.Time
		user.UpdatedAt = up.Time

		users = append(users, &user)
	}

	return users, nil
}

func (r *Repository) CreateUser(ctx context.Context, user *entity.User) (int, error) {
	if user == nil {
		return 0, nil
	}

	ds := goqu.Insert("user").Rows(goqu.Record{
		"name":              zeronull.Text(user.Name),
		"email":             zeronull.Text(user.Name),
		"phone":             zeronull.Text(user.Phone),
		"telegram_username": zeronull.Text(user.TelegramUsername),
		"telegram_id":       user.TelegramID,
		"created_at":        time.Now().UTC(),
	}).
		OnConflict(
			goqu.DoUpdate("telegram_id", goqu.Record{"updated_at": time.Now().UTC()})).
		Returning("id")

	sql, _, err := ds.ToSQL()
	if err != nil {
		return 0, fmt.Errorf("ds.ToSQL: %w", err)
	}

	var id int

	err = r.pool.QueryRow(ctx, sql).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("r.pool.QueryRow: %w", err)
	}

	return id, nil
}
