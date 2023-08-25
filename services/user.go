package services

import (
	"firstGolangModule/interfaces"
	"fmt"
)

type UserServices interface {
	Create(interfaces.User) (interfaces.User, error)
	Update(interfaces.User) (interfaces.User, error)
	FindAll() ([]interfaces.User, error)
	Delete(string) (string, error)
}

type usersService struct {
	// users []interfaces.User
	users map[string]interfaces.User
}

// Create implements UserServices.
func (services *usersService) Create(user interfaces.User) (interfaces.User, error) {
	// services.users = append(services.users, user)
	_, exists := services.users[user.USERNAME]
	if exists {
		return interfaces.User{}, fmt.Errorf("%v %v %v", "user with this username:", user.USERNAME, "already exists")
	}
	services.users[user.USERNAME] = user
	return user, nil
}

// Delete implements UserServices.
func (services *usersService) Delete(username string) (string, error) {
	_, exists := services.users[username]
	if !exists {
		return "", fmt.Errorf("%v", "No Such Users found")
	}
	delete(services.users, username)
	return "user deleted successfully", nil
}

// FindAll implements UserServices.
func (services *usersService) FindAll() ([]interfaces.User, error) {
	var user []interfaces.User
	if len(services.users) == 0 {
		return []interfaces.User{}, fmt.Errorf("%v", "No Users found")
	}
	for _, val := range services.users {
		user = append(user, val)
	}

	return user, nil
}

// Update implements UserServices.
func (services *usersService) Update(user interfaces.User) (interfaces.User, error) {
	_, exists := services.users[user.USERNAME]
	if !exists {
		return interfaces.User{}, fmt.Errorf("%v", "No Such Users found")
	}
	services.users[user.USERNAME] = user
	return user, nil
}

func New() UserServices {
	return &usersService{
		users: make(map[string]interfaces.User),
	}
}
