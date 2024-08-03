package db

import (
	"context"
	"database/sql"
	"log/slog"
	"strings"
)

// idempotent save
// accepts ALL fields of entity and save as is
func (user *User) Save(ctx context.Context) error {
	_, _, _ = user.RefresTimestamps()

	_, err := sqliteConn.ExecContext(
		ctx,
		`INSERT INTO user(id, tgid, name, tgusername, chatid, createdat, updatedat)
        VALUES($1, $2, $3, $4, $5, $6, $7)
        ON CONFLICT(tgid) DO UPDATE SET name=$3, tgusername=$4, chatid=$5, updatedat=$6
        RETURNING id;`,
		&user.ID,
		&user.TGId,
		&user.Name,
		&user.TGusername,
		&user.ChatId,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		slog.Error("Error when trying to save user: " + err.Error())
		return err
	}
	slog.Debug("User created/updated")

	return nil
}

func (user *User) Filter(ctx context.Context) ([]User, error) {
	where := []string{}
	if user.TGId != 0 {
		where = append(where, "tgid=$tgid")
	}
	if user.TGusername != "" {
		where = append(where, "tgusername=$tgusername")
	}

	where_ := strings.Join(where, " AND ")
	query := `SELECT id, tgid, name, tgusername, chatid, createdat, updatedat FROM user WHERE ` + where_ + `;`

	rows, err := sqliteConn.QueryContext(
		ctx,
		query,
		sql.Named("tgid", user.TGId),
		sql.Named("tgusername", user.TGusername),
	)
	if err != nil {
		slog.Error("Error when filtering users " + err.Error())
		return nil, err
	}
	defer rows.Close()

	users := []User{}

	for rows.Next() {
		user := User{}
		err := rows.Scan(
			&user.ID,
			&user.TGId,
			&user.Name,
			&user.TGusername,
			&user.ChatId,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			slog.Error("Error fetching users by filter params: " + err.Error())
			continue
		}
		users = append(users, user)
	}

	return users, nil
}
