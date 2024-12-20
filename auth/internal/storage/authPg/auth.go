package authPg

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/kechdarho/FinTrack/auth/internal/models"
)

func (s *Storage) CreateUser(ctx context.Context, email, username, password string) (userID uint, err error) {
	query := `INSERT INTO public.users (email, username, password_hash)
			VALUES ($1,$2,$3)
			RETURNING user_id`

	err = s.db.QueryRow(ctx, query, email, username, password).Scan(&userID)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.ConstraintName {
			case "users_email_key":
				return 0, fmt.Errorf("email %s already exists", email)
			case "users_username_key":
				return 0, fmt.Errorf("username %s already exists", username)
			default:
				return 0, fmt.Errorf("username %s already exists", pgErr.Error())
			}
		}
		return
	}
	return
}

func (s *Storage) GetUser(ctx context.Context, login string) (user models.User, err error) {
	query := `SELECT 
				user_id,
				password_hash,
				role
			 FROM public.users
			 WHERE email = $1 or username = $1`
	err = s.db.QueryRow(ctx, query, login).Scan(&user.UserID, &user.Password, &user.Role)
	if err != nil {
		return
	}

	return
}
