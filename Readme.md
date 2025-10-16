# Auth Service

This is a simple authentication service built with Go. It provides endpoints for user registration, login, and profile management.

## Features

- User Registration
- User Login (JWT based)
- User Profile Retrieval
- User Profile Update
- Password Hashing with bcrypt
- JWT Authentication Middleware

## Technologies Used

- Go
- Gorilla Mux (for routing)
- JWT (JSON Web Tokens)
- Bcrypt (for password hashing)
- PostgreSQL (for database) - *Note: This README assumes a PostgreSQL database. If you're using a different database, adjust the connection string and migrations accordingly.*

## Getting Started

### Prerequisites

- Go (version 1.18 or higher)
- PostgreSQL database
- Git

### Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/your-username/auth-service.git
   cd auth-service
   ```

2. **Set up environment variables:**

   Create a `.env` file in the root directory of the project and add the following:

   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=your_db_user
   DB_PASSWORD=your_db_password
   DB_NAME=auth_db
   JWT_SECRET=your_jwt_secret_key
   PORT=8080

   ```
   *Replace `your_db_user`, `your_db_password`, `your_jwt_secret_key` with your actual credentials.*

3. **Install dependencies:**

   ```bash
   go mod tidy
   ```

4. **Run database migrations:**

   You'll need to create a `users` table in your PostgreSQL database. Here's an example SQL schema:

   ```sql
   CREATE TABLE users (
       id SERIAL PRIMARY KEY,
       username VARCHAR(255) UNIQUE NOT NULL,
       email VARCHAR(255) UNIQUE NOT NULL,
       password_hash VARCHAR(255) NOT NULL,
       created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
       updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
   );
