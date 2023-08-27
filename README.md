# Auth API

![Zoo Keeper Logo](img/../assets/img/golady.png)

> This Project is a basic Authorization API .

---

**Technologies used:**

- GO
- Gin
- Gorm
- JWT Authentication

---

**Prerequisites**

Make sure you have the following tools installed on your system:

- Go (https://golang.org/doc/install)
- MySQL (or another database of your choice)

**Configuration**

1. Clone this repository to your local environment.

2. Create a MySQL database for the application. For example, you can use the following MySQL command:

   ```sql
   CREATE DATABASE authentication_api;

   ```

3. Dplicate `.env.exemple` fill with your information.

4. Install the application dependencies using the following command:

- go mod tidy

5. Run the application using the following command:

- go run main.go

---

**Endpoints**

#### Registration (Register)

POST /register

Creates a new user in the database.

Parameters:

username: User's username
password: User's password

#### Login

POST /login

Authenticates the user and generates a JWT token.

Parameters:

username: User's username
password: User's password
The response will include a JWT token that should be used in subsequent calls requiring authentication.

#### JWT Authentication

This API uses JWT (JSON Web Tokens) for authentication to secure routes that require authentication. You should include the JWT token in the request header for protected routes.

```
Authorization: Bearer YOUR_JWT_TOKEN_HERE
```
