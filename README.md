# Infinity Service

A secure authentication and user management service built with Go, Fiber, and GORM.

## Features

- ✅ Secure HMAC-SHA256 token-based authentication
- ✅ Role-based access control (RBAC)
- ✅ Session management with expiration
- ✅ Bcrypt password hashing
- ✅ MySQL database with GORM ORM
- ✅ RESTful API design
- ✅ Structured logging with slog

## Quick Start

### Prerequisites

- Go 1.21 or higher
- MySQL 8.0 or higher
- Make (optional, for using Makefile commands)

### Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd infinity-be
```

2. Install dependencies:
```bash
go mod download
go mod vendor
```

3. Setup the database:
```bash
# Create database
mysql -u root -p -e "CREATE DATABASE infinity_service;"

# Run migrations
mysql -u root -p infinity_service < db/mysql/deploy/0000-init.sql

# Seed initial data (creates roles and admin user)
mysql -u root -p infinity_service < db/mysql/seed/0000-init.sql
```

4. Configure the application:
```bash
# Edit config_files/service-config.json
# Update database credentials and other settings
```

5. **IMPORTANT: Update the secret key** (see Security Configuration below)

6. Run the application:
```bash
go run cmd/qms-engine/main.go
```

The server will start on `http://localhost:8089` (or the port specified in config).

## Security Configuration

### ⚠️ Update Secret Key Before Deployment

The default secret key in `config_files/service-config.json` is **NOT SECURE** for production use.

**Current (INSECURE):**
```json
{
  "security": {
    "secretKey": "is-this-secret-key?"
  }
}
```

**Update to a strong secret key:**
```json
{
  "security": {
    "secretKey": "your-very-long-random-secret-key-at-least-32-characters-or-more"
  }
}
```

**Generate a secure secret key:**

```bash
# Option 1: Using openssl
openssl rand -base64 32

# Option 2: Using Python
python3 -c "import secrets; print(secrets.token_urlsafe(32))"

# Option 3: Using Go
go run -c 'package main; import ("crypto/rand"; "encoding/base64"; "fmt"); func main() { b := make([]byte, 32); rand.Read(b); fmt.Println(base64.URLEncoding.EncodeToString(b)) }'
```

**Best Practices:**
- Minimum 32 characters recommended
- Use alphanumeric and special characters
- Never commit production secrets to version control
- Use environment variables or secure vaults in production
- Rotate secret keys periodically

## Token Security

This service implements **HMAC-SHA256 token authentication**:

- **Client receives**: Raw token (44 chars, base64url encoded)
- **Database stores**: Hashed token (64 chars, hex encoded)
- **Secret key**: Cryptographically binds tokens to your application
- **Security**: Even with database access, stored tokens cannot be used

### How it works:

1. **Login**: Generates random token → Hashes with secret → Stores hash → Returns raw token
2. **Authentication**: Receives raw token → Hashes with secret → Looks up hash in DB → Validates

See [AUTHENTICATION_GUIDE.md](AUTHENTICATION_GUIDE.md) for detailed documentation.

## Default Credentials

After running the seed script:

- **Email**: `admin@example.com`
- **Password**: `Admin@123`
- **Role**: admin

**⚠️ Change these credentials immediately in production!**

## API Documentation

See [AUTHENTICATION_GUIDE.md](AUTHENTICATION_GUIDE.md) for complete API documentation including:

- Login endpoint
- User creation
- Authentication flow
- Error handling
- Example requests

## Project Structure

```
infinity-be/
├── cmd/qms-engine/          # Application entry point
├── config_files/            # Configuration files
├── db/mysql/               # Database scripts
│   ├── deploy/            # Schema migrations
│   └── seed/              # Seed data
├── handler/               # HTTP handlers
├── internal/
│   ├── common/           # Common utilities and errors
│   ├── config/           # Configuration setup
│   ├── entity/           # Database entities
│   ├── model/            # Request/response models
│   ├── repository/       # Data access layer
│   └── service/          # Business logic
│       └── user/        # User service with token utils
├── server/               # Server configuration
│   ├── config/          # Server config types
│   └── middleware/      # HTTP middleware
└── vendor/              # Vendored dependencies
```

## Development

### Generate Password Hash

To generate a bcrypt hash for a password:

```bash
go run hash.go
```

Edit the password in `hash.go` before running.

### Database Migrations

Migrations are located in `db/mysql/deploy/`. To add new migrations:

1. Create a new file: `XXXX-description.sql`
2. Run: `mysql -u root -p infinity_service < db/mysql/deploy/XXXX-description.sql`

## Environment Variables

While the current implementation uses `service-config.json`, you can extend it to support environment variables:

- `DB_HOST`: Database host
- `DB_PORT`: Database port
- `DB_NAME`: Database name
- `DB_USER`: Database username
- `DB_PASSWORD`: Database password
- `TOKEN_SECRET`: Secret key for HMAC token generation
- `PORT`: Server port

## Security Features

1. ✅ HMAC-SHA256 token generation with secret key
2. ✅ Bcrypt password hashing (cost: 10)
3. ✅ Session expiration (24 hours default)
4. ✅ Role-based access control
5. ✅ Timing-attack resistant token comparison
6. ✅ Database breach protection (tokens hashed)

## License

[Add your license here]

## Contributing

[Add contribution guidelines]

## Support

For questions or issues, please refer to [AUTHENTICATION_GUIDE.md](AUTHENTICATION_GUIDE.md) or open an issue.

