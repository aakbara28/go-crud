# Go CRUD API - Student Management

CRUD API for student data management following Go coding standards.

## Project Structure

```
go-crud/
├── internal/               # Private application code
│   ├── constparam/        # Constants and parameters
│   ├── environment/       # Environment configuration
│   ├── logging/           # Logging utilities
│   ├── utils/             # Utility functions
│   └── v1/handler/response/ # Response handlers
├── config/                # Database configuration
├── controller/            # HTTP handlers
├── model/                 # Data models
├── service/               # Business logic
├── *.json                 # Environment configuration files
├── go.mod
├── main.go
└── README.md
```

## Features

- ✅ Clean directory structure following coding standards
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

## API Documentation

Swagger UI available at: `http://localhost:8080/swagger/`

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