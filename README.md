# Go Authentication and Authorization Service

A RESTful authentication and user management API built with Go, Gin, MongoDB, and JSON Web Tokens (JWT). This service provides secure password hashing, JWT-based authentication, role-based access control (USER and ADMIN), and user pagination.

---

## Features

- User registration and login with bcrypt password hashing
- Stateless JWT authentication with access and refresh tokens
- Role-based authorization (USER and ADMIN roles)
- Token-authenticated route protection via Gin middleware
- Paginated user list aggregation
- MongoDB data persistence using official MongoDB Go Driver v2

---

## Project Structure

```
.
|-- controllers/
|   `-- userController.go     # Request handlers for signup, login, user queries
|-- database/
|   `-- databaseConnection.go # MongoDB client initialization and collection helper
|-- helpers/
|   |-- authHelper.go         # Role and user permission validation
|   `-- tokenHelper.go        # JWT creation, claim parsing, and validation
|-- middleware/
|   `-- authMiddleware.go     # Gin middleware for validating request token header
|-- models/
|   `-- userModel.go          # User struct definition and validation tags
|-- routes/
|   |-- authRouter.go         # Public authentication routes (/signup, /login)
|   `-- userRouter.go         # Protected user routes (/users, /users/:user_id)
|-- main.go                   # Application entry point and router setup
|-- go.mod
|-- go.sum
`-- .env
```

---

## Prerequisites

Ensure you have the following installed on your system:

- Go (version 1.20 or later)
- MongoDB instance (local or MongoDB Atlas cluster)
- Git

---

## Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/JeetDas5/go-auth.git
cd go-auth
```

### 2. Configure Environment Variables

Create a `.env` file in the root directory:

```env
PORT=8080
MONGODB_URL=mongodb://localhost:27017/go-auth
SECRET_KEY=your_jwt_secret_key_here
```

| Variable | Description | Default |
| --- | --- | --- |
| `PORT` | Port number on which the HTTP server listens | `8080` |
| `MONGODB_URL` | MongoDB connection URI string | Required |
| `SECRET_KEY` | Secret key used for signing and verifying JWT tokens | Required |

### 3. Install Dependencies

```bash
go mod download
```

### 4. Run the Application

```bash
go run main.go
```

The server will initialize the MongoDB connection and listen on `http://localhost:8080` (or your configured `PORT`).

---

## API Documentation

Base URL: `http://localhost:8080`

### Public Endpoints

#### 1. Health Check Endpoints

- **`GET /api-1`**
  - Description: Basic service health ping.
  - Response:
    ```json
    {
      "success": "API 1 is running"
    }
    ```

- **`GET /api-2`**
  - Description: Secondary health ping.
  - Response:
    ```json
    {
      "success": "API 2 is running"
    }
    ```

#### 2. User Signup

- **Method**: `POST`
- **Endpoint**: `/signup`
- **Description**: Registers a new user with hashed password and generates initial JWT tokens.
- **Headers**:
  - `Content-Type: application/json`
- **Request Body**:
  ```json
  {
    "first_name": "Jeet",
    "last_name": "Das",
    "password": "password123",
    "email": "jeet@example.com",
    "phone": "9876543210",
    "user_type": "USER"
  }
  ```
  - `first_name`: string (2-50 characters, required)
  - `last_name`: string (2-50 characters, required)
  - `password`: string (6-100 characters, required)
  - `email`: string (valid email format, required, unique)
  - `phone`: string (minimum 10 digits, required, unique)
  - `user_type`: string (`USER` or `ADMIN`, required)
- **Response (200 OK)**:
  ```json
  {
    "success": "User created successfully",
    "user_id": "64d0f6..."
  }
  ```
- **Error Responses**:
  - `400 Bad Request`: Validation failure or duplicate email/phone.
  - `500 Internal Server Error`: Database insertion error.

#### 3. User Login

- **Method**: `POST`
- **Endpoint**: `/login`
- **Description**: Authenticates user credentials, refreshes tokens, and returns user details.
- **Headers**:
  - `Content-Type: application/json`
