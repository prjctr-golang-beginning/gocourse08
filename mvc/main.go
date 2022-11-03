// Rough brush stroke
package main

import (
	"fmt"
	"mvc/app/controllers"
	"mvc/app/models"
	"mvc/app/views"
)

func main() {
	var ci int
	uc := controllers.UserController{}
	cc := controllers.CourseController{}

Onemoretime:
	fmt.Printf("CHOOSE WHAT TO DO: " +
		"\nCreate user: cu" +
		"\nCreate course: cc" +
		"\nAdd user to couse: au" +
		"\nDelete user from course: du" +
		"\nExit: e\n")
	var command string
	_, _ = fmt.Scan(&command)

	switch command {
	case `cu`:
		u := &models.User{`robot 1`, `1@1.com`, false}
		uc.CreateUser(u)
		fmt.Println(" ----- User created")
		if err := views.PrintUser(u); err != nil {
			fmt.Println(err)
		}
		goto Onemoretime
	case `cc`:
		ci++
		c := models.NewCourse(ci)
		cc.CreateCourse(c)
		fmt.Println(" ----- Course created")
		if err := views.PrintCourse(c); err != nil {
			fmt.Println(err)
		}
		goto Onemoretime
	case `au`:
		fmt.Printf("<user email>?")
		var userEmail string
		_, _ = fmt.Scan(&userEmail)
		user2add := uc.FindUser(userEmail)
		if user2add == nil {
			fmt.Printf("User with %s not found", userEmail)
			goto Onemoretime
		}
		fmt.Printf("<course id>?")
		var courseId int
		_, _ = fmt.Scan(&courseId)
		course := cc.AddUser(courseId, user2add)
		fmt.Println(" ----- User added")
		if err := views.PrintCourse(course); err != nil {
			fmt.Println(err)
		}
		goto Onemoretime
	case `du`:
		fmt.Printf("<user email>?")
		var userEmail string
		_, _ = fmt.Scan(&userEmail)
		user2delete := uc.FindUser(userEmail)
		if user2delete == nil {
			fmt.Printf("User with %s not found", userEmail)
			goto Onemoretime
		}
		fmt.Printf("<course id>?")
		var courseId int
		_, _ = fmt.Scan(&courseId)
		cc.DeleteUser(courseId, user2delete)
		fmt.Println(" ----- User deleted")
		goto Onemoretime
	case `e`:
		break
	}

	fmt.Println("Closed by user")
}
