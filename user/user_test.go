package user

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreate(t *testing.T) {

	// Given there are no users in the system
	users = []User{}

	// When we create a user named "Naren"
	Create("Naren")

	// there should be one user in the system
	assert.Len(t, users, 1, "there should be 1 user")
	// and name of that user should be Naren
	assert.Equal(t, "Naren", users[0].Name, "the first user should have the name Naren")
}

func TestFindByID(t *testing.T) {

	// create a user named "Naren"
	Create("Naren")

	firstUser := users[0]
	userID := firstUser.ID
	foundUser := FindByID(userID)

	assert.Equal(t, "Naren", foundUser.Name)
	assert.Equal(t, userID, foundUser.ID)
}

func TestGetAll(t *testing.T) {

	// because previous tests can affect
	users = []User{}

	Create("Naren")
	Create("GoLang")
	Create("Playground")

	users := GetAll()

	assert.Len(t, users, 3)
}

func TestUpdate(t *testing.T) {

	// create a user named "Naren"
	Create("Naren")

	firstUser := users[0]
	userID := firstUser.ID
	firstUser.Name = "Go Lang"
	Update(firstUser)

	foundUser := FindByID(userID)

	assert.Equal(t, "Go Lang", foundUser.Name)
	assert.Equal(t, userID, foundUser.ID)
}
