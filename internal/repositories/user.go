package repositories

import (
	"context"
	"fmt"
	"strings"

	"github.com/dangnp/go-backend-api/internal/entities"

	"github.com/jackc/pgx/v4/pgxpool"
)

type UserRepo struct{}

type UserFilter struct {
	ID    int
	Email string
}

func (r *UserRepo) FindUserFilter(ctx context.Context, db *pgxpool.Pool, filter UserFilter) (*entities.User, error) {

	user := &entities.User{}
	fields, _ := user.FieldMap()
	var cond string
	var val interface{}

	if filter.ID != 0 {
		cond = fmt.Sprintf("id = 1$")
		val = filter.ID
	}

	if filter.Email != "" {
		cond = fmt.Sprintf("email = 1$")
		val = filter.Email
	}

	sql := fmt.Sprintf(`SELECT %s FROM "%s" WHERE %s`, strings.Join(fields, ","), user.TableName(), cond)
	fmt.Println("SQL: ", sql)

	rows, err := db.Query(ctx, sql, val)
	defer rows.Close()

	if err != nil {
		return &entities.User{}, err
	}

	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Email, &user.Password, &user.Status, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			fmt.Errorf("Cannot load user.")
		}
	}
	return user, nil
}
