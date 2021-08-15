package handlers

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/basudebpalwebdev/admybrand-friend-finder/api/dbconn"
	db "github.com/basudebpalwebdev/admybrand-friend-finder/db/sqlc"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// ListUsers godoc
// @Summary List users
// @Description Get a paginated list of users
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
		pageNo = 1
	}
	limit, err := strconv.Atoi(c.Query("limit", "10"))
	if err != nil {
		limit = 10
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

// CreateNewUser godoc
// @Summary Create new users
// @Description Create a new user
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param username body string true "Please enter your desired username"
// @Param description body string true "Please describe yourself"
// @Param dob body string true "Please enter your date of birth(yyyy-mm-dd)"
// @Param address body string true "Please enter your address"
// @Success 200 {object} db.User
// @Header 200 {string} Token "qwerty"
// @Router /users [post]
func CreateNewUser(c *fiber.Ctx) error {
	type NewUserReqParams struct {
		Username    string `json:"username" validate:"required,min=3,max=64"`
		Description string `json:"description" validate:"required,min=3,max=128"`
		Dob         string `json:"dob" validate:"required,min=10,max=10"`
		Address     string `json:"address" validate:"required,min=3,max=128"`
	}

	newUserReqParams := new(NewUserReqParams)
	parseErr := c.BodyParser(newUserReqParams)
	if parseErr != nil {
		return c.Status(fiber.StatusBadRequest).SendString(parseErr.Error())
	}
	validate := validator.New()
	valErr := validate.Struct(newUserReqParams)
	if valErr != nil {
		return c.Status(fiber.StatusBadRequest).SendString(valErr.Error())
	}
	dob := strings.Split(newUserReqParams.Dob, "-")
	year, yrErr := strconv.Atoi(dob[0])
	month, mthErr := strconv.Atoi(dob[1])
	day, dayErr := strconv.Atoi(dob[2])
	if yrErr != nil || mthErr != nil || dayErr != nil {
		return c.Status(fiber.StatusBadRequest).SendString("D.O.B. value should be in the format 'YYYY-MM-DD'")
	}
	newUser, crErr := dbconn.DBQueries.CreateUser(context.Background(), db.CreateUserParams{
		Username:    newUserReqParams.Username,
		Description: newUserReqParams.Description,
		Dob:         time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC).UTC(),
		Address:     newUserReqParams.Address,
		CreatedAt:   time.Now().UTC(),
	})
	if crErr != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(crErr.Error())
	}
	return c.Status(fiber.StatusOK).JSON(newUser)
}

// GetUser godoc
// @Summary Get an user by ID
// @Description Get an user by ID
// @Accept  json
// @Produce  json
// @Param id path number false "Please enter the id"
// @Success 200 {object} db.User
// @Header 200 {string} Token "qwerty"
// @Router /users/{id} [get]
func GetUser(c *fiber.Ctx) error {
	userIDStr := c.Params("id")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("User id value is not a valid number")
	}
	foundUser, getErr := dbconn.DBQueries.GetUser(context.Background(), userID)
	if getErr != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("User not found")
	}
	return c.Status(fiber.StatusOK).JSON(foundUser)
}

// FindUserByUsername godoc
// @Summary Find an user by Username
// @Description Find an user by Username
// @Accept  json
// @Produce  json
// @Param username path string true "Please enter the username"
// @Success 200 {object} db.User
// @Header 200 {string} Token "qwerty"
// @Router /users/find/{username} [get]
func FindUserByUsername(c *fiber.Ctx) error {
	username := c.Params("username")

	foundUser, getErr := dbconn.DBQueries.FindUserByUsername(context.Background(), username)
	if getErr != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("User not found")
	}
	return c.Status(fiber.StatusOK).JSON(foundUser)
}

// UpdateNewUser godoc
// @Summary Update new users
// @Description Update a new user
// @Accept  json
// @Produce  json
// @Param id path string true "Enter the id"
// @Param description body string true "Update description"
// @Param dob body string true "Update date of birth(yyyy-mm-dd)"
// @Param address body string true "Update address"
// @Success 200 {object} db.User
// @Header 200 {string} Token "qwerty"
// @Router /users [post]
func UpdateUser(c *fiber.Ctx) error {
	type UpdateUserReqParams struct {
		Description string `json:"description" validate:"required,min=3,max=128"`
		Dob         string `json:"dob" validate:"required,min=10,max=10"`
		Address     string `json:"address" validate:"required,min=3,max=128"`
	}

	updateUserReqParams := new(UpdateUserReqParams)
	parseErr := c.BodyParser(updateUserReqParams)
	if parseErr != nil {
		return c.Status(fiber.StatusBadRequest).SendString(parseErr.Error())
	}
	validate := validator.New()
	valErr := validate.Struct(updateUserReqParams)
	if valErr != nil {
		return c.Status(fiber.StatusBadRequest).SendString(valErr.Error())
	}
	dob, success := convertDate(updateUserReqParams.Dob, c)
	if !success {
		return c.Status(fiber.StatusBadRequest).SendString("D.O.B. value should be in the format 'YYYY-MM-DD'")
	}
	id, idErr := strconv.ParseInt(c.Params("id"), 10, 64)
	if idErr != nil {
		return c.Status(fiber.StatusBadRequest).SendString("User ID value illegal")
	}
	updatedUser, crErr := dbconn.DBQueries.UpdateUserDetails(context.Background(), db.UpdateUserDetailsParams{
		ID:          id,
		Description: updateUserReqParams.Description,
		Dob:         dob,
		Address:     updateUserReqParams.Address,
	})
	if crErr != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(crErr.Error())
	}
	return c.Status(fiber.StatusOK).JSON(updatedUser)
}

func convertDate(dob string, c *fiber.Ctx) (time.Time, bool) {
	dobArr := strings.Split(dob, "-")
	if len(dobArr) < 3 {
		return time.Time{}, false
	}
	year, yrErr := strconv.Atoi(dobArr[0])
	month, mthErr := strconv.Atoi(dobArr[1])
	day, dayErr := strconv.Atoi(dobArr[2])
	if yrErr != nil || mthErr != nil || dayErr != nil {
		return time.Time{}, false
	}
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC).UTC(), true
}

// DeleteUser godoc
// @Summary Delete an user by ID
// @Description Delete an user by ID
// @Accept  json
// @Produce  json
// @Param id path string true "Please enter the id"
// @Success 200 {object} string
// @Header 200 {string} Token "qwerty"
// @Router /users/{id} [delete]
func DeleteUser(c *fiber.Ctx) error {
	userIDStr := c.Params("id")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("User id value is not a valid number")
	}
	deletedUser, delErr := dbconn.DBQueries.DeleteUser(context.Background(), userID)
	if delErr != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("User not found")
	}
	if deletedUser.ID == 0 {
		return c.Status(fiber.StatusInternalServerError).SendString("User not found")
	}
	return c.Status(fiber.StatusOK).SendString(fmt.Sprintf("Deleted user with id: %d", userID))
}
