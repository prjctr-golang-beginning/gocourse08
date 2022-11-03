package controllers

import (
	"fmt"
	"mvc/app/models"
)

type CourseController struct {
	courses []*models.Course
}

func (cc *CourseController) CreateCourse(c *models.Course) {
	cc.courses = append(cc.courses, c)
}

func (cc *CourseController) AddUser(cId int, u *models.User) *models.Course {
	for _, course := range cc.courses {
		if course.Id() == cId {
			course.AddUser(u)
			return course
		}
	}

	return nil
}

func (cc *CourseController) DeleteUser(cId int, u *models.User) {
	for _, course := range cc.courses {
		if course.Id() == cId {
			if err := course.DeleteUser(u); err != nil {
				fmt.Println(err)
			}
		}
	}
}
