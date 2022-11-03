// Rough brush stroke
package main

import (
	"fmt"
)

func main() {
	// user controller
	// create

	// course controller
	// create
	// add user
	// delete user

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
		fmt.Println(" ----- User created")
		goto Onemoretime
	case `cc`:
		fmt.Println(" ----- Course created")
		goto Onemoretime
	case `au`:
		fmt.Printf("<user id>?")
		var userId int
		_, _ = fmt.Scan(&userId)
		fmt.Printf("<course id>?")
		var courseId int
		_, _ = fmt.Scan(&courseId)
		fmt.Println(" ----- User added")
		goto Onemoretime
	case `du`:
		fmt.Printf("<user id>?")
		var userId int
		_, _ = fmt.Scan(&userId)
		fmt.Printf("<course id>?")
		var courseId int
		_, _ = fmt.Scan(&courseId)
		fmt.Println(" ----- User deleted")
		goto Onemoretime
	case `e`:
		break
	}

	fmt.Println("Closed by user")
}
