package repository

import (
	"context"

	"github.com/Jenietor2/app-inventory/internal/entity"
)

const (
	qryInsertUser = `
		insert into USERS(email, username, password)
		values(?,?,?)
	`

	qryGetUserByEmail = `
		select 
			id,
			email,
			username,
			password
		from USERS
		where email=?
	`
)

func (r *repo) SaveUser(ctx context.Context, email, name, password string) error {
	_, err := r.db.ExecContext(ctx, qryInsertUser, email, name, password)
	return err
}

func (r *repo) GetUserByEmail(ctx context.Context, email string) (*entity.User, error){
	u := &entity.User{}
	err := r.db.GetContext(ctx, &u, qryGetUserByEmail, email)
	if err != nil{
		return nil, err
	}
	
	return u, nil
}