# HTTP-Rest-API Using Go & GoFr

Hello and welcome to the world of Go (also known as Golang), GORM, and the innovative GoFr framework. In this Project, we’ll take you on a practical journey to build a REST API using Go, combined with GORM for handling databases, and powered by GoFr Framework.
Before getting started get’s get to know them one by one:

### Go(Golang) :- 
Go, the renowned programming language, provides a solid foundation for building web applications, microservices, and network-intensive tasks with its elegance, simplicity.

### GORM :- 
GORM, the Object-Relational Mapping (ORM) framework, acts as a bridge between Go objects and relational databases.

### GoFr :- 
GoFr is Opinionated Web Framework written in Go (Golang). It helps in building robust and scalable applications. It makes creating RESTful APIs and web applications a breeze, even if you’re relatively new to Go. It prioritise simplicity over complexity.



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
