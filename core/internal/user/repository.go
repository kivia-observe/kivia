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
	INSERT INTO users (id, name, email, password, profile_picture, provider, joined_at) VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := r.db.Exec(context.Background(), query, user.Id, user.Name, user.Email, user.Password, user.ProfilePicture, user.Provider, user.JoinedAt)

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
		SELECT id, name, email, COALESCE(password, ''), profile_picture, COALESCE(provider, 'local'), joined_at FROM users
		WHERE email = $1
	`

	row := r.db.QueryRow(context.Background(), query, email)

	row.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.ProfilePicture, &user.Provider, &user.JoinedAt)

	return user
}

func (r Repository) FindById(id string) User {

	var user User

	query := `
		SELECT id, name, email, COALESCE(password, ''), profile_picture, COALESCE(provider, 'local'), joined_at FROM users
		WHERE id = $1
	`

	row := r.db.QueryRow(context.Background(), query, id)
	row.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.ProfilePicture, &user.Provider, &user.JoinedAt)

	return user
}

func (r Repository) DeleteById(id string) error {

	query := `
		DELETE FROM users WHERE id = $1
	`

	tag, err := r.db.Exec(context.Background(), query, id)
	
	if tag.RowsAffected() == 0 {
		return UserNotFound
	}

	return err
}

func (r Repository) UpdateById(id string, user editUserRequest) error {

	query := `
		UPDATE users SET name = $1, email = $2 WHERE id = $3
	`

	tag, err := r.db.Exec(context.Background(), query, user.Name, user.Email, id)
	
	if tag.RowsAffected() == 0 {
		return UserNotFound
	}

	return err
}