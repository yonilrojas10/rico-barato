package roles

import (
	"time"

	"github.com/asaskevich/govalidator"
)

// Roles  Model struct Roles
type Roles struct {
	ID        string    `json:"id" db:"id" valid:"required,uuid"`
	Name      string    `json:"name" db:"name" valid:"required"`
	UserId    string    `json:"user_id" db:"user_id" valid:"required"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func NewRoles(id string, name string, userId string) *Roles {
	return &Roles{
		ID:     id,
		Name:   name,
		UserId: userId,
	}
}

func (m *Roles) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
