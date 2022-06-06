# Grades Inquiries System for teacher 
Grades Inquiries System for teacher is a backend system for teacher base on Gin to register students grades.

It provides five kind of method to controll the student grades with postman,respectively are ListStudent, GetStudent, CreateStudent, UpdateStudent, DeleteStudent,and I choose CockroachDB to be the database in this module,and it also five kind of method to controll the database, respectively are ListStudent, GetStudent, CreateStudent, UpdateStudent, DeleteStudent in crdb_repository.go file


## Schema design 
This system only contain one table, three field, respectively are student_name, id, point, and student_name is string and unquie, id is int and primary key, point is int.

![alt text](https://github.com/sanoisaboy/gin_hw/blob/main/student_grade_schema.PNG)

## API use
### ListStudent
use postman to GET localhost:8002/v1/users and you will get all the students grades
![alt text](https://github.com/sanoisaboy/gin_hw/blob/main/555319.png)

### GetStudent
use postman to GET localhost:8002/v1/users/id?id=<student_id> and you will get the designated student name, id , point.
![alt text](https://github.com/sanoisaboy/gin_hw/blob/main/16216348053558.png)

### CreateStudent
use postman to POST localhost:8002/v1/users?studnet_name=<student_name>&id=<student_id>&point=<student_point> and will create a new student grades in the database.
![alt text](https://github.com/sanoisaboy/gin_hw/blob/main/16216352730447.png)

### UpdateStudent
use postman to PATCH localhost:8002/v1/users?id=<student_id>&point=<student_update_point> and will update the designated student point.
![alt text](https://github.com/sanoisaboy/gin_hw/blob/main/16216355519108.png)

### DeleteStudent
use postman to DELETE localhost:8002/v1/users?id=<student_id> and will delete the designated student in the database.
![alt text](https://github.com/sanoisaboy/gin_hw/blob/main/16216357563794.png)

## Quick start

    $go run cmd/gin_hw/main.go --connect_string "<databaes_connect_string>"
and use browser to connect to localhost:8002.


