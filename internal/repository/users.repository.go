package repository

import (
	"context"

	"github.com/disturb/inventory/internal/entity"
)

/* const(
	qryInserUser = `
	insert into USERS (email,name,password)
	values (?,?,?));
	`
) */

func (r *repo) SaveUser(ctx context.Context, email, name, password string) error {
	return nil
}
func (r *repo) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	return nil, nil
}
