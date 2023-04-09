package uow

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type RepositoryFactory func(tx *sql.Tx) interface{}

type UowInterface interface {
	Register(name string, fc RepositoryFactory)
	GetRepository(ctx context.Context, name string) (interface{}, error)
	Do(ctx context.Context, fn func(uow UowInterface) error) error
	CommitOrRollback() error
	Rollback() error
	UnRegister(name string)
}

type uow struct {
	DB           *sql.DB
	TX           *sql.Tx
	Repositories map[string]RepositoryFactory
}

func NewUow(ctx context.Context, db *sql.DB) UowInterface {
	return &uow{
		DB:           db,
		Repositories: make(map[string]RepositoryFactory),
	}
}

func (u *uow) Register(name string, fc RepositoryFactory) {
	u.Repositories[name] = fc
}

func (u *uow) UnRegister(name string) {
	delete(u.Repositories, name)
}

func (u *uow) GetRepository(ctx context.Context, name string) (interface{}, error) {
	if u.TX == nil {
		tx, err := u.DB.BeginTx(ctx, nil)
		if err != nil {
			return nil, err
		}
		u.TX = tx
	}
	repository := u.Repositories[name](u.TX)
	return repository, nil
}

func (u *uow) Do(ctx context.Context, fn func(uow UowInterface) error) error {
	if u.TX != nil {
		return errors.New("transaction already started")
	}

	tx, err := u.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	u.TX = tx

	if err = fn(u); err != nil {
		if errRollback := u.Rollback(); errRollback != nil {
			return fmt.Errorf("original error: %s, rollback error: %s", err, errRollback)
		}
		return err
	}
	return u.CommitOrRollback()
}

func (u *uow) CommitOrRollback() error {
	if err := u.TX.Commit(); err != nil {
		if errRollback := u.Rollback(); errRollback != nil {
			return fmt.Errorf("original error: %s, rollback error: %s", err, errRollback)
		}
		return err
	}
	u.TX = nil
	return nil
}

func (u *uow) Rollback() error {
	if u.TX != nil {
		return errors.New("no transaction to rollback")
	}

	if err := u.TX.Rollback(); err != nil {
		return err
	}

	u.TX = nil
	return nil
}
