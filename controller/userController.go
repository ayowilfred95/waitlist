package controller

import (
	"log"

	"github.com/ayowilfred95/database"
	"github.com/ayowilfred95/model"
	"github.com/gofiber/fiber/v2"
)

type user struct {
	Name  string `json:"name"`
	Email string  `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
}

func CreateUser(c *fiber.Ctx) error {
	// Get data from the req body and validate the body
	// we store the data first in a struct
	var body user
	if err := c.BodyParser(&body); err != nil {
		log.Println("Error parsing request body:", err)
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid body",
		})
	}

	// craete  a new user object with data from the request data
	newUser := &model.User{Name: body.Name, Email: body.Email, PhoneNumber:body.PhoneNumber}

	// create the user in the database
	result := database.DB.Create(newUser) 

	if result.Error != nil {
		log.Println("Database error:", result.Error)
		return c.Status(500).JSON(fiber.Map{
			"error" : "Failed to create user",
		})
	}

	response := fiber.Map{
		"name":body.Name,
		"email":body.Email,
		"phoneNumber":body.PhoneNumber,
	}
	return c.Status(201).JSON(response)

}


// create a user with just email input field
func CreateUserWithEmail(c *fiber.Ctx)error{
		// Get data from the req body and validate the body
	// we store the data first in a struct
	var body user
	if err := c.BodyParser(&body); err != nil {
		log.Println("Error parsing request body:", err)
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid body",
		})
	}

	// check if a user with the same email already exists in the database
	var existingUser model.User
	result := database.DB.Where("email=?", body.Email).First(&existingUser)
	if result.Error == nil {
		// user with the same email already exists
		return c.Status(400).JSON(fiber.Map{
			"error": "User with this email already exists",
		})
	}

		// craete  a new user object with data from the request data
		newUser := &model.User{Email: body.Email}

		// create the user in the database
	result = database.DB.Create(newUser) 

	if result.Error != nil {
		log.Println("Database error:", result.Error)
		return c.Status(500).JSON(fiber.Map{
			"error" : "Failed to create user",
		})
	}


		response := fiber.Map{
			"email":body.Email,
		}
		return c.Status(201).JSON(response)
}


// create a struct to hold the data while being updated in the database
type updateUser struct {
	Name  string `json:"name"`
	Email string  `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
}

func UpdateUser(c *fiber.Ctx) error {
	// validate the body
	var body updateUser
	if err := c.BodyParser(&body); err != nil {
		log.Println("Error parsing request body:", err)
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid body",
		})
	}

	// get the email from the request body
	email:= body.Email
	if email == "" {
		return c.Status(400).JSON(fiber.Map{
			"error":"Please provide an email",
		})
	}

	// Find the user by email in the database
	var existingUser model.User
	result := database.DB.Where("email = ?", email).First(&existingUser)
	if result.Error != nil {
		log.Println("Database error:", result.Error)
		return c.Status(500).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	// Updata user information if provided in the requestr
	if body.Name != "" {
		existingUser.Name = body.Name
	}
	if body.PhoneNumber != "" {
		existingUser.PhoneNumber = body.PhoneNumber
	}
	
	// save the updated user to the database
	result = database.DB.Save(&existingUser)
	if result.Error != nil {
		log.Println("Database error:", result.Error)
		return c.Status(500).JSON(fiber.Map{
			"error":"Failed to update user",
		})
	}

	response := fiber.Map{
		"name": existingUser.Name,
		"email": existingUser.Email,
		"phoneNumber": existingUser.PhoneNumber,
	}
	return c.Status(200).JSON(response)
}


// Get all users created
func GetUsers(c *fiber.Ctx)error{
	// get all users
// store the return data from the daabase in a slice
	var users []model.User

	result := database.DB.Find(&users)

	if result.Error !=nil {
		return c.Status(401).JSON(fiber.Map{
			"error":" error fetching users data from database",
		})
	} 
	
		return c.Status(200).JSON(fiber.Map{
			"users": users,
		})
}


// delete users 

// func DeleteUser(c *fiber.Ctx)error{
// 	// 
	

// }
