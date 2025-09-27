# Social Media Platform Backend

A high-performance RESTful API built with Go and Fiber framework for social networking, featuring real-time notifications, friendship management, and content sharing capabilities.

## 📱 Overview

The Social Media Platform Backend is a modern backend service built with Go that provides a complete solution for social networking. It supports user management, friendship connections, post creation, and real-time notifications when friends share new content.

## ✨ Features

- **User Management**
  - User registration with validation
  - User profile retrieval
  - Account deletion functionality
  - Password security (10-15 characters)
  - List all users in the system

- **Friendship System**
  - Create bidirectional friendships
  - View user's friend list
  - Remove friendships
  - UUID-based relationship tracking

- **Post Management**
  - Create posts with content validation (max 2000 chars)
  - View user's posts
  - Delete posts
  - User-specific post feeds

- **Real-time Notifications**
  - Channel-based notification system
  - Automatic friend notifications on new posts
  - Persistent notification channels
  - System hydration on startup

- **Performance Optimization**
  - Redis caching for friend lists
  - 24-hour cache TTL
  - Automatic cache invalidation
  - Connection pooling

- **Security & Validation**
  - Input validation using struct tags
  - UUID-based identifiers
  - Comprehensive error handling
  - Request body size limits (16MB)

## 🛠️ Tech Stack

- **Backend Framework:** Fiber v2.52.9
- **Language:** Go 1.25
- **Database:** PostgreSQL with GORM
- **Cache:** Redis
- **ORM:** GORM v1.31.0
- **Validation:** go-playground/validator v10
- **UUID:** google/uuid
- **Build Tool:** Go Modules

## 📋 Prerequisites

Before running this application, ensure you have the following installed:

- Go 1.25 or higher
- PostgreSQL 12 or higher
- Redis Server 6.0 or higher
- Git

## 🚀 Installation & Setup

### 1. Clone the Repository

```bash
git clone https://github.com/HrithikKeshri10/Social-Media-Platform-Backend.git
cd Social-Media-Platform-Backend
```

### 2. Configure PostgreSQL Database

Create a PostgreSQL database and enable UUID extension:

```sql
CREATE DATABASE social_media_db;
CREATE USER your_db_user WITH PASSWORD 'your_password';
GRANT ALL PRIVILEGES ON DATABASE social_media_db TO your_db_user;

-- Connect to the database
\c social_media_db;

-- Enable UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
```

Update the database credentials in `internals/database/db.go`:

```go
dsn := "user=your_db_user database=social_media_db sslmode=disable password=your_password"
```

### 3. Configure Redis

Ensure Redis is running on your system. The application supports environment configuration:

```bash
# Optional: Set Redis URL (defaults to localhost:6379)
export REDIS_URL=redis://localhost:6379

# Or with authentication
export REDIS_URL=redis://:password@localhost:6379
```

### 4. Install Dependencies

```bash
go mod download
go mod tidy
```

### 5. Run the Application

```bash
go run cmd/main.go
```

The application will start on `http://localhost:3015`

## 📡 API Endpoints

### User Management

| Method | Endpoint | Description | Access |
|--------|----------|-------------|--------|
| POST | `/socio/users` | Register a new user | Public |
| GET | `/socio/users` | Get all users | Public |
| GET | `/socio/users/:id` | Get user by ID | Public |
| DELETE | `/socio/users/:id` | Delete user account | Public |

### Friendship Management

| Method | Endpoint | Description | Access |
|--------|----------|-------------|--------|
| POST | `/socio/friends` | Create friendship connection | Public |
| GET | `/socio/friends/:id` | Get user's friend list | Public |
| DELETE | `/socio/friends/:id?f_id=:friend_id` | Remove friendship | Public |

### Post Management

| Method | Endpoint | Description | Access |
|--------|----------|-------------|--------|
| POST | `/socio/users/:id/posts` | Create a new post | Public |
| GET | `/socio/users/:id/posts` | Get user's posts | Public |
| DELETE | `/socio/users/:id/posts/:post_id` | Delete a post | Public |

## 📝 API Request Examples

### Register a User

```json
POST /socio/users
{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "SecurePass123"
}
```

**Response:**
```json
{
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "name": "John Doe",
    "email": "john@example.com",
    "password": "SecurePass123",
    "created_at": "2024-12-01T10:30:00Z"
}
```

### Create Friendship

```json
POST /socio/friends
{
    "user_id": "550e8400-e29b-41d4-a716-446655440000",
    "friend_id": "660e8400-e29b-41d4-a716-446655440001"
}
```

**Response:**
```json
{
    "id": 1,
    "user_id": "550e8400-e29b-41d4-a716-446655440000",
    "friend_id": "660e8400-e29b-41d4-a716-446655440001",
    "created_at": "2024-12-01T10:35:00Z"
}
```

### Create a Post

```json
POST /socio/users/550e8400-e29b-41d4-a716-446655440000/posts
{
    "content": "Hello World! This is my first post on this amazing platform."
}
```

**Response:**
```json
{
    "id": "770e8400-e29b-41d4-a716-446655440002",
    "content": "Hello World! This is my first post on this amazing platform.",
    "user_id": "550e8400-e29b-41d4-a716-446655440000",
    "created_at": "2024-12-01T10:40:00Z"
}
```

### Get User's Friends

```json
GET /socio/friends/550e8400-e29b-41d4-a716-446655440000
```

