package views

import (
	"encoding/json"
	"fmt"
	"mvc/app/models"
)

func PrintCourse(c *models.Course) error {
	if res, err := json.Marshal(c); err != nil {
		return err
	} else {
		fmt.Println(string(res))
		return nil
	}
}
