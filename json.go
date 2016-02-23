package main

import (
	"fmt"
	"encoding/json"
)

type UserProfile struct{
	User User `json:"user"`
	Videos []Video `json:"videos"`
}

type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

type Video struct {
	Id int `json:"id"`
	Link string `json:"link"`
}

func main() {
	up := UserProfile{
		User{1234, "Alan"},
		[]Video{{5678, "http://vibbidi.com/5678"}},
	}
	json, _ := json.Marshal(up);
	fmt.Print(string(json))
}
