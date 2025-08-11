package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizkyitaji/foodmarket-backend-go/database"
	"github.com/rizkyitaji/foodmarket-backend-go/models"
	"golang.org/x/crypto/bcrypt"
)

// GET Users
func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	err := database.DB.Select(&users, "SELECT * FROM users")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(users)
}

// CREATE User
func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	_, err := database.DB.Exec(
		// "INSERT INTO users (id, name, email, password, phone_number, address, profile_photo, created_at, updated_at) VALUES(nextval('users_id_seq'::regclass), $1, $2, $3, '$4', '$5', '$6', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)",
		"INSERT INTO users (name, email, password, phone_number, address, profile_photo, created_at, updated_at) VALUES($1, $2, $3, '$4', '$5', '$6')",
		user.Name, user.Email, user.Password, user.PhoneNumber, user.Address, user.ProfilePhoto,
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "User created successfully",
	})
}

// Login User
func LoginUser(c *fiber.Ctx) error {
	type LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	var user models.User
	err := database.DB.Get(&user, "SELECT * FROM users WHERE email = $1", req.Email)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid email or password"})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid email or password"})
	}

	return c.JSON(fiber.Map{
		"message": "Login successful",
		"data": fiber.Map{
			"id":            user.ID,
			"name":          user.Name,
			"email":         user.Email,
			"phone_number":  user.PhoneNumber,
			"address":       user.Address,
			"profile_photo": user.ProfilePhoto,
		},
	})
}
