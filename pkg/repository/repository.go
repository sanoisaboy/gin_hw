package repository

import "context"

type Repository interface {
	ListStudent(ctx context.Context, id int) ([]User, error)
	GetStudent(ctx context.Context) ([]User, error)
	CreateStudent(ctx context.Context, name string, id int, point int, http_status int) (int error)
	UpdateStudent(ctx context.Context, id int, point int, http_status int) (int, []User, error)
	DeleteStudent(ctx context.Context, id int) error
}