**Response:**
```json
[
    {
        "friend_id": "660e8400-e29b-41d4-a716-446655440001"
    },
    {
        "friend_id": "880e8400-e29b-41d4-a716-446655440003"
    }
]
```

## 🗄️ Database Schema

### Users Table
- **id** (UUID, Primary Key, auto-generated)
- **name** (VARCHAR, max 100 chars)
- **email** (VARCHAR, unique)
- **password** (VARCHAR, 10-15 chars)
- **created_at** (TIMESTAMP)
- **updated_at** (TIMESTAMP)

### Friendships Table
- **id** (SERIAL, Primary Key)
- **user_id** (UUID, Foreign Key → Users.id)
- **friend_id** (UUID, Foreign Key → Users.id)
- **created_at** (TIMESTAMP)
- **updated_at** (TIMESTAMP)
- **deleted_at** (TIMESTAMP, soft delete)
- **Unique Constraint:** (user_id, friend_id)

### Posts Table
- **id** (UUID, Primary Key, auto-generated)
- **content** (TEXT, max 2000 chars)
- **user_id** (UUID, Foreign Key → Users.id)
- **created_at** (TIMESTAMP)
- **updated_at** (TIMESTAMP)

## ⚙️ Configuration

Key configuration properties:

```go
// Server Configuration
Port: ":3015"
BodyLimit: 16 * 1024 * 1024 // 16MB

// Database Configuration
Database: "social_media_db"
User: "your_db_user"
SSLMode: "disable"

// Redis Configuration
Default: "localhost:6379"
Cache TTL: 24 hours

// Validation Rules
Password: min=10, max=15
Post Content: max=2000
User Name: max=100
```

### Environment Variables

```bash
# Redis Configuration (Optional)
REDIS_URL=redis://localhost:6379

# Future: Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_db_user
DB_PASSWORD=your_password
DB_NAME=social_media_db
```

## 🔒 Security

The application implements several security measures:

- **UUID Identifiers:** Prevents sequential ID attacks
- **Input Validation:** Comprehensive validation for all inputs
- **Password Requirements:** Enforced 10-15 character passwords
- **Request Size Limits:** 16MB body limit to prevent DoS
- **Error Handling:** Centralized error handling with appropriate status codes
- **Soft Deletes:** Data recovery capability for friendships

## 📂 Project Structure

```
Social-Media-Platform-Backend/
├── cmd/
│   ├── app/
│   │   └── app.go              # Application initialization
│   └── main.go                 # Entry point
├── controllers/
│   ├── friendships/
│   │   └── friends.go          # Friendship handlers
│   ├── posts/
│   │   └── posts.go            # Post handlers
│   └── users/
│       └── users.go            # User handlers
├── internals/
│   ├── cache/
│   │   └── cache.go            # Redis configuration
│   ├── config/
│   │   └── db.go               # Database migrations
│   ├── database/
│   │   └── db.go               # PostgreSQL connection
│   ├── dto/
│   │   ├── friendships.go      # Friendship DTOs
│   │   ├── posts.go            # Post DTOs
│   │   └── users.go            # User DTOs
│   ├── notifications/
│   │   └── notifications.go    # Real-time notifications
│   ├── server/
│   │   ├── handlers.go         # Route handlers
│   │   ├── middleware.go       # Middleware setup
│   │   └── server.go           # Server configuration
│   └── validator/
│       ├── users.go            # User validation
│       └── utils.go            # Validation utilities
├── models/
│   ├── friendship/
│   │   └── friendship.go       # Friendship model
│   ├── posts/
│   │   └── posts.go            # Post model
│   └── users/
│       └── users.go            # User model
├── routes/
│   ├── friendships.go          # Friendship routes
│   ├── posts.go                # Post routes
│   └── users.go                # User routes
├── services/
│   ├── friendships/
│   │   └── friendships.go      # Friendship business logic
│   ├── posts/
│   │   └── posts.go            # Post business logic
│   └── users/
│       └── users.go            # User business logic
├── .gitignore
├── go.mod                      # Go module definition
└── go.sum                      # Dependency checksums
```

## 🎯 Features in Detail

### Real-time Notification System
- Channel-based architecture for instant notifications
- Automatic friend notification when posts are created
- Persistent channels maintained for active users
- System hydration on startup to restore active connections
- Graceful shutdown handling with context cancellation

### Caching Strategy
- Redis caching for friend list queries
- 24-hour TTL for optimal performance
- Automatic cache invalidation on friendship changes
- JSON serialization for complex data structures

### Request Validation
- Struct tag-based validation
- Custom validation messages
- Field-level constraints:
  - Email format validation
  - Password length (10-15 characters)
  - Post content limit (2000 characters)
  - Name length limit (100 characters)

### Error Handling
- Centralized error handler
- Consistent error response format
- Appropriate HTTP status codes
- Recovery middleware for panic handling
- Detailed logging for debugging

## 🚦 Performance Optimizations

- **Connection Pooling:** Efficient database connection management
- **Goroutines:** Concurrent notification delivery
- **Caching Layer:** Redis for frequently accessed data
- **Optimized Queries:** GORM query optimization
- **Middleware Pipeline:** Efficient request processing

## 👨‍💻 Author

**Hrithik Keshri**
- LinkedIn: [Hrithik Keshri](https://linkedin.com/in/hrithikkeshri10)
---

Made with ❤️ by Hrithik Keshri
