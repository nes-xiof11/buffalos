package domain

import (
	"errors"
	"strings"
	"time"
)

type User struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Validate checks whether the User struct contains valid and required data.
// It returns an error describing what is invalid or missing.
func (u *User) Validate() error {
	// Username is required
	if strings.TrimSpace(u.Username) == "" {
		return errors.New("username is required")
	}

	// Email is required
	if strings.TrimSpace(u.Email) == "" {
		return errors.New("email is required")
	}

	// Very basic email format check
	if !strings.Contains(u.Email, "@") || !strings.Contains(u.Email, ".") {
		return errors.New("invalid email format")
	}

	// Password required (only for users that are not yet created)
	if u.ID == 0 && strings.TrimSpace(u.Password) == "" {
		return errors.New("password is required")
	}

	// Role default check or validation rule
	if u.Role == "" {
		u.Role = "user"
	}
	//if !u.CreatedAt.IsZero() && u.UpdatedAt.Before(u.CreatedAt) {
	//	return errors.New("UpdatedAt cannot be before CreatedAt")
	//}

	return nil
}
