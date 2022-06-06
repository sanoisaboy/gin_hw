# Grades Inquiries System for teacher 
Grades Inquiries System for teacher is a backend system for teacher to register students grades



## Quick start

    $go run cmd/gin_hw/main.go --connect_string "<databaes_connect_string>"

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
