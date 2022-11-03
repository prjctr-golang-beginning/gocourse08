package views

import (
	"encoding/json"
	"fmt"
	"mvc/app/models"
)

func PrintUser(u *models.User) error {
	if res, err := json.Marshal(u); err != nil {
		return err
	} else {
		fmt.Println(string(res))
		return nil
	}
}
