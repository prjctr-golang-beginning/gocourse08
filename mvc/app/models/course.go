package models

import (
	"fmt"
	"time"
)

func NewCourse(id int) *Course {
	return &Course{
		id:        id,
		Name:      fmt.Sprintf("course %s", id),
		StartDate: time.Now(),
	}
}

type Course struct {
	id        int
	Name      string    `json:"name"`
	Users     []*User   `json:"users"`
	StartDate time.Time `json:"start_date"`
}

func (c Course) Id() int {
	return c.id
}

func (c Course) AddUser(u *User) {
	c.Users = append(c.Users, u)
}

func (c *Course) DeleteUser(u *User) error {
	for i, user := range c.Users {
		if user.GetEmail() == u.GetEmail() {
			if len(c.Users) == 1 {
				c.Users = []*User{}
			} else {
				c.Users = c.Users[i:len(c.Users)]
			}
			return nil
		}
	}

	return fmt.Errorf("user %s not found", u.GetEmail())
}
