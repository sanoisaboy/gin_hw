# Grades Inquiries System for teacher 
Grades Inquiries System for teacher is a backend system for teacher base on Gin to register students grades.

It provides five kind of method to controll the student grades with postman,respectively are ListStudent, GetStudent, CreateStudent, UpdateStudent, DeleteStudent,and I choose CockroachDB to be the database in this module,and it also five kind of method to controll the database, respectively are ListStudent, GetStudent, CreateStudent, UpdateStudent, DeleteStudent in crdb_repository.go file


## Schema design 
This system only contain one table, three field, respectively are student_name, id, point,and student_name is string, id is primary key,point is int.
![alt text](file:///C:/Users/sanoisaboy/Desktop/student_grade_schema.PNG)

## Quick start

    $go run cmd/gin_hw/main.go --connect_string "<databaes_connect_string>"
and use browser to connect to localhost:8002.

## Pre-require
* Go
* CockroachDB

## Module 1: Distributed SQL
In this module, we need to familiar with Distributed SQL (a kind of SQL database that design for scale, maybe base on Google Spanner or Calvin/SLOG papers). I choose CockroachDB to practice with this module.

### Objectives
* Admin CockroachDB (eg: setup database and user-schema...)
* Design application schema (eg: tables, indexes...)

## Module 2: Repository & Entity & Interactor
In this module, we need to use Go to transfer SQL command into Go code. You may use some database driver or ORM libraries like GORM or pgxpool.

### Objectives
* Implement ```crdbRepository struct``` and ```Repository interface``` 

## Module 3: Protobuf & gRPC & HTTP
In this module, we need to build a client interface, in other words, we need to provide HTTP endpoints . We can use GIN.
* Implement GIN
* Implemment HTTP endpoint

## How to use
first run the code.

    $go run cmd/gin_hw/main.go -connect_string "<database_connect_string>"

and use browser to connect to localhost:8002.
