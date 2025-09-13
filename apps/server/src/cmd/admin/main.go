package main

import (
	"fmt"
	"log"
	"sodnix/apps/server/src/common"
	"sodnix/apps/server/src/config"
	"sodnix/apps/server/src/database"

	users "sodnix/apps/server/src/modules/users"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func main() {
	config.LoadEnv()
	database.ConnectDatabase()

	db := database.DB

	// Check if an admin user already exists
	var adminUser users.User
	result := db.Where("email = ?", config.ADMIN_EMAIL).First(&adminUser)
	if result.Error == nil {
		fmt.Println("Admin user already exists.")
		return
	}
	if result.Error != gorm.ErrRecordNotFound {
		log.Fatalf("Failed to check for existing admin user: %v", result.Error)
	}

	// Create admin user
	password := config.ADMIN_PASSWORD
	if password == "" {
		log.Fatal("ADMIN_PASSWORD environment variable is not set.")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}

	adminID := uuid.New() // Generate a new UUID for the admin user

	admin := users.User{
		Email:       config.ADMIN_EMAIL,
		Password:    string(hashedPassword),
		DisplayName: "Admin",
		Username:    "admin",
		UUIDTypeModel: common.UUIDTypeModel{
			ID: adminID,
			AuditFields: common.AuditFields{
				CreatedByID: adminID,
				UpdatedByID: adminID,
			},
		},
	}

	if err := db.Create(&admin).Error; err != nil {
		log.Fatalf("Failed to create admin user: %v", err)
	}

	fmt.Println("Admin user created successfully!")
	fmt.Printf("Email: %s\n", config.ADMIN_EMAIL)
	fmt.Printf("Password: %s\n", password) // For demonstration, do not print in production
}
