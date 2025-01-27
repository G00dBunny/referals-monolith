package controllers

import (
	"referals/src/database"
	"referals/src/models"

	"github.com/gofiber/fiber/v3"
	"golang.org/x/crypto/bcrypt"
)


func Register(c fiber.Ctx) error {

	var data map[string]string

	if err := c.Bind().JSON(&data); err != nil {
        return err
    }

	if data["password"] != data["password_confirm"]{
		c.Status(fiber.StatusBadRequest).SendString("Bad Request")
		return c.JSON(fiber.Map{
			"message" : "password do no match",
		})
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 12)

	user := models.User{
		FirstName: data["first_name"],
		LastName: data["last_name"],
		Email: data["email"],
		Password : password,
		IsAmbassador: false,
	}

	database.DB.Create(&user)

	return c.JSON(fiber.Map{
		"message" : "hello",
	})
}

func Login(c fiber.Ctx) error {

	var data map[string]string

	if err := c.Bind().JSON(&data); err != nil {
        return err
    }

	var user models.User

	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.Id == 0 {
		c.Status(fiber.StatusBadRequest)

		return c.JSON(fiber.Map{
			"message" : "Invalid credentials",
		})
	}
	
	
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])) ; err  != nil{
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message" : "Invalid credentials",
		})
	}

	return c.JSON(user)
}
