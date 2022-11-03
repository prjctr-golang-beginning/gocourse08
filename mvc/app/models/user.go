package models

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	OnCourse bool   `json:"on_course"`
}

func (u *User) GetEmail() string {
	return u.Email
}
