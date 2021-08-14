package handlers

import (
	"context"
	"strconv"

	"github.com/basudebpalwebdev/admybrand-friend-finder/api/dbconn"
	db "github.com/basudebpalwebdev/admybrand-friend-finder/db/sqlc"
	"github.com/gofiber/fiber/v2"
)

// ShowAccount godoc
// @Summary List users
// @Description Get a paginated list of users
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param limit path int false "limit"
// @Param page_no path int false "page_no"
// @Success 200 {object} db.User
// @Header 200 {string} Token "qwerty"
// @Router /users [get]
func ListUsers(c *fiber.Ctx) error {
	pageNo, err := strconv.Atoi(c.Query("page_no", "1"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	limit, err := strconv.Atoi(c.Query("limit", "10"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	if pageNo < 1 {
		pageNo = 1
	}
	if limit < 10 {
		limit = 10
	}
	userList, err := dbconn.DBQueries.ListUsers(context.Background(), db.ListUsersParams{
		Limit:  int32(limit),
		Offset: int32((pageNo - 1) * limit),
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(userList)
}
