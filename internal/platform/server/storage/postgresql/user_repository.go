package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/huandu/go-sqlbuilder"
	domain "simci-go/internal"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Save(ctx context.Context, user domain.User) error {
	courseSQLStruct := sqlbuilder.NewStruct(new(sqlUser))
	query, args := courseSQLStruct.InsertInto(tableUsers, sqlUser{
		ID:       user.ID(),
		Name:     user.Name(),
		Email:    user.Email(),
		Password: "", // TODO: change for encrypt password
	}).Build()

	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("error trying to persist user on database: %v", err)
	}

	return nil
}
