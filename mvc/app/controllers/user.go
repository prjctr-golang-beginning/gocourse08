package controllers

import "mvc/app/models"

type UserController struct {
	users []*models.User
}

func (cc *UserController) CreateUser(u *models.User) {
	cc.users = append(cc.users, u)
}

func (cc *UserController) FindUser(email string) *models.User {
	for _, user := range cc.users {
		if user.GetEmail() == email {
			return user
		}
	}

	return nil
}
