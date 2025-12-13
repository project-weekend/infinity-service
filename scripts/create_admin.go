package main

import (
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// This script generates SQL statements to create an admin user with a bcrypt-hashed password
// Usage:
//   go run scripts/create_admin.go [password]
//   If no password is provided, defaults to "Admin@123"

func main() {
	password := "Admin@123"
	if len(os.Args) > 1 {
		password = os.Args[1]
	}

	// Generate bcrypt hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}

	// Generate UUID for admin user
	adminUUID := uuid.New().String()

	fmt.Println("-- ==============================================")
	fmt.Println("-- Admin User Creation Script")
	fmt.Println("-- ==============================================")
	fmt.Println("-- Generated at:", fmt.Sprintf("%v", os.Args[0]))
	fmt.Println("-- Password (plaintext for reference):", password)
	fmt.Println("-- Password (bcrypt hash):", string(hashedPassword))
	fmt.Println("-- User UUID:", adminUUID)
	fmt.Println("-- ==============================================")
	fmt.Println()

	// SQL to insert admin role (if not exists)
	fmt.Println("-- Insert admin role")
	fmt.Println("INSERT INTO `roles` (`name`, `created_at`, `updated_at`)")
	fmt.Println("VALUES ('admin', NOW(), NOW())")
	fmt.Println("ON DUPLICATE KEY UPDATE `updated_at` = NOW();")
	fmt.Println()

	// SQL to insert admin user
	fmt.Println("-- Insert admin user")
	fmt.Println("-- Note: Replace 'system' in created_by if needed")
	fmt.Printf("INSERT INTO `users` (`role_id`, `user_code`, `email`, `status`, `created_by`, `password`, `created_at`, `updated_at`)\n")
	fmt.Printf("VALUES (\n")
	fmt.Printf("    (SELECT id FROM roles WHERE name = 'admin'),\n")
	fmt.Printf("    '%s',\n", adminUUID)
	fmt.Printf("    'admin@infinity.local',\n")
	fmt.Printf("    'active',\n")
	fmt.Printf("    'system',\n")
	fmt.Printf("    '%s',\n", string(hashedPassword))
	fmt.Printf("    NOW(),\n")
	fmt.Printf("    NOW()\n")
	fmt.Printf(");\n")
	fmt.Println()

	fmt.Println("-- ==============================================")
	fmt.Println("-- Login credentials:")
	fmt.Println("-- Email: admin@infinity.local")
	fmt.Println("-- Password:", password)
	fmt.Println("-- ==============================================")
}
