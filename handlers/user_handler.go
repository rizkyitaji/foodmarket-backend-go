package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizkyitaji/foodmarket-backend-go/database"
	"github.com/rizkyitaji/foodmarket-backend-go/models"
)

// GET all users
func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	err := database.DB.Select(&users, "SELECT * FROM users")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(users)
}

// CREATE user
func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	_, err := database.DB.Exec(
		// "INSERT INTO users (id, name, email, password, phone_number, address, profile_photo, created_at, updated_at) VALUES(nextval('users_id_seq'::regclass), $1, $2, $3, '$4', '$5', '$6', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)",
		"INSERT INTO users (name, email, password, phone_number, address, profile_photo, created_at, updated_at) VALUES($1, $2, $3, '$4', '$5', '$6')",
		user.Name, user.Email, user.Password, user.PhoneNumber, user.Address, user.ProfilePhoto,
	)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "User created successfully",
	})
}
