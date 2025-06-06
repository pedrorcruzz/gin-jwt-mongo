# Gin + JWT + MongoDB Starter Project

## 📌 Status

### 🚧 in developing, but it is currently on hold. I will resume it soon.


This project was created to test and explore the main features of the [Gin](https://gin-gonic.com/) web framework in combination with [JWT](https://jwt.io/) for authentication and [MongoDB](https://www.mongodb.com/) as the database.

## 🧱 Architecture

This project follows the **Layered Architecture** pattern, separating concerns into different layers such as `handler`, `service`, `repository`, `entities`, and `utils`. This structure improves code organization, maintainability, and scalability as the project grows.

## 🔧 Technologies Used

- [Gin](https://github.com/gin-gonic/gin) – A high-performance HTTP web framework written in Go.
- [JWT](https://github.com/golang-jwt/jwt) – JSON Web Tokens for secure authentication.
- [MongoDB](https://www.mongodb.com/) – NoSQL database for data storage.
- [Go](https://golang.org/) – Programming language used in the project.

## 🚀 Features

- User registration and login with JWT authentication
- Protected routes using middleware
- Integration with MongoDB to persist user data
- Clean and modular project structure for scalability



## Run Project

### Run
```
go run ./cmd/api
```
### Build

```
go build ./cmd/api
```
