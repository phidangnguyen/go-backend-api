package entities

import "github.com/jackc/pgtype"

type User struct {
	ID pgtype.Int8 `json:"id"`
	//
	Email    pgtype.Text `json:"email"`
	Password pgtype.Text `json:"password,omitempty"`
	Status   pgtype.Text `json:"status,omitempty"`

	CreatedAt pgtype.Int8 `json:"created_at"`
	UpdatedAt pgtype.Int8 `json:"updated_at"`
}

func (e *User) TableName() string {
	return "user"
}

func (e *User) FieldMap() ([]string, []interface{}) {
	return []string{
			"id",
			"email",
			"password",
			"status",
			"created_at",
			"updated_at",
		}, []interface{}{
			&e.ID,
			&e.Email,
			&e.Password,
			&e.Status,
			&e.CreatedAt,
			&e.UpdatedAt,
		}
}
