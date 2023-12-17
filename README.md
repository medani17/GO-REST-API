Hello, I have created a simple Go project that uses Gofr, Gorm, and MySQL.

This project aims to create REST APIs for basic operations, including create, read, update, and delete. I have create simple example involving a student to illustrate the use case for creation.

This project can be consider for future modification for learning purpose where i can add more feature for a student such as subjects a student can have.

## To test this project i have created a database locally using mysql cli with the following commands:

1. create database students;

2. use database students;

3. CREATE TABLE students (
    id INT PRIMARY KEY,
    name VARCHAR(255),
    class VARCHAR(50),
    age INT
   );

4. After that i have used create API and generated students records.

## 🙇 Application Requirement

1.  Install Go version - 1.21.5 via 
 * https://go.dev/doc/install

2. Once installed, Please check via below command
```bash   
    go version
```
3. Install GoFr
```bash   
    go get gofr.dev

## ⚙️Project-Structure and Setup
``` 
    HTTP-Rest-API/
        ├── configs/
        │     └── .env
        │
        ├── handler/
        │     └── handler.go
        │
        ├── model/
        │     └── model.go
        │
        ├── datastore/
        │     ├── datastore.go
        │     └── interface.go
        │
        ├── main.go
        ├── go.mod

```

After our project is succesfully built we have run it with the command go run main.go
and it then pops the result of successfully connected to the database.(you can check the image [run ss.jpeg]) 

##Test

Then we have used postman and tested different apis that we have made, and got the output similar to the attached images below: 

