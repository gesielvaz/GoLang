package repository

import (
	"context"

	"github.com/disturb/inventory/internal/entity"
)

const (
	qryInserUser = `
		insert into USERS (email,name,password)
		values (?,?,?);
	`
	qryGetUserByEmail = `
	select 
		id,
		email,
		name,
		password
		from USERS;
		where email = ?;
	`
)

func (r *repo) SaveUser(ctx context.Context, email, name, password string) error {
	_, err := r.db.ExecContext(ctx, qryInserUser, email, name, password)
	return err
}
func (r *repo) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	u := &entity.User{}
	err := r.db.GetContext(ctx, u, qryGetUserByEmail, email)
	return u, err
}
