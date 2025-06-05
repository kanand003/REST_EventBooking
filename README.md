# REST Event Booking API

A comprehensive REST API built with Go (Golang) and Gin framework for managing events and registrations with user authentication.

## Features

- **User Management**:
  - User registration and authentication
  - JWT-based token authentication
  - Password hashing with bcrypt

- **Event Management**:
  - Create, read, update, and delete events
  - Event owner authorization
  - Event listing and filtering

- **Registration System**:
  - Register for events
  - Cancel registrations
  - Authorization checks for registration actions

- **Database**:
  - SQLite database for ease of setup
  - Structured schema with relationships between users, events, and registrations

## Tech Stack

- **Language**: Go (Golang)
- **Web Framework**: Gin
- **Database**: SQLite
- **Authentication**: JWT (JSON Web Tokens)
- **Password Security**: bcrypt

## Project Structure

```
│   api.db                    # SQLite database file
│   go.mod                    # Go module file
│   go.sum                    # Go dependencies lockfile
│   main.go                   # Application entry point
│   README.md                 # Project documentation
│
├───api-test                  # HTTP test requests
│       cancel-registration.http
│       create-event.http
│       create-user.http
│       delete-event.http
│       get-event.http
│       getsingle-event.http
│       login.http
│       register.http
│       update-event.http
│
├───db                        # Database configuration
│       db.go                 # Database initialization and schema
│
├───middleware                # HTTP middleware
│       auth.go               # Authentication middleware
│
├───models                    # Data models
│       event.go              # Event model and methods
│       user.go               # User model and methods
│
├───routes                    # API routes and handlers
│       events.go             # Event-related endpoints
│       registrations.go      # Registration-related endpoints
│       routes.go             # Route registration
│       users.go              # User-related endpoints
│
└───utils                     # Utility functions
        hash.go               # Password hashing utilities
        jwt.go                # JWT token management
```

## API Endpoints

### Public Endpoints

- **GET /events** - List all events
- **GET /events/:id** - Get a specific event by ID
- **POST /signup** - Register a new user
- **POST /login** - Login and get authentication token

### Protected Endpoints (Require Authentication)

- **POST /api/events** - Create a new event
- **PUT /api/events/:id** - Update an existing event (owner only)
- **DELETE /api/events/:id** - Delete an event (owner only)
- **POST /api/events/:id/register** - Register for an event
- **DELETE /api/events/:id/register** - Cancel an event registration

## Getting Started

### Prerequisites

- Go 1.19 or higher
- SQLite3

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/REST_EventBooking.git
   cd REST_EventBooking
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Run the application:
   ```bash
   go run main.go
   ```

   The server will start on `http://localhost:8080`

### Using the API

1. **Create a user account**:
   ```http
   POST http://localhost:8080/signup
   Content-Type: application/json

   {
     "email": "user@example.com",
     "password": "securepassword"
   }
   ```

2. **Login to get a JWT token**:
   ```http
   POST http://localhost:8080/login
   Content-Type: application/json

   {
     "email": "user@example.com",
     "password": "securepassword"
   }
   ```

3. **Create an event (with auth token)**:
   ```http
   POST http://localhost:8080/api/events
   Content-Type: application/json
   Authorization: YOUR_JWT_TOKEN

   {
     "name": "Tech Conference 2023",
     "description": "Annual technology conference",
     "location": "Convention Center",
     "start_time": "2023-10-01T09:00:00Z",
     "end_time": "2023-10-01T18:00:00Z"
   }
   ```

4. **Register for an event**:
   ```http
   POST http://localhost:8080/api/events/1/register
   Authorization: YOUR_JWT_TOKEN
   ```

## Database Schema

### Users Table
- `id` - Primary key
- `email` - User email (unique)
- `password` - Hashed password

### Events Table
- `id` - Primary key
- `name` - Event name
- `description` - Event description
- `start_time` - Event start time
- `end_time` - Event end time
- `location` - Event location
- `user_id` - Creator ID (foreign key)
- `created_at` - Creation timestamp
- `updated_at` - Last update timestamp

### Registrations Table
- `id` - Primary key
- `event_id` - Event ID (foreign key)
- `user_id` - User ID (foreign key)
- `created_at` - Registration timestamp

## Security

- Passwords are hashed using bcrypt before storage
- Authentication is handled via JWT tokens
- Route protection with middleware
- Authorization checks for event management

## Future Enhancements

- Add user roles (admin, organizer, attendee)
- Implement event categories and tags
- Add search functionality
- Implement pagination for event listings
- Add email notifications for registrations

## License

[MIT License](LICENSE)
