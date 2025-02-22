package controllers

import (
	"referals/src/database"
	"referals/src/models"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v3"
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


	user := models.User{
		FirstName: data["first_name"],
		LastName: data["last_name"],
		Email: data["email"],
		IsAmbassador: false,
	}

	user.SetPassword(data["password"])

	database.DB.Create(&user)

	return c.JSON(user)
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
	
	
	if err := user.ComparePassword(data["password"]); err  != nil{
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message" : "Invalid credentials",
		})
	}

	payload := jwt.StandardClaims{
		Subject : strconv.Itoa(int(user.Id)),
		ExpiresAt : time.Now().Add(time.Hour * 24).Unix(),
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString([]byte("secret"))

	if err != nil {
		c.Status(fiber.StatusBadRequest)

		return c.JSON(fiber.Map{
			"message" : "Invalid credentials",
		})
	}

	cookie := fiber.Cookie{
		Name: "jwt",
		Value : token,
		Expires : time.Now().Add(time.Hour*24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message" : "success",
	})
}

func User(c fiber.Ctx) error{
	cookie := c.Cookies("jwt")
	
	token,err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func (token *jwt.Token) (interface{},error)  {
		return []byte("secret"),nil
	})

	if err != nil || !token.Valid{
		c.Status(fiber.StatusUnauthorized)

		return c.JSON(fiber.Map{
			"message" : "unauthenticated",
		})
	}

	payload := token.Claims.(*jwt.StandardClaims)

	var user models.User

	database.DB.Where("id=?", payload.Subject).First(&user)

	return c.JSON(user)
}
