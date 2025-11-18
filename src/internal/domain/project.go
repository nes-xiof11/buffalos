package domain

import (
	"errors"
	"strings"
	"time"
)

type Project struct {
	ID        int64     `json:"id"`
	OwnerID   int64     `json:"owner_id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (p *Project) Validate() error {
	if p.OwnerID <= 0 {
		return errors.New("OwnerID must be greater than zero")
	}

	if strings.TrimSpace(p.Name) == "" {
		return errors.New("name cannot be empty")
	}

	if len(p.Name) > 255 {
		return errors.New("name cannot exceed 255 characters")
	}

	//if !p.CreatedAt.IsZero() && p.UpdatedAt.Before(p.CreatedAt) {
	//	return errors.New("UpdatedAt cannot be before CreatedAt")
	//}

	return nil
}
