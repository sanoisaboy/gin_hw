package repository

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

type crdbRepository struct {
	pool *pgxpool.Pool
}

type User struct {
	Student_name string
	ID           int
	POINT        int
}

func NewCrdbRepository(conn string) *crdbRepository {

	pool, err := pgxpool.Connect(context.Background(), conn)

	if err != nil {
		log.Fatal("Unable to connect database: ", err)
	}
	return &crdbRepository{
		pool: pool,
	}
}
func (r *crdbRepository) GetStudent(ctx context.Context, id int) ([]User, error) {
	student_get := []User{}

	rows, _ := r.pool.Query(context.Background(),
		"SELECT * FROM student WHERE id=$1", id)

	for rows.Next() {
		var connect string
		var point int
		var id int
		err := rows.Scan(&connect, &id, &point)
		if err != nil {
			log.Fatalln(err)
		}
		student_get = append(student_get, User{Student_name: connect, ID: id, POINT: point})

	}
	log.Printf("%+v", student_get)

	return student_get, nil
}

func (r *crdbRepository) ListStudent(ctx context.Context) ([]User, error) {
	student_get := []User{}

	rows, _ := r.pool.Query(context.Background(),
		"SELECT * FROM student")

	for rows.Next() {
		var connect string
		var point int
		var id int
		err := rows.Scan(&connect, &id, &point)
		if err != nil {
			log.Fatalln(err)
		}
		student_get = append(student_get, User{Student_name: connect, ID: id, POINT: point})
	}
	log.Printf("%+v", student_get)

	return student_get, nil
}

func (r *crdbRepository) CreateStudent(ctx context.Context, name string, id int, point int, http_status int) (int, error) {
	if http_status == 403 {
		return 403, nil
	}

	if _, err := r.pool.Exec(context.Background(),
		"INSERT INTO student (student_name,id, point) VALUES ($1,$2, $3)", name, id, point); err != nil {
		log.Println(err)
		return 403, err
	}
	return 201, nil
}

func (r *crdbRepository) UpdateStudent(ctx context.Context, id int, point int, http_status int) (int, []User, error) {
	if http_status == 400 {
		return 400, nil, nil
	}

	if _, err := r.pool.Exec(context.Background(),
		"UPDATE student SET point = $1 WHERE id = $2", point, id); err != nil {
		log.Println(err)
		return 400, nil, err
	}
	student_get := []User{}

	rows, _ := r.pool.Query(context.Background(),
		"SELECT * FROM student WHERE id=$1", id)

	for rows.Next() {
		var connect string
		var point int
		var id int
		err := rows.Scan(&connect, &id, &point)
		if err != nil {
			log.Fatalln(err)
		}
		student_get = append(student_get, User{Student_name: connect, ID: id, POINT: point})
	}
	log.Printf("%+v", student_get)

	return 200, student_get, nil

}

func (r *crdbRepository) DeleteStudent(ctx context.Context, id int) (int, error) {
	if _, err := r.pool.Exec(context.Background(),
		"DELETE FROM student WHERE id IN ($1)", id); err != nil {
		log.Println(err)
		return 400, err
	}
	return 204, nil
}
