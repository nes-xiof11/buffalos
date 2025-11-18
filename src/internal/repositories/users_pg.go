package repos

import (
	"buffalos/src/internal/domain"
	"context"
	"database/sql"
)

type UserPG struct {
	DB *sql.DB
}

func NewUserPG(db *sql.DB) *UserPG {
	return &UserPG{DB: db}
}

func (r *UserPG) Create(ctx context.Context, dto *domain.User) error {
	query := `
		INSERT INTO users (username, email, password, role) 
		VALUES ($1, $2, $3)
		RETURNING id, created_at, updated_at;
	`
	err := r.DB.QueryRowContext(ctx, query,
		dto.Username,
		dto.Email,
		dto.Password,
		dto.Role,
	).Scan(&dto.ID, &dto.CreatedAt, &dto.UpdatedAt)
	return err
}

func (r *UserPG) GetByID(ctx context.Context, id int64) (*domain.User, error) {
	dtos, err := r.fetchAll(ctx, `SELECT * FROM users WHERE id = $1;`, id)
	if err != nil || len(dtos) < 1 {
		return nil, err
	}

	return &dtos[0], nil
}

func (r *UserPG) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
	dtos, err := r.fetchAll(ctx, `SELECT * FROM users WHERE username = $1;`, username)
	if err != nil || len(dtos) < 1 {
		return nil, err
	}

	return &dtos[0], nil
}

func (r *UserPG) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	dtos, err := r.fetchAll(ctx, `SELECT * FROM users WHERE email = $1;`, email)
	if err != nil || len(dtos) < 1 {
		return nil, err

	}

	return &dtos[0], nil
}

func (r *UserPG) GetAll(ctx context.Context, limit, offset int64) ([]domain.User, error) {
	dtos, err := r.fetchAll(ctx,
		`SELECT * FROM users LIMIT $1 OFFSET $2;`,
		limit, offset,
	)

	if err != nil || len(dtos) < 1 {
		return nil, err
	}

	return dtos, nil
}

func (r *UserPG) Update(ctx context.Context, dto *domain.User) error {
	//
	return nil
}

func (r *UserPG) Delete(ctx context.Context, id int64) error {
	rows := r.DB.QueryRowContext(ctx, "DELETE FROM users WHERE id = $1;", id)

	if err := rows.Err(); err != nil {
		return err
	}
	return nil
}

func (r *UserPG) fetchAll(ctx context.Context, query string, args ...interface{}) ([]domain.User, error) {
	rows, err := r.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var dtos []domain.User
	for rows.Next() {
		var dto domain.User
		err = rows.Scan(
			&dto.ID,
			&dto.Username,
			&dto.Email,
			&dto.Password,
			&dto.Role,
			&dto.CreatedAt,
			&dto.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		dtos = append(dtos, dto)
	}
	return dtos, nil
}
