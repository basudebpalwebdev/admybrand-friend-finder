package handlers_test

import (
	"database/sql"
	"fmt"
	"log"
	"net/http/httptest"
	"testing"

	"github.com/basudebpalwebdev/admybrand-friend-finder/api"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"

	_ "github.com/lib/pq"
)

// func TestMain(m *testing.M) {
// 	dbConn, err := sql.Open("postgres", "postgresql://basu:Basudeb@2021@localhost:5432/admybrand_friend_finder?sslmode=disable")
// 	if err != nil {
// 		log.Fatal("Cannot connect to the database :", err)
// 	}
// 	testServer := api.NewServer(dbConn)
// }

const URL = "localhost:9999"

var testServer *api.Server

func TestMain(t *testing.M) {
	dbConn, err := sql.Open("postgres", "postgresql://basu:Basudeb@2021@localhost:5432/admybrand_friend_finder?sslmode=disable")
	if err != nil {
		log.Fatal("Cannot connect to the database :", err)
	}
	testServer = api.NewServer(dbConn)
}

func TestRoot(t *testing.T) {
	testServer.Start(URL)
	req := httptest.NewRequest("GET", URL, nil)
	res, err := testServer.App.Test(req, -1)

	assert.Equal(t, nil, err)
	assert.Equal(t, fiber.StatusOK, res.StatusCode)
	assert.Equal(t, "Hello, World!", res.Body)
	testServer.Stop()
}

func TestListUsers(t *testing.T) {
	testServer.Start(URL)
	req := httptest.NewRequest("GET", fmt.Sprintf("%s/users", URL), nil)
	res, err := testServer.App.Test(req, -1)

	assert.Equal(t, nil, err)
	assert.Equal(t, fiber.StatusOK, res.StatusCode)

	req = httptest.NewRequest("GET", fmt.Sprintf("%s/users?limit=%d&page_no=%d", URL, 12, 2), nil)
	res, err = testServer.App.Test(req, -1)

	assert.Equal(t, nil, err)
	assert.Equal(t, fiber.StatusOK, res.StatusCode)

	testServer.Stop()
}
