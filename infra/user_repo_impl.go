package infra

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/komugi8/clava/domain"
)

type UserRepositoryImpl struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) GetUsersByIDs(c context.Context, userIDs []domain.UserID) (*[]domain.User, error) {
	users := &[]domain.User{}
	query, args, err := sqlx.In(`SELECT * FROM user WHERE id IN (?)`, userIDs)
	if err != nil {
		return nil, err
	}
	query = r.db.Rebind(query)
	err = r.db.Select(users, query, args...)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepositoryImpl) GetUserByID(c context.Context, userID domain.UserID) (*domain.User, error) {
	user := &domain.User{}
	query := `SELECT * FROM user WHERE id = ?`
	err := r.db.Get(&user, query, userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepositoryImpl) CreateUser(c context.Context, user *domain.User) (*domain.User, error) {
	query := `INSERT INTO user (name, password, role) VALUES (?, ?, ?)`
	_, err := r.db.Exec(query, user.Name, user.Password, user.Role)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepositoryImpl) DeleteUser(c context.Context, userID domain.UserID) (*domain.User, error) {
	user, err := r.GetUserByID(c, userID)
	if err != nil {
		return nil, err
	}
	query := `DELETE FROM user WHERE id = ?`
	_, err = r.db.Exec(query, userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepositoryImpl) UpdateUser(c context.Context, user *domain.User) (*domain.User, error) {
	query := `UPDATE user SET name = ?, password = ?, role = ? WHERE id = ?`
	_, err := r.db.Exec(query, user.Name, user.Password, user.Role, user.ID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
