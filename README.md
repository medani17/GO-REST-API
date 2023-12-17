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
```
## ⚙️Project-Structure and Setup

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
        │── test
        │    ├── datastore_test.go
        │    ├── handler_test.go
        │    ├── model_test.go
        |                 
        ├── main.go
        ├── go.mod

```
```
After our project is succesfully built we have run it with the command go run main.go
and it then pops the result of successfully connected to the database.
![run ss](https://github.com/medani17/GO-REST-API/assets/71040928/774a03d0-d9b4-42fc-a502-41b73a034b81)


## Unit Test
Then we have used postman and tested different apis that we have made, and got the output similar to the attached images below: 
![create](https://github.com/medani17/GO-REST-API/assets/71040928/4662d0d8-a7d5-4cff-ac94-a3d74e10fd97)
![delete ss](https://github.com/medani17/GO-REST-API/assets/71040928/7a21d2e1-c7ae-4dde-b4be-78e684f8a432)


## Diagrams
1. Architecture Diagram![Architecture Diagram](https://github.com/medani17/GO-REST-API/assets/71040928/5c49370b-db36-4d38-aacf-516ff8c54021)
2. Sequence Diagram![Sequence Diagram](https://github.com/medani17/GO-REST-API/assets/71040928/2899f84e-6348-443d-8907-2869c21b9135)






