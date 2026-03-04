package user

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		db: db,
	}
}

func (r Repository) Save(user User) error {

	query := `

	INSERT INTO users (id, name, email, password, joined_at) VALUES ($1, $2, $3, $4, $5)
	`

	_, err := r.db.Exec(context.Background(), query, user.Id, user.Name, user.Email, user.Password, user.JoinedAt)

	return err
}

func (r Repository) Delete(userId string) error {

	query := `
		DELETE FROM users WHERE id = $1
	`
	_, err := r.db.Exec(context.Background(), query, userId)

	return err

}

func (r Repository) ExistsByEmail(email string) bool {

	var exists bool 

	query := `
		SELECT EXISTS (
			SELECT 1 FROM users
			WHERE email = $1
		)
	`

	row := r.db.QueryRow(context.Background(), query, email)

	row.Scan(&exists)

	return exists
}

func (r Repository) FindByEmail(email string) User {

	var user User 

	query := `
		SELECT id, name, email, password, joined_at FROM users
		WHERE email = $1
	`

	row := r.db.QueryRow(context.Background(), query, email)

	row.Scan(&user.Id, &user.Name, &user.Email,&user.Password, &user.JoinedAt)

	return user
}

func (r Repository) FindById(id string) User {

	var user User 

	query := `
		SELECT id, name, email, password, joined_at FROM users
		WHERE id = $1
	`

	row := r.db.QueryRow(context.Background(), query, id)
	row.Scan(&user.Id, &user.Name, &user.Email,&user.Password, &user.JoinedAt)

	return user
}
