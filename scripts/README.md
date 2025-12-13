# Scripts Directory

This directory contains utility scripts for managing the Infinity service.

## create_admin.go

Generates SQL statements to create an admin user with a bcrypt-hashed password.

### Usage

#### Using Make (Recommended)

Generate admin SQL with default password (`Admin@123`):
```bash
make create-admin
```

Generate admin SQL with custom password:
```bash
make create-admin PASSWORD="MySecurePassword123"
```

#### Using Go Run Directly

With default password:
```bash
go run scripts/create_admin.go
```

With custom password:
```bash
go run scripts/create_admin.go "MySecurePassword123"
```

Save output to file:
```bash
go run scripts/create_admin.go "MyPassword" > db/mysql/seed/0001-create-admin.sql
```

### Output

The script generates SQL statements that:
1. Insert the `admin` role (if it doesn't exist)
2. Insert an admin user with:
   - Unique UUID
   - Bcrypt-hashed password
   - Email: `admin@infinity.local`
   - Name: `Admin User`
   - Status: `active`

### Example Output

```sql
-- Insert admin role
INSERT INTO `roles` (`name`, `created_at`, `updated_at`)
VALUES ('admin', NOW(), NOW())
ON DUPLICATE KEY UPDATE `updated_at` = NOW();

-- Insert admin user
INSERT INTO `users` (`role_id`, `user_iid`, `name`, `email`, `status`, `created_by`, `password`, `created_at`, `updated_at`)
VALUES (
    (SELECT id FROM roles WHERE name = 'admin'),
    '2bfa645a-43e8-4048-830f-cbb50746f124',
    'Admin User',
    'admin@infinity.local',
    'active',
    'system',
    '$2a$10$KHkPyWLfMyNJ72BszthxfefGuNx88mqWgWy81HSZN3Lpu/ZaDuyq2',
    NOW(),
    NOW()
);
```

### Default Credentials

- **Email**: `admin@infinity.local`
- **Password**: `Admin@123` (or your custom password)

## Database Management

### Initialize Database Schema
```bash
make db-init
```

### Seed Admin User
```bash
make db-seed
```

### Reset Database (Init + Seed)
```bash
make db-reset
```

### Notes

- The generated password hash uses bcrypt with default cost (10)
- Each time the script runs, it generates a new UUID for the admin user
- The admin role insert uses `ON DUPLICATE KEY UPDATE` to avoid errors if the role already exists
- Change the default email (`admin@infinity.local`) in the script if needed

