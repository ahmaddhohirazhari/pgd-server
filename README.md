<h1 align="center">pdg-server</h1>
<p align="center"><img src="https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Go_Logo_Blue.svg/2560px-Go_Logo_Blue.svg.png" width="400px" alt="Golang.jpg" /></p>
<p align="center">
    <a href="https://golang.org/" target="blank">More about Golang</a>
</p>

## 🔗 Description

This Backend Application is used for vehicle rental systems such as car rental, motorbikes, and bicycles. In the application, users can add, change, delete, and read the data of the vehicle they want to rent. In addition, users can also see the rental history. This application was built using the Golang programming language with the Gin Framework and uses GORM, a Database that is used using PostgreSQL and deployed on the Heroku website.

## 🔗 Installation Gorilla/Mux

- Install Gin

```sh
  go get -u github.com/gin-gonic/gin
```

## 🔗 Feature

- CRUD customer

## 🔗 Installation Step

- Go to the project directory

```sh
  mkdir pgd-server
  cd pgd-server

  go mod init pgd-server
  # add file main.go
```

- Clone the project

- Add Env

```sh
  APP_PORT= Your Port
  JWT_KEYS= Your Secret Keys

  DB_USER = Your DB User
  DB_HOST = Your DB Host
  DB_NAME = Your DB Name
  DB_PASS = Your DB Password
```

- Install dependencies

```sh
  go get -u ./..
  # or
  go mod tidy
```

- Start the server

```sh
  go run main.go
```

## 💻 Built with

- [Golang](https://go.dev/): Programming
- [Gin](https://github.com/gin-gonic/gin): for handle http request
- [Postgres](https://www.postgresql.org/): for DBMS

## 💻 Deploy

Link Deploy :

## 💻 End Point

Postmant Documentation :

- Postman : [link](https://documenter.getpostman.com/view/25631499/2sA3kYgyeE)

## 🚀 About Me

- Github : [ahmaddhohirazhari](https://github.com/ahmaddhohirazhari/)
