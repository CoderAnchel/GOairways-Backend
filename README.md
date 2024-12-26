# GOairways-Backend
![Sin título-2024-11-22-1530](https://github.com/user-attachments/assets/646b89f5-ddbc-42fa-9171-06f408b22f4c)


## Overview
This project is a backend application built using the Go programming language and the Fiber framework. The application follows a microservices architecture, employing JWT authentication with 2FA (Two-Factor Authentication), bcrypt for password hashing, and secure HttpOnly cookies for session management. The backend enables users to perform a variety of actions such as creating and approving routes, managing buy models, registering plans, creating accounts, and logging in securely.

---

## Features

- **Microservices Architecture**: Each service is isolated for scalability and maintainability.
- **Authentication**:
  - JWT-based authentication with Two-Factor Authentication (2FA).
  - Secure password storage using bcrypt.
  - Session management using HttpOnly cookies.
- **Core Functionalities**:
  - User registration and account management.
  - Secure login with optional 2FA.
  - Creation and approval of routes.
  - Management of buy models.
  - Plane registration and updates.
- **Security Features**:
  - Enforced secure practices using bcrypt and HttpOnly cookies.
  - Validation of JWT tokens for API request authorization.

---

## Tech Stack

### Language
- **Go**

### Framework
- **Fiber** (a fast, lightweight web framework inspired by Express.js)

### Libraries & Tools
- **bcrypt**: For secure password hashing.
- **JWT**: For token-based authentication.
- **Fiber Middleware**: Various middleware for request handling, error management, and CORS.
- **gorilla/mux**: For additional routing needs.

### Databases
- **PostgreSQL**: For user and transactional data.
- **Redis**: For session storage and caching.

---

## Installation

1. **Prerequisites**:
   - Go (version 1.20 or later).
   - PostgreSQL database.
   - Redis server.

2. **Clone the Repository**:
   ```bash
   git clone https://github.com/your-repo/backend-go-fiber
   cd backend-go-fiber
   ```

3. **Environment Setup**:
   - Create a `.env` file and configure the following variables:
     ```env
     DB_HOST=localhost
     DB_PORT=5432
     DB_USER=yourusername
     DB_PASSWORD=yourpassword
     DB_NAME=yourdatabase
     REDIS_URL=redis://localhost:6379
     JWT_SECRET=yourjwtsecret
     ```

4. **Install Dependencies**:
   ```bash
   go mod tidy
   ```

5. **Run Database Migrations** (if any):
   ```bash
   go run migrate.go
   ```

6. **Run the Application**:
   ```bash
   go run main.go
   ```

---
