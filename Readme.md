# User Management Golang Application

## Overview

This is a User Management application developed in Go. It provides a set of APIs for managing user data. The application runs locally and connects to a MongoDB database. It also includes Swagger documentation for easy exploration of the APIs.


## Run Locally

Clone the project

```bash
  git clone https://github.com/VaishaliBavche/GolangCourse.git
```

Go to the project directory

```bash
  cd GolangCourse
```

Import dependencies

```bash
  go mod tidy
```

Start the server

```bash
  go run .
```


## API Reference

#### Get all Users

```http
  GET /users
```
gets list of all the users to count of users present.

#### Get User by Id

```http
  GET /users/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of user to fetch |

gets user by provided id.

#### Delete User by Id

```http
  DELETE /users/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of user to delete|

deletes user by provided id.

#### Create new user

```http
  POST /users
```
Payload
```json
{
    "name": "string",       // required
    "email": "string",      // required
    "email": "string",
    "age": integer,
    "isActive": boolean,
    "type": "string"
}
```
create new user with provided payload and return the id of user.
#### Update User by Id

```http
  PATCH /users/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of user to update|


Payload
```json
{
    "name": "string",       // required
    "email": "string",      // required
    "age": integer,
    "isActive": boolean,
    "type": "string"
}
```
update the user details by provided id and payload.



## API Endpoints
Swagger Documentation

Swagger UI: http://localhost:3000/swagger/index.html
## Authors

- [VaishaliBavche](https://www.github.com/VaishaliBavche)


## 🛠 Tech Stacks
Golang, Mongodb, Swagger Apis, Echo Framework etc.


## Support

For support, email us.
- [Vaishali Bavche](mailto:vbavche71198@gmail.com )