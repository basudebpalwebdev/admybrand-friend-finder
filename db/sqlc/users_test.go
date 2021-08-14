package db

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	arg := CreateUserParams{
		Username:    "test.create.user",
		Description: "Test user creation",
		Dob:         time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
		Address:     "Test user address",
		CreatedAt:   time.Now().UTC(),
	}
	createdUser, crErr := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, crErr)
	require.NotEmpty(t, createdUser)
	require.NotZero(t, createdUser.ID)
	require.Equal(t, arg.Address, createdUser.Address)
	require.Equal(t, arg.Description, createdUser.Description)
	require.Equal(t, arg.Dob, createdUser.Dob)
	require.InDelta(t, arg.CreatedAt.UnixNano(), createdUser.CreatedAt.UnixNano(), 1000)
}

func TestListUsers(t *testing.T) {
	for i := 0; i < 20; i++ {
		arg := CreateUserParams{
			Username:    "test.list.user." + fmt.Sprintf("%d", i),
			Description: "Test user#" + fmt.Sprintf("%d", i) + " list",
			Dob:         time.Date(2000-i, time.January, 1, 0, 0, 0, 0, time.UTC),
			Address:     "Test user address",
			CreatedAt:   time.Now().UTC(),
		}
		createdUser, crErr := testQueries.CreateUser(context.Background(), arg)

		require.NoError(t, crErr)
		require.NotEmpty(t, createdUser)
		require.NotZero(t, createdUser.ID)
	}

	// Checking for normal args
	userList, listErr := testQueries.ListUsers(context.Background(), ListUsersParams{
		Limit:  10,
		Offset: 0,
	})
	require.NoError(t, listErr)
	require.Equal(t, 10, len(userList))

	// Checking for invalid limit
	userList, listErr = testQueries.ListUsers(context.Background(), ListUsersParams{
		Limit:  -1,
		Offset: 0,
	})
	require.Error(t, listErr)
	require.Empty(t, userList)

	// Checking for invalid offset
	userList, listErr = testQueries.ListUsers(context.Background(), ListUsersParams{
		Limit:  1,
		Offset: -1,
	})
	require.Error(t, listErr)
	require.Empty(t, userList)

	// Checking for large offset value (bigger than total user count)
	userList, listErr = testQueries.ListUsers(context.Background(), ListUsersParams{
		Limit:  1,
		Offset: 100000,
	})
	require.NoError(t, listErr)
	require.Empty(t, userList)
}

func TestUpdateUserDetails(t *testing.T) {
	arg := CreateUserParams{
		Username:    "test.update.user",
		Description: "Test user update",
		Dob:         time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
		Address:     "Test user address",
		CreatedAt:   time.Now().UTC(),
	}
	createdUser, crErr := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, crErr)
	require.NotEmpty(t, createdUser)
	require.NotZero(t, createdUser.ID)

	updateArgs := UpdateUserDetailsParams{
		ID:          createdUser.ID,
		Description: "Updated test user",
		Dob:         arg.Dob.AddDate(1, 1, 1),
		Address:     "Updated test address",
	}
	updatedUser, updErr := testQueries.UpdateUserDetails(context.Background(), updateArgs)
	require.NoError(t, updErr)
	require.NotEmpty(t, updatedUser)
	require.Equal(t, createdUser.ID, updatedUser.ID)
	require.Equal(t, updateArgs.Address, updatedUser.Address)
	require.Equal(t, updateArgs.Description, updatedUser.Description)
	require.Equal(t, updateArgs.Dob, updatedUser.Dob)

	// Checking for invalid ID
	updatedUser, updErr = testQueries.UpdateUserDetails(context.Background(), UpdateUserDetailsParams{
		ID:          -1,
		Description: "lorem",
	})
	require.Error(t, updErr)
	require.Empty(t, updatedUser)
}

func TestGetUser(t *testing.T) {
	arg := CreateUserParams{
		Username:    "test.get.user",
		Description: "Test get user",
		Dob:         time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
		Address:     "Test user address",
		CreatedAt:   time.Now().UTC(),
	}
	createdUser, crErr := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, crErr)
	require.NotEmpty(t, createdUser)
	require.NotZero(t, createdUser.ID)

	foundUser, getErr := testQueries.GetUser(context.Background(), createdUser.ID)
	require.NoError(t, getErr)
	require.NotEmpty(t, foundUser)
	require.Equal(t, createdUser.ID, foundUser.ID)
	require.Equal(t, createdUser.Address, foundUser.Address)
	require.Equal(t, createdUser.Username, foundUser.Username)
	require.Equal(t, createdUser.Description, foundUser.Description)
	require.Equal(t, createdUser.Dob, foundUser.Dob)
	require.Equal(t, createdUser.CreatedAt, foundUser.CreatedAt)

	// Checking for invalid ID
	foundUser, getErr = testQueries.GetUser(context.Background(), -1)
	require.Error(t, getErr)
	require.Empty(t, foundUser)
}

func TestFindUserByUsername(t *testing.T) {
	arg := CreateUserParams{
		Username:    "test.find.user",
		Description: "Test find user",
		Dob:         time.Date(2006, time.January, 1, 0, 0, 0, 0, time.UTC),
		Address:     "Test user address",
		CreatedAt:   time.Now().UTC(),
	}
	createdUser, crErr := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, crErr)
	require.NotEmpty(t, createdUser)
	require.NotZero(t, createdUser.ID)

	foundUser, findErr := testQueries.FindUserByUsername(context.Background(), createdUser.Username)
	require.NoError(t, findErr)
	require.NotEmpty(t, foundUser)
	require.Equal(t, createdUser.ID, foundUser.ID)
	require.Equal(t, createdUser.Address, foundUser.Address)
	require.Equal(t, createdUser.Description, foundUser.Description)
	require.Equal(t, createdUser.Dob, foundUser.Dob)
	require.Equal(t, createdUser.CreatedAt, foundUser.CreatedAt)

	// Checking for empty Username
	foundUser, findErr = testQueries.FindUserByUsername(context.Background(), "")
	require.Error(t, findErr)
	require.Empty(t, foundUser)
}

func TestDeleteUser(t *testing.T) {
	arg := CreateUserParams{
		Username:    "test.delete.user",
		Description: "Test delete user",
		Dob:         time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
		Address:     "Test user address",
		CreatedAt:   time.Now().UTC(),
	}
	createdUser, crErr := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, crErr)
	require.NotEmpty(t, createdUser)
	require.NotZero(t, createdUser.ID)

	delErr := testQueries.DeleteUser(context.Background(), createdUser.ID)
	require.NoError(t, delErr)

	// Checking for user with deleted ID
	foundUser, getErr := testQueries.GetUser(context.Background(), createdUser.ID)
	require.Error(t, getErr)
	require.Empty(t, foundUser)
}
