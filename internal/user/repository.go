package user

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *repository {
	return &repository{
		db: db,
	}
}

func (r repository) Save(user User) error {

	query := `

	INSERT INTO users VALUES (id, name, email, password, joined_at) ($1, $2, $3, $4, $5)
	`

	_, err := r.db.Exec(context.Background(), query, user.Id, user.Name, user.Email, user.Password, user.JoinedAt)

	return err
}

func (r repository) Delete(userId string) error {

	query := `
		DELETE FROM users WHERE id = $1
	`
	_, err := r.db.Exec(context.Background(), query, userId)

	return err

}
