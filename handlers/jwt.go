package handlers

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
)

type JWT struct {
	Expires_in uint32

	Access_token   string
	Provider_token string
	Refresh_token  string
	Token_type     string
}

// CheckJWT will check the request to fetch for JWT tokens and Authenticate the user.
func CheckJWT(c *fiber.Ctx) bool {
	headers := c.GetReqHeaders() // prob wrong
	return headers["access_token"] != ""
}

// CheckJWT will check the request to fetch for JWT tokens and Authenticate the user.
func CheckFillJWT(c *fiber.Ctx) (*JWT, error) {
	log.Println("CheckFillJWT")
	qString := string(c.Request().URI().QueryString())
	log.Println(qString)

	// if headers["access_token"] == "" {
	// 	return nil, errors.New("token is not present")
	// }

	// i, err := strconv.ParseUint(headers["expires_in"], 10, 32)
	// if err != nil {
	// 	return nil, errors.New("not able to parse expiry")
	// }

	// jwt := &JWT{
	// 	Expires_in: uint32(i),

	// 	Access_token:   headers["access_token"],
	// 	Provider_token: headers["provider_token"],
	// 	Refresh_token:  headers["refresh_token"],
	// 	Token_type:     headers["token_type"],
	// }

	return nil, errors.New("err")
}
