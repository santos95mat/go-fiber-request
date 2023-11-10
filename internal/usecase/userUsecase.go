package usecase

import (
	"encoding/json"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/santos95mat/go-fiber-request/internal/dto"
)

type UserRequest struct{}

func (*UserRequest) GetMany(c *fiber.Ctx) error {
	req := fiber.Get("http://localhost:3000/v1/user")
	// to set headers
	token := c.Cookies("Authorization")
	req.Cookie("Authorization", token)

	statusCode, data, errs := req.Bytes()
	if len(errs) > 0 {
		return c.Status(statusCode).JSON(fiber.Map{
			"error": errs,
		})
	}

	var users []dto.UserResponseDTO
	jsonErr := json.Unmarshal(data, &users)

	if jsonErr != nil {
		var resBody dto.ResBodyDTO
		jsonErr = json.Unmarshal(data, &resBody)

		if jsonErr != nil {
			return c.Status(statusCode).JSON(fiber.Map{
				"error": jsonErr,
			})
		}

		return c.Status(statusCode).JSON(resBody)
	}

	return c.Status(statusCode).JSON(users)
}

func (*UserRequest) Create(c *fiber.Ctx) error {
	var userData dto.UserCreateDTO
	err := c.BodyParser(&userData)

	if err != nil {
		panic(err)
	}

	req := fiber.Post("http://localhost:3000/v1/user")
	// to set JSON BODY
	req.JSON(userData)

	statusCode, data, errs := req.Bytes()

	if len(errs) > 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": errs,
		})
	}

	var user dto.ResBodyDTO
	jsonErr := json.Unmarshal(data, &user)
	if jsonErr != nil {
		panic(jsonErr)
	}

	return c.Status(statusCode).JSON(user)
}

func (*UserRequest) Login(c *fiber.Ctx) error {
	var userLogin dto.UserLoginDTO
	err := c.BodyParser(&userLogin)

	if err != nil {
		panic(err)
	}

	req := fiber.Post("http://localhost:3000/v1/login")
	// to set JSON BODY
	req.JSON(userLogin)

	statusCode, data, errs := req.Bytes()

	if len(errs) > 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": errs,
		})
	}

	var user dto.ResBodyDTO
	jsonErr := json.Unmarshal(data, &user)
	if jsonErr != nil {
		panic(jsonErr)
	}

	cookie := new(fiber.Cookie)
	cookie.Name = "Authorization"
	cookie.Value = user.Token
	cookie.Expires = time.Now().Add(24 * time.Hour)

	c.Cookie(cookie)

	return c.Status(statusCode).JSON(user)
}
