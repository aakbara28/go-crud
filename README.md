# Go CRUD API - Student Management

CRUD API for student data management following Go coding standards.

## Project Structure

```
go-crud/
├── cmd/                   # Application entry points
│   └── main.go
├── internal/              # Private application code
│   ├── constparam/        # Constants and parameters
│   ├── environment/       # Environment configuration
│   ├── logging/           # Logging utilities
│   ├── utils/             # Utility functions
│   └── v1/                # API version 1
│       ├── controllers/   # HTTP handlers
│       ├── handler/       # Request/Response structures
│       │   └── response/
│       ├── models/        # Data models
│       └── services/      # Business logic
├── config/                # Database configuration
├── docs/                  # API documentation (Swagger)
├── *.json                 # Environment configuration files
├── go.mod
├── go.sum
├── main.go
└── README.md
```

## Features

- ✅ Versioned API structure (v1)
- ✅ Logging pattern with requestId and requestKey
- ✅ Consistent error handling
- ✅ Standardized response format
- ✅ Environment-based configuration
- ✅ Database transaction management
- ✅ Swagger documentation
- ✅ UUID for request tracking

## Configuration

### Environment Files

- `local.json` - Local development
- `dev.json` - Development environment  
- `prod.json` - Production environment

### Database Configuration

Update file konfigurasi sesuai environment:

```json
{
  "database": {
    "host": "localhost",
    "port": "5432",
    "user": "your_user",
    "password": "your_password",
    "name": "your_database",
    "sslmode": "disable"
  },
  "server": {
    "port": "8080"
  }
}
```

## Installation

### Prerequisites
- Go 1.19 or higher
- PostgreSQL database
- Git

### Setup

1. Clone the repository:
```bash
git clone <repository-url>
cd go-crud
```

2. Install dependencies:
```bash
go mod download
```

3. Configure environment:
```bash
# Copy and edit configuration file
cp local.json.example local.json
# Edit database connection settings
```

4. Run database migrations:
```bash
# The application will auto-migrate on startup
# Or manually create the student table:
CREATE TABLE students (
    student_id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    age INTEGER NOT NULL,
    major VARCHAR(255) NOT NULL
);
```

## Running the Application

```bash
# Local environment (default)
go run main.go

# Development environment
go run main.go -env=dev

# Production environment
go run main.go -env=prod

```

## API Endpoints

### Base URL: `/v1`

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/v1/student` | Create new student |
| GET | `/v1/student/{studentId}` | Get student by Student ID |
| PUT | `/v1/student/{studentId}` | Update student by Student ID |
| DELETE | `/v1/student/{studentId}` | Delete student by Student ID |

### Request Headers

- `requestId` (optional) - Request tracking ID, auto-generated if not provided

### Response Format

```json
{
  "requestId": "uuid-string",
  "requestKey": "operation-key",
  "code": 200,
  "status": "OK",
  "message": "Success message",
  "data": {}
}
```

## API Examples

### Create Student
```bash
curl -X POST http://localhost:8080/v1/student \
  -H "Content-Type: application/json" \
  -H "requestId: 12345-67890" \
  -d '{
    "studentId": "STD001",
    "name": "John Doe",
    "age": 20,
    "major": "Computer Science"
  }'
```

### Get Student
```bash
curl -X GET http://localhost:8080/v1/student/STD001 \
  -H "requestId: 12345-67890"
```

### Update Student
```bash
curl -X PUT http://localhost:8080/v1/student/STD001 \
  -H "Content-Type: application/json" \
  -H "requestId: 12345-67890" \
  -d '{
    "name": "James Smith",
    "age": 21,
    "major": "Software Engineering"
  }'
```

### Delete Student
```bash
curl -X DELETE http://localhost:8080/v1/student/STD001 \
  -H "requestId: 12345-67890"
```

## API Documentation

Swagger UI available at: `http://localhost:8080/swagger/`

## Testing

```bash
# Run tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific test
go test ./internal/v1/services -v
```

## Logging

Semua request dan response akan dicatat dengan format:

```json
{
  "requestId": "uuid",
  "requestKey": "operation|identifier", 
  "code": "200",
  "type": "REQUEST|RESPONSE",
  "method": "POST|GET|PUT|DELETE",
  "function": "function_name",
  "data": {},
  "logType": "log_student_crud",
  "timestamp": "2006-01-02 15:04:05"
}
```

## Dependencies

- **Gorilla Mux**: HTTP web framework
- **GORM**: ORM library
- **UUID**: Unique identifier generation
- **Swagger**: API documentation
- **PostgreSQL**: Database driver
