package services

import (
	"buffalos/src/internal/domain"
	"buffalos/src/internal/misc"
	repositories "buffalos/src/internal/repositories"
	"context"
	"errors"
	"log"
)

type User struct {
	Repo *repositories.UserPG
}

func NewUser(repo *repositories.UserPG) *User {
	return &User{Repo: repo}
}

// Create validates the incoming user data and attempts to persist it in the repository.
//
// Return values:
//
//   - bool: indicates whether validation was successful.
//
//   - false → validation failed (treated as a Bad Request).
//
//   - true  → validation passed (but an error may still occur when accessing the repository).
//
//   - error:
//
//   - returned when validation fails (Bad Request).
//
//   - returned when the repository operation fails (Internal Server Error).
//
// Behavior:
//   - Ensures username and password are provided.
//   - Hashes the password before persistence.
//   - Assigns a default role ("user") when not provided.
func (us *User) Create(ctx context.Context, dto *domain.User) (bool, error) {
	if err := dto.Validate(); err != nil {
		return false, err
	}

	dto.Password = misc.SHA256(dto.Password)
	return true, us.Repo.Create(ctx, dto)
}

func (us *User) Login(ctx context.Context, dto *domain.User) (bool, *domain.User, error) {
	user, err := us.Repo.GetByEmail(ctx, dto.Email)
	if err != nil || user == nil {
		log.Printf("error getting user by email: %v", err)
		return false, nil, errors.New("user not found")
	}

	if user.Password == misc.SHA256(dto.Password) {
		return true, user, nil
	}
	return false, user, nil
}

func (us *User) GetByID(ctx context.Context, id int64) (*domain.User, error) {
	user, err := us.Repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *User) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
	user, err := us.Repo.GetByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *User) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	user, err := us.Repo.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *User) GetAll(ctx context.Context, limit, offset int64) ([]domain.User, error) {
	users, err := us.Repo.GetAll(ctx, limit, offset)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (us *User) Update(ctx context.Context, dto *domain.User) (bool, error) {
	if err := dto.Validate(); err != nil {
		return false, err
	}
	dto.Password = misc.SHA256(dto.Password)
	return true, us.Repo.Update(ctx, dto)
}

func (us *User) Delete(ctx context.Context, dto *domain.User) (bool, error) {
	//if err := dto.Validate(); err != nil {
	//	return false, err
	//}
	return true, us.Repo.Delete(ctx, dto.ID)
}