- **Request Body**:
  ```json
  {
    "email": "jeet@example.com",
    "password": "password123"
  }
  ```
- **Response (200 OK)**:
  ```json
  {
    "id": "64d0f6...",
    "first_name": "Jeet",
    "last_name": "Das",
    "password": "<bcrypt_hash>",
    "email": "jeet@example.com",
    "phone": "9876543210",
    "token": "<jwt_access_token>",
    "user_type": "USER",
    "refresh_token": "<jwt_refresh_token>",
    "created_at": "2026-09-06T14:30:00Z",
    "updated_at": "2026-09-06T14:35:00Z",
    "user_id": "64d0f6..."
  }
  ```
- **Error Responses**:
  - `400 Bad Request`: Invalid request JSON.
  - `404 Not Found`: No user found with the given email.
  - `401 Unauthorized`: Password verification failure.

---

### Protected Endpoints

All protected endpoints require the client to pass the JWT access token in the `token` HTTP header:

```http
token: <jwt_access_token>
```

#### 4. List Users (Paginated)

- **Method**: `GET`
- **Endpoint**: `/users`
- **Description**: Returns a paginated list of all users. Accessible by `ADMIN` users only.
- **Headers**:
  - `token: <jwt_access_token>`
- **Query Parameters**:
  - `recordPerPage` (optional, default: `10`): Number of records per page.
  - `page` (optional, default: `1`): Page number (1-indexed).
- **Example Request**:
  ```http
  GET /users?recordPerPage=5&page=1 HTTP/1.1
  Host: localhost:8080
  token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
  ```
- **Response (200 OK)**:
  ```json
  {
    "total_count": 25,
    "user_items": [
      {
        "id": "64d0f6...",
        "first_name": "Jeet",
        "last_name": "Das",
        "email": "jeet@example.com",
        "phone": "9876543210",
        "user_type": "USER",
        "user_id": "64d0f6...",
        "created_at": "2026-09-06T14:30:00Z",
        "updated_at": "2026-09-06T14:35:00Z"
      }
    ]
  }
  ```
- **Error Responses**:
  - `400 Bad Request`: Unauthorized access if user role is not `ADMIN`.
  - `401 Unauthorized`: Missing or invalid JWT token.
  - `500 Internal Server Error`: Aggregation failure.

#### 5. Get User by ID

- **Method**: `GET`
- **Endpoint**: `/users/:user_id`
- **Description**: Retrieves a single user record by their `user_id`. Accessible by `ADMIN` users or by `USER` users accessing their own `user_id`.
- **Headers**:
  - `token: <jwt_access_token>`
- **URL Parameters**:
  - `user_id`: The unique user identifier string.
- **Example Request**:
  ```http
  GET /users/64d0f682ab1234567890cdef HTTP/1.1
  Host: localhost:8080
  token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
  ```
- **Response (200 OK)**:
  ```json
  {
    "id": "64d0f682ab1234567890cdef",
    "first_name": "Jeet",
    "last_name": "Das",
    "password": "<bcrypt_hash>",
    "email": "jeet@example.com",
    "phone": "9876543210",
    "token": "<jwt_access_token>",
    "user_type": "USER",
    "refresh_token": "<jwt_refresh_token>",
    "created_at": "2026-09-06T14:30:00Z",
    "updated_at": "2026-09-06T14:35:00Z",
    "user_id": "64d0f682ab1234567890cdef"
  }
  ```
- **Error Responses**:
  - `400 Bad Request`: Unauthorized access if a `USER` attempts to access another user's profile.
  - `401 Unauthorized`: Missing, expired, or malformed token.
  - `500 Internal Server Error`: Database query error.

---

## Token Lifecycle

- **Access Token**: Valid for 24 hours. Contains user claims (`email`, `first_name`, `last_name`, `uid`, `user_type`).
- **Refresh Token**: Valid for 7 days (168 hours). Used to regenerate access tokens when expired.
