# Sanbercode Golang Backend Development Batch 63 | Quiz 3 - Mini Project

Mini Project:

- Create project for input and manage books, categories, and users authentication
- The database used is a relational database including mysql or postgresql
- Deploy production project on railway

## Features

- **Create User and Login**: Allows creation and login for users.
- **Categories Management**: CRUD operations for managing categories.
- **Books Management**: CRUD operations for managing books.
- **JWT Middleware**: Books and Categories implementing JWT middleware for authentication
- **Swagger Documentation**: Web Interface for API Documentation & Testing

## Installation

#### Make sure go have been installed

1. **Clone the repository**:

   ```bash
   - git clone <project-url>
   - cd <project-folder>

   ```

2. **Install Dependencies**:

   ```bash
   - go mod tidy

   ```

3. **Rename `.env.example` to `.env` and fill in the variables with actual value**:

   ```bash
   - mv .env.example .env

   or rename using gui

   ```

4. **Run Application**:

   ```bash
   - go run main.go
   ```

## Documentation
Navigate to `/swagger/index.html` to view API Documentation

- **Example (local)**: `http://localhost:8080/swagger/index.html`