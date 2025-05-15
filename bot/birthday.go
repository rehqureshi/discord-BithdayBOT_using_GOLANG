package bot

import (
	"encoding/json"
	"fmt"
	"os"
)

type Birthday struct {
	UserID   string `json:"user_id"`
	Date     string `json:"date"`
	Username string `json:"username"`
}

var birthday []Birthday

func loadBirthdays() {
	file, err := os.ReadFile("./birthday.json")
	if err != nil {
		fmt.Println(err.Error())
		birthday = []Birthday{}
		return
	}
	_ = json.Unmarshal(file, &birthday)
}

func saveBirthdays() {
	data, _ := json.MarshalIndent(birthday, "", " ")
	_ = os.WriteFile("./birthday.json", data, 0644)
}
