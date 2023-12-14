Hello, I have created a simple go project which user gofr, gorm, and mysql.

This project has a goal to create Rest APIs such as Create , update , delete and put . I have taken firmly small example student which has simple use case for creation .

This project can be consider for future modification for learning purpose where i can add more feature for a student such as subjects a student can have.

To test this project i have created a database locally using mysql cli with the following commands

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
4. Docker :
   To simplify the setup and management of the database for your REST API, consider using a Docker image for SQL. This approach provides a consistent and isolated environment for your       database, making it easier to manage and maintain. You can use the official MySQL Docker image from this link - https://hub.docker.com/_/mysql





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
