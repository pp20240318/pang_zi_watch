package service

import (
	"errors"
	"log"
	"os"

	"watch/backend/internal/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var ErrInvalidCredentials = errors.New("invalid credentials")

func EnsureDefaultAdmin(db *gorm.DB) {
	var count int64
	if err := db.Model(&models.Admin{}).Count(&count).Error; err != nil {
		log.Printf("failed to count admins: %v", err)
		return
	}
	if count > 0 {
		return
	}

	username := os.Getenv("WATCH_DEFAULT_ADMIN")
	password := os.Getenv("WATCH_DEFAULT_ADMIN_PWD")
	if username == "" {
		username = "admin"
	}
	if password == "" {
		password = "admin123"
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	role := models.Role{Name: "superadmin", Permissions: "all"}
	if err := db.Create(&role).Error; err != nil {
		log.Printf("failed to create default role: %v", err)
		return
	}
	admin := models.Admin{Username: username, Password: string(hashed), RoleID: role.ID}
	if err := db.Create(&admin).Error; err != nil {
		log.Printf("failed to create default admin: %v", err)
		return
	}
	log.Printf("default admin created: %s / %s", username, password)
}

func Login(db *gorm.DB, username, password string) (*models.Admin, error) {
	var admin models.Admin
	if err := db.Preload("Role").Where("username = ?", username).First(&admin).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrInvalidCredentials
		}
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password)); err != nil {
		return nil, ErrInvalidCredentials
	}
	return &admin, nil
}

func ChangePassword(db *gorm.DB, adminID uint, oldPassword, newPassword string) error {
	var admin models.Admin
	if err := db.First(&admin, adminID).Error; err != nil {
		return err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(oldPassword)); err != nil {
		return ErrInvalidCredentials
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	admin.Password = string(hashed)
	return db.Save(&admin).Error
}
