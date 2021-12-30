package user

import "github.com/google/uuid"

var users []User

func New(ID string, name string) User {

	return User{
		ID:   ID,
		Name: name,
	}
}

func Create(name string) (User, error) {
	user := New(uuid.New().String(), name)
	users = append(users, user)

	return user, nil
}

func FindByName(name string) *User {
	for _, user := range users {
		if user.Name == name {
			return &user
		}
	}
	return nil
}

func FindByID(ID string) *User {
	for _, user := range users {
		if user.ID == ID {
			return &user
		}
	}
	return nil
}

func Update(user User) bool {
	for i := range users {
		if users[i].ID == user.ID {
			users[i].Name = user.Name
			return true
		}
	}
	return false
}

func DeleteByID(ID string) bool {
	for i := range users {
		if users[i].ID == ID {
			users[i] = users[len(users)-1]
			users = users[:len(users)-1]
			return true
		}
	}
	return false
}

func GetAll() []User {

	return users
}
