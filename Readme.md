# Auth Service

This is a simple authentication service built with Go. It provides endpoints for user registration, login, and profile management.

## Features

- User Registration
- User Login (JWT based)
- User Profile Retrieval
- User Profile Update
- Password Hashing with bcrypt
- JWT Authentication Middleware
- CRUD for Book Collection

## Technologies Used

- Go
- Gorilla Mux (for routing)
- JWT (JSON Web Tokens)
- Bcrypt (for password hashing)
- Sqlite3 (for database)

## Getting Started

### Prerequisites

- Go (version 1.25 or higher)
- SQLite3
- Git

### Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/your-username/auth-service.git
   cd auth-service
   ```

2. **Set up environment variables:**

   *This uses sqlite3 for the Database, upon running this app, it will be created automatically*
   [API Documentation](https://documenter.getpostman.com/view/28337725/2sB3QQJSy5)

3. **Install dependencies:**

   ```bash
   go mod tidy
   ```

4. **Run database migrations:**

   You'll need to create a `users` and `books` table in your PostgreSQL database. Here's an example SQL schema:

   ```sql
   CREATE TABLE users (
       id SERIAL PRIMARY KEY,
       username VARCHAR(255) UNIQUE NOT NULL,
       email VARCHAR(255) UNIQUE NOT NULL,
       password_hash VARCHAR(255) NOT NULL,
       created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
       updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
   );

   CREATE TABLE books (
       id SERIAL PRIMARY KEY,
       title VARCHAR(255) UNIQUE NOT NULL,
       author VARCHAR(255) UNIQUE NOT NULL,
       created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
       updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
   );

